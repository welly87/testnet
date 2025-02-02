package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"github.com/trisacrypto/testnet/pkg/rvasp/config"
	"github.com/trisacrypto/testnet/pkg/rvasp/jsonpb"
	pb "github.com/trisacrypto/testnet/pkg/rvasp/pb/v1"
	"github.com/trisacrypto/trisa/pkg/ivms101"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DB is a wrapper around a gorm.DB instance that restricts query results to a single
// VASP.
type DB struct {
	db   *gorm.DB
	vasp VASP
}

func NewDB(conf *config.Config) (d *DB, err error) {
	d = &DB{}
	if d.db, err = OpenDB(conf); err != nil {
		return nil, err
	}

	if err = MigrateDB(d.db); err != nil {
		return nil, err
	}

	if err = d.db.Where("name = ?", conf.Name).First(&d.vasp).Error; err != nil {
		return nil, fmt.Errorf("could not fetch VASP info from database: %s", err)
	}

	if conf.Name != d.vasp.Name {
		return nil, fmt.Errorf("expected name %q but have database name %q", conf.Name, d.vasp.Name)
	}

	return d, nil
}

func (d *DB) GetVASP() VASP {
	return d.vasp
}

func (d *DB) GetDB() *gorm.DB {
	return d.db
}

func (d *DB) Query() *gorm.DB {
	return d.db.Where("vasp_id = ?", d.vasp.ID)
}

func (d *DB) Create(value interface{}) *gorm.DB {
	return d.db.Create(value)
}

func (d *DB) Save(value interface{}) *gorm.DB {
	return d.db.Save(value)
}

// LookupAccount by email address or wallet address.
func (d *DB) LookupAccount(account string) *gorm.DB {
	return d.Query().Where("email = ?", account).Or("wallet_address = ?", account)
}

// LookupAnyAccount by email address or wallet address, not restricted to the local
// rVASP.
func (d *DB) LookupAnyAccount(account string) *gorm.DB {
	return d.db.Where("email = ?", account).Or("wallet_address = ?", account)
}

// LookupBeneficiary by email address or wallet address.
func (d *DB) LookupBeneficiary(beneficiary string) *gorm.DB {
	return d.Query().Preload("Provider").Where("email = ?", beneficiary).Or("address = ?", beneficiary)
}

// LookupAnyBeneficiary by email address or wallet address, not restricted to the local
// rVASP.
func (d *DB) LookupAnyBeneficiary(beneficiary string) *gorm.DB {
	return d.db.Preload("Provider").Where("email = ?", beneficiary).Or("address = ?", beneficiary)
}

// LookupIdentity by wallet address.
func (d *DB) LookupIdentity(walletAddress string) *gorm.DB {
	return d.Query().Where("wallet_address = ?", walletAddress)
}

// LookupPending returns the pending transactions.
func (d *DB) LookupPending() *gorm.DB {
	return d.Query().Where("state in (?, ?, ?)", pb.TransactionState_PENDING_SENT, pb.TransactionState_PENDING_RECEIVED, pb.TransactionState_PENDING_ACKNOWLEDGED)
}

// LookupTransaction by envelope ID.
func (d *DB) LookupTransaction(envelope string) *gorm.DB {
	return d.Query().Where("envelope = ?", envelope)
}

// LookupWallet by wallet address.
func (d *DB) LookupWallet(address string) *gorm.DB {
	return d.Query().Where("address = ?", address)
}

// MakeTransaction returns a new Transaction from the originator and beneficiary
// wallet addresses. Note: this does not store the transaction in the database to allow
// the caller to modify the transaction fields before storage.
func (d *DB) MakeTransaction(originator string, beneficiary string) (*Transaction, error) {
	var originatorIdentity, beneficiaryIdentity Identity

	// Fetch originator identity record
	if err := d.LookupIdentity(originator).FirstOrInit(&originatorIdentity, Identity{}).Error; err != nil {
		log.Error().Err(err).Msg("could not lookup originator identity")
		return nil, status.Errorf(codes.FailedPrecondition, "could not lookup originator identity: %s", err)
	}

	// If originator identity does not exist then create it
	if originatorIdentity.ID == 0 {
		originatorIdentity.WalletAddress = originator
		originatorIdentity.Vasp = d.vasp

		if err := d.Create(&originatorIdentity).Error; err != nil {
			log.Error().Err(err).Msg("could not save originator identity")
			return nil, status.Errorf(codes.FailedPrecondition, "could not save originator identity: %s", err)
		}
	}

	// Fetch beneficiary identity record
	if err := d.LookupIdentity(beneficiary).FirstOrInit(&beneficiaryIdentity, Identity{}).Error; err != nil {
		log.Error().Err(err).Msg("could not lookup beneficiary identity")
		return nil, status.Errorf(codes.FailedPrecondition, "could not lookup beneficiary identity: %s", err)
	}

	// If the beneficiary identity does not exist then create it
	if beneficiaryIdentity.ID == 0 {
		beneficiaryIdentity.WalletAddress = beneficiary
		beneficiaryIdentity.Vasp = d.vasp

		if err := d.Create(&beneficiaryIdentity).Error; err != nil {
			log.Error().Err(err).Msg("could not save beneficiary identity")
			return nil, status.Errorf(codes.FailedPrecondition, "could not save beneficiary identity: %s", err)
		}
	}

	return &Transaction{
		Envelope:    uuid.New().String(),
		Originator:  originatorIdentity,
		Beneficiary: beneficiaryIdentity,
		State:       pb.TransactionState_AWAITING_REPLY,
		StateString: pb.TransactionState_AWAITING_REPLY.String(),
		Timestamp:   time.Now(),
		Vasp:        d.vasp,
	}, nil
}

// VASP is a record of known partner VASPs and caches TRISA protocol information. This
// table also contains IVMS101 data that identifies the VASP (but only for the local
// VASP - we assume that VASPs do not have IVMS101 data on each other and have to use
// the directory service for that).
// TODO: modify VASP ID to a GUID
type VASP struct {
	gorm.Model
	Name      string     `gorm:"uniqueIndex;size:255;not null"`
	LegalName *string    `gorm:"column:legal_name;null"`
	URL       *string    `gorm:"null"`
	Country   *string    `gorm:"null"`
	Endpoint  *string    `gorm:"null"`
	PubKey    *string    `gorm:"null"`
	NotAfter  *time.Time `gorm:"null"`
	IVMS101   string     `gorm:"column:ivms101"`
}

// TableName explicitly defines the name of the table for the model
func (VASP) TableName() string {
	return "vasps"
}

type PolicyType string

const (
	SendPartial PolicyType = "SendPartial"
	SendFull    PolicyType = "SendFull"
	SendError   PolicyType = "SendError"
	SyncRepair  PolicyType = "SyncRepair"
	SyncRequire PolicyType = "SyncRequire"
	AsyncRepair PolicyType = "AsyncRepair"
	AsyncReject PolicyType = "AsyncReject"
)

// Returns True if this is a valid policy for the originator
func isValidOriginatorPolicy(policy PolicyType) bool {
	return policy == SendPartial || policy == SendFull || policy == SendError
}

// Returns True if this is a valid policy for the beneficiary
func isValidBeneficiaryPolicy(policy PolicyType) bool {
	return policy == SyncRepair || policy == SyncRequire || policy == AsyncRepair || policy == AsyncReject
}

// Wallet is a mapping of wallet IDs to VASPs to determine where to send transactions.
// Provider lookups can happen by wallet address or by email.
type Wallet struct {
	gorm.Model
	Address           string     `gorm:"uniqueIndex"`
	Email             string     `gorm:"uniqueIndex"`
	OriginatorPolicy  PolicyType `gorm:"column:originator_policy"`
	BeneficiaryPolicy PolicyType `gorm:"column:beneficiary_policy"`
	ProviderID        uint       `gorm:"not null"`
	Provider          VASP       `gorm:"foreignKey:ProviderID"`
	VaspID            uint       `gorm:"not null"`
	Vasp              VASP       `gorm:"foreignKey:VaspID"`
}

// TableName explicitly defines the name of the table for the model
func (Wallet) TableName() string {
	return "wallets"
}

// Account contains details about the transactions that are served by the local VASP.
// It also contains the IVMS 101 data for KYC verification, in this table it is just
// stored as a JSON string rather than breaking it down to the field level. Only
// customers of the VASP have accounts.
type Account struct {
	gorm.Model
	Name          string          `gorm:"not null"`
	Email         string          `gorm:"uniqueIndex;not null"`
	WalletAddress string          `gorm:"uniqueIndex;not null;column:wallet_address"`
	Wallet        Wallet          `gorm:"foreignKey:WalletAddress;references:Address"`
	Balance       decimal.Decimal `gorm:"type:numeric(15,2);default:0.0"`
	Completed     uint64          `gorm:"not null;default:0"`
	Pending       uint64          `gorm:"not null;default:0"`
	IVMS101       string          `gorm:"column:ivms101;not null"`
	VaspID        uint            `gorm:"not null"`
	Vasp          VASP            `gorm:"foreignKey:VaspID"`
}

// TableName explicitly defines the name of the table for the model
func (Account) TableName() string {
	return "accounts"
}

// Transaction holds exchange information to send money from one account to another. It
// also contains the decrypted identity payload that was sent as part of the TRISA
// protocol and the envelope ID that uniquely identifies the message chain.
// TODO: Add a field for the transaction payload marshaled as a string.
type Transaction struct {
	gorm.Model
	Envelope      string              `gorm:"not null"`
	AccountID     uint                `gorm:"not null"`
	Account       Account             `gorm:"foreignKey:AccountID"`
	OriginatorID  uint                `gorm:"column:originator_id;not null"`
	Originator    Identity            `gorm:"foreignKey:OriginatorID"`
	BeneficiaryID uint                `gorm:"column:beneficiary_id;not null"`
	Beneficiary   Identity            `gorm:"foreignKey:BeneficiaryID"`
	Amount        decimal.Decimal     `gorm:"type:decimal(15,8)"`
	Debit         bool                `gorm:"not null"`
	State         pb.TransactionState `gorm:"not null;default:0"`
	StateString   string              `gorm:"column:state_string;not null"`
	Timestamp     time.Time           `gorm:"not null"`
	NotBefore     time.Time           `gorm:"not null"`
	NotAfter      time.Time           `gorm:"not null"`
	Identity      string              `gorm:"not null"`
	Transaction   string              `gorm:"not null"`
	VaspID        uint                `gorm:"not null"`
	Vasp          VASP                `gorm:"foreignKey:VaspID"`
}

// TableName explicitly defines the name of the table for the model
func (Transaction) TableName() string {
	return "transactions"
}

// Set the transaction state to a new value
func (t *Transaction) SetState(state pb.TransactionState) {
	t.State = state
	t.StateString = state.String()
}

// Identity holds raw data for an originator or a beneficiary that was sent as
// part of the transaction process. This should not be stored in the wallet since the
// wallet is a representation of the local VASPs knowledge about customers and bercause
// the identity information could change between transactions. This intermediate table
// is designed to more closely mimic data storage as part of a blockchain transaction.
type Identity struct {
	gorm.Model
	WalletAddress string `gorm:"not null;column:wallet_address;index:wallet_index,unique"`
	Email         string `gorm:"not null"`
	Provider      string `gorm:"not null"`
	VaspID        uint   `gorm:"not null;index:wallet_index,unique"`
	Vasp          VASP   `gorm:"foreignKey:VaspID"`
}

// TableName explicitly defines the name of the table for the model
func (Identity) TableName() string {
	return "identities"
}

// BalanceFloat converts the balance decmial into an exact two precision float32 for
// use with the protocol buffers.
func (a Account) BalanceFloat() float32 {
	bal, _ := a.Balance.Truncate(2).Float64()
	return float32(bal)
}

// Return the VASP associated with the account.
func (a Account) GetVASP(d *DB) (vasp *VASP, err error) {
	vasp = &VASP{}
	if err = d.db.Where("id = ?", a.VaspID).First(vasp).Error; err != nil {
		return nil, err
	}
	return
}

// Return the wallet associated with the account.
func (a Account) GetWallet(db *DB) (wallet *Wallet, err error) {
	wallet = &Wallet{}
	if err = db.Query().Where("address = ?", a.WalletAddress).First(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

// Transactions returns an ordered list of transactions associated with the account
// ordered by the timestamp of the transaction, listing any pending transactions at the
// top. This function may also support pagination and limiting functions, which is why
// we're using it rather than having a direct relationship on the model.
func (a Account) Transactions(db *DB) (records []Transaction, err error) {
	if err = db.Query().Preload(clause.Associations).Where("account_id = ?", a.ID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// Return the account associated with the transaction.
func (t Transaction) GetAccount(db *DB) (account *Account, err error) {
	account = &Account{}
	if err = db.Query().Where("id = ?", t.AccountID).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

// Return the originator associated with the transaction.
func (t Transaction) GetOriginator(db *DB) (identity *Identity, err error) {
	identity = &Identity{}
	if err = db.Query().Where("id = ?", t.OriginatorID).First(identity).Error; err != nil {
		return nil, err
	}
	return identity, nil
}

// Return the originator address associated with the transaction.
func (t Transaction) GetBeneficiary(db *DB) (identity *Identity, err error) {
	identity = &Identity{}
	if err = db.Query().Where("id = ?", t.BeneficiaryID).First(identity).Error; err != nil {
		return nil, err
	}
	return identity, nil
}

// AmountFloat converts the amount decmial into an exact two precision float32 for
// use with the protocol buffers.
func (t Transaction) AmountFloat() float32 {
	bal, _ := t.Amount.Truncate(2).Float64()
	return float32(bal)
}

// Proto converts the transaction into a protocol buffer transaction
func (t Transaction) Proto() *pb.Transaction {
	return &pb.Transaction{
		Originator: &pb.Account{
			WalletAddress: t.Originator.WalletAddress,
			Email:         t.Originator.Email,
			Provider:      t.Originator.Provider,
		},
		Beneficiary: &pb.Account{
			WalletAddress: t.Beneficiary.WalletAddress,
			Email:         t.Beneficiary.Email,
			Provider:      t.Beneficiary.Provider,
		},
		Amount:     t.AmountFloat(),
		Timestamp:  t.Timestamp.Format(time.RFC3339),
		EnvelopeId: t.Envelope,
		Identity:   t.Identity,
		State:      t.State,
	}
}

// LoadIdentity returns the ivms101.Person for the VASP.
func (v VASP) LoadIdentity() (person *ivms101.Person, err error) {
	if v.IVMS101 == "" {
		return nil, fmt.Errorf("vasp record %d does not have IVMS101 data", v.ID)
	}

	person = new(ivms101.Person)
	if err = jsonpb.UnmarshalString(v.IVMS101, person); err != nil {
		return nil, fmt.Errorf("could not unmarshal identity: %s", err)
	}
	return person, nil
}

// LoadIdentity returns the ivms101.Person for the Account.
func (a Account) LoadIdentity() (person *ivms101.Person, err error) {
	if a.IVMS101 == "" {
		return nil, fmt.Errorf("account record %d does not have IVMS101 data", a.ID)
	}

	person = new(ivms101.Person)
	if err = jsonpb.UnmarshalString(a.IVMS101, person); err != nil {
		return nil, fmt.Errorf("could not unmarshal identity: %s", err)
	}
	return person, nil
}

// OpenDB opens a connection to a database with retries and returns the gorm database
// pointer.
func OpenDB(conf *config.Config) (db *gorm.DB, err error) {
	for i := 0; i < conf.Database.MaxRetries+1; i++ {
		if db, err = gorm.Open(postgres.Open(conf.Database.DSN), &gorm.Config{}); err == nil {
			return db, nil
		}
		time.Sleep(time.Second)
	}

	return db, fmt.Errorf("could not connect to database after %d retries: %s", conf.Database.MaxRetries, err)
}

// MigrateDB the schema based on the models defined above.
func MigrateDB(gdb *gorm.DB) (err error) {
	// Migrate models
	if err = gdb.AutoMigrate(&VASP{}, &Wallet{}, &Account{}, &Transaction{}, &Identity{}); err != nil {
		return err
	}

	return nil
}

// ResetDB resets the database using the JSON fixtures.
func ResetDB(gdb *gorm.DB, fixturesPath string) (err error) {
	var (
		vasps    []VASP
		wallets  []Wallet
		accounts []Account
	)

	// Load the VASP fixtures
	if vasps, err = LoadVASPs(fixturesPath); err != nil {
		return err
	}

	// Load the wallet and account fixtures
	if wallets, accounts, err = LoadWallets(fixturesPath); err != nil {
		return err
	}

	// Reset the database
	if err = gdb.Migrator().DropTable(&VASP{}, &Wallet{}, &Account{}, &Transaction{}, &Identity{}); err != nil {
		return err
	}

	// Migration to create the tables
	if err = MigrateDB(gdb); err != nil {
		return err
	}

	// Insert the VASP fixtures into the database
	if err = gdb.Create(&vasps).Error; err != nil {
		return err
	}

	// Insert the wallet fixtures into the database
	if err = gdb.Create(&wallets).Error; err != nil {
		return err
	}

	// Insert the account fixtures into the database
	if err = gdb.Create(&accounts).Error; err != nil {
		return err
	}

	return nil
}
