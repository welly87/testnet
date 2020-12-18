// Specification Copyright (c) 2020 Joint Working Group on interVASP Messaging Standards
// https://intervasp.org/
// https://intervasp.org/wp-content/uploads/2020/05/IVMS101-interVASP-data-model-standard-issue-1-FINAL.pdf

// Protocol Buffer Specification Copyright (c) 2020 CipherTrace, Inc. https://ciphertrace.com

// Licensed under MIT License

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// NOTE ON THE SPECIFICATION MAPPING
// This protocol buffers specification has applied the Protocol Buffers style guide
// https://developers.google.com/protocol-buffers/docs/style to the ISVM101
// specification to be consistent with other Protocol Buffers specifications and to
// avoid common pitfalls when generating language specific classes.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: ivms101/enum.proto

package ivms101

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Definition: A single value corresponding to the nature of name being adopted.
type NaturalPersonNameTypeCode int32

const (
	// Alias name
	// A name other than the legal name by which a natural person is also known.
	NaturalPersonNameTypeCode_NATURAL_PERSON_NAME_TYPE_CODE_ALIA NaturalPersonNameTypeCode = 0
	// Name at birth
	// The name given to a natural person at birth.
	NaturalPersonNameTypeCode_NATURAL_PERSON_NAME_TYPE_CODE_BIRT NaturalPersonNameTypeCode = 1
	// Maiden name
	// The original name of a natural person who has changed their name after marriage.
	NaturalPersonNameTypeCode_NATURAL_PERSON_NAME_TYPE_CODE_MAID NaturalPersonNameTypeCode = 2
	// Legal name
	// Identifies a natural person for legal, official or administrative purposes.
	NaturalPersonNameTypeCode_NATURAL_PERSON_NAME_TYPE_CODE_LEGL NaturalPersonNameTypeCode = 3
	// Unspecified
	// A name by which a natural person may be known but which cannot otherwise be
	// categorized or the category of which the sender is unable to determine.
	NaturalPersonNameTypeCode_NATURAL_PERSON_NAME_TYPE_CODE_MISC NaturalPersonNameTypeCode = 4
)

// Enum value maps for NaturalPersonNameTypeCode.
var (
	NaturalPersonNameTypeCode_name = map[int32]string{
		0: "NATURAL_PERSON_NAME_TYPE_CODE_ALIA",
		1: "NATURAL_PERSON_NAME_TYPE_CODE_BIRT",
		2: "NATURAL_PERSON_NAME_TYPE_CODE_MAID",
		3: "NATURAL_PERSON_NAME_TYPE_CODE_LEGL",
		4: "NATURAL_PERSON_NAME_TYPE_CODE_MISC",
	}
	NaturalPersonNameTypeCode_value = map[string]int32{
		"NATURAL_PERSON_NAME_TYPE_CODE_ALIA": 0,
		"NATURAL_PERSON_NAME_TYPE_CODE_BIRT": 1,
		"NATURAL_PERSON_NAME_TYPE_CODE_MAID": 2,
		"NATURAL_PERSON_NAME_TYPE_CODE_LEGL": 3,
		"NATURAL_PERSON_NAME_TYPE_CODE_MISC": 4,
	}
)

func (x NaturalPersonNameTypeCode) Enum() *NaturalPersonNameTypeCode {
	p := new(NaturalPersonNameTypeCode)
	*p = x
	return p
}

func (x NaturalPersonNameTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NaturalPersonNameTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ivms101_enum_proto_enumTypes[0].Descriptor()
}

func (NaturalPersonNameTypeCode) Type() protoreflect.EnumType {
	return &file_ivms101_enum_proto_enumTypes[0]
}

func (x NaturalPersonNameTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NaturalPersonNameTypeCode.Descriptor instead.
func (NaturalPersonNameTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_ivms101_enum_proto_rawDescGZIP(), []int{0}
}

//  Definition: A single value corresponding to the nature of name being specified
// for the legal person.
type LegalPersonNameTypeCode int32

const (
	// Legal name
	// Official name under which an organisation is registered.
	LegalPersonNameTypeCode_LEGAL_PERSON_NAME_TYPE_CODE_LEGL LegalPersonNameTypeCode = 0
	// Short name
	// Specifies the short name of the organisation.
	LegalPersonNameTypeCode_LEGAL_PERSON_NAME_TYPE_CODE_SHRT LegalPersonNameTypeCode = 1
	// Trading name
	// Name used by a business for commercial purposes, although its registered legal
	// name, used for contracts and other formal situations, may be another.
	LegalPersonNameTypeCode_LEGAL_PERSON_NAME_TYPE_CODE_TRAD LegalPersonNameTypeCode = 2
)

// Enum value maps for LegalPersonNameTypeCode.
var (
	LegalPersonNameTypeCode_name = map[int32]string{
		0: "LEGAL_PERSON_NAME_TYPE_CODE_LEGL",
		1: "LEGAL_PERSON_NAME_TYPE_CODE_SHRT",
		2: "LEGAL_PERSON_NAME_TYPE_CODE_TRAD",
	}
	LegalPersonNameTypeCode_value = map[string]int32{
		"LEGAL_PERSON_NAME_TYPE_CODE_LEGL": 0,
		"LEGAL_PERSON_NAME_TYPE_CODE_SHRT": 1,
		"LEGAL_PERSON_NAME_TYPE_CODE_TRAD": 2,
	}
)

func (x LegalPersonNameTypeCode) Enum() *LegalPersonNameTypeCode {
	p := new(LegalPersonNameTypeCode)
	*p = x
	return p
}

func (x LegalPersonNameTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LegalPersonNameTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ivms101_enum_proto_enumTypes[1].Descriptor()
}

func (LegalPersonNameTypeCode) Type() protoreflect.EnumType {
	return &file_ivms101_enum_proto_enumTypes[1]
}

func (x LegalPersonNameTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LegalPersonNameTypeCode.Descriptor instead.
func (LegalPersonNameTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_ivms101_enum_proto_rawDescGZIP(), []int{1}
}

// Definition: Identifies the nature of the address.
type AddressTypeCode int32

const (
	// Residential
	// Address is the home address.
	AddressTypeCode_ADDRESS_TYPE_CODE_HOME AddressTypeCode = 0
	// Business
	// Address is the business address.
	AddressTypeCode_ADDRESS_TYPE_CODE_BIZZ AddressTypeCode = 1
	// Geographic
	// Address is the unspecified physical (geographical) address suitable for
	// identification of the natural or legal person.
	AddressTypeCode_ADDRESS_TYPE_CODE_GEOG AddressTypeCode = 2
)

// Enum value maps for AddressTypeCode.
var (
	AddressTypeCode_name = map[int32]string{
		0: "ADDRESS_TYPE_CODE_HOME",
		1: "ADDRESS_TYPE_CODE_BIZZ",
		2: "ADDRESS_TYPE_CODE_GEOG",
	}
	AddressTypeCode_value = map[string]int32{
		"ADDRESS_TYPE_CODE_HOME": 0,
		"ADDRESS_TYPE_CODE_BIZZ": 1,
		"ADDRESS_TYPE_CODE_GEOG": 2,
	}
)

func (x AddressTypeCode) Enum() *AddressTypeCode {
	p := new(AddressTypeCode)
	*p = x
	return p
}

func (x AddressTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddressTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ivms101_enum_proto_enumTypes[2].Descriptor()
}

func (AddressTypeCode) Type() protoreflect.EnumType {
	return &file_ivms101_enum_proto_enumTypes[2]
}

func (x AddressTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddressTypeCode.Descriptor instead.
func (AddressTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_ivms101_enum_proto_rawDescGZIP(), []int{2}
}

// Definition: Identifies the national identification type.
// NationalIdentifierTypeCode applies a restriction over the codes present in ISO20022
// datatype ‘TypeOfIdentification4Code’.
type NationalIdentifierTypeCode int32

const (
	// Alien registration number
	// Number assigned by a government agency to identify foreign nationals.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_ARNU NationalIdentifierTypeCode = 0
	// Passport number
	// Number assigned by a passport authority.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_CCPT NationalIdentifierTypeCode = 1
	// Registration authority identifier
	// Identifier of a legal entity as maintained by a registration authority.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_RAID NationalIdentifierTypeCode = 2
	// Driver license number
	// Number assigned to a driver's license.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_DRLC NationalIdentifierTypeCode = 3
	// Foreign investment identity number
	// Number assigned to a foreign investor (other than the alien number).
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_FIIN NationalIdentifierTypeCode = 4
	// Tax identification number
	// Number assigned by a tax authority to an entity.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_TXID NationalIdentifierTypeCode = 5
	// Social security number
	// Number assigned by a social security agency.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_SOCS NationalIdentifierTypeCode = 6
	// Identity card number
	// Number assigned by a national authority to an identity card.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_IDCD NationalIdentifierTypeCode = 7
	// Legal Entity Identifier
	// Legal Entity Identifier (LEI) assigned in accordance with ISO 17442.
	// The LEI is a 20-character, alpha-numeric code that enables clear and unique
	// identification of legal entities participating in financial transactions.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_LEIX NationalIdentifierTypeCode = 8
	// Unspecified
	// A national identifier which may be known but which cannot otherwise be
	// categorized or the category of which the sender is unable to determine.
	NationalIdentifierTypeCode_NATIONAL_IDENTIFIER_TYPE_CODE_MISC NationalIdentifierTypeCode = 9
)

// Enum value maps for NationalIdentifierTypeCode.
var (
	NationalIdentifierTypeCode_name = map[int32]string{
		0: "NATIONAL_IDENTIFIER_TYPE_CODE_ARNU",
		1: "NATIONAL_IDENTIFIER_TYPE_CODE_CCPT",
		2: "NATIONAL_IDENTIFIER_TYPE_CODE_RAID",
		3: "NATIONAL_IDENTIFIER_TYPE_CODE_DRLC",
		4: "NATIONAL_IDENTIFIER_TYPE_CODE_FIIN",
		5: "NATIONAL_IDENTIFIER_TYPE_CODE_TXID",
		6: "NATIONAL_IDENTIFIER_TYPE_CODE_SOCS",
		7: "NATIONAL_IDENTIFIER_TYPE_CODE_IDCD",
		8: "NATIONAL_IDENTIFIER_TYPE_CODE_LEIX",
		9: "NATIONAL_IDENTIFIER_TYPE_CODE_MISC",
	}
	NationalIdentifierTypeCode_value = map[string]int32{
		"NATIONAL_IDENTIFIER_TYPE_CODE_ARNU": 0,
		"NATIONAL_IDENTIFIER_TYPE_CODE_CCPT": 1,
		"NATIONAL_IDENTIFIER_TYPE_CODE_RAID": 2,
		"NATIONAL_IDENTIFIER_TYPE_CODE_DRLC": 3,
		"NATIONAL_IDENTIFIER_TYPE_CODE_FIIN": 4,
		"NATIONAL_IDENTIFIER_TYPE_CODE_TXID": 5,
		"NATIONAL_IDENTIFIER_TYPE_CODE_SOCS": 6,
		"NATIONAL_IDENTIFIER_TYPE_CODE_IDCD": 7,
		"NATIONAL_IDENTIFIER_TYPE_CODE_LEIX": 8,
		"NATIONAL_IDENTIFIER_TYPE_CODE_MISC": 9,
	}
)

func (x NationalIdentifierTypeCode) Enum() *NationalIdentifierTypeCode {
	p := new(NationalIdentifierTypeCode)
	*p = x
	return p
}

func (x NationalIdentifierTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NationalIdentifierTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ivms101_enum_proto_enumTypes[3].Descriptor()
}

func (NationalIdentifierTypeCode) Type() protoreflect.EnumType {
	return &file_ivms101_enum_proto_enumTypes[3]
}

func (x NationalIdentifierTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NationalIdentifierTypeCode.Descriptor instead.
func (NationalIdentifierTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_ivms101_enum_proto_rawDescGZIP(), []int{3}
}

// Definition: Identifies the national script from which transliteration to Latin
// script is applied.
type TransliterationMethodCode int32

const (
	// Arabic (Arabic language)
	// ISO 233-2:1993
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_ARAB TransliterationMethodCode = 0
	// Arabic (Persian language)
	// ISO 233-3:1999
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_ARAN TransliterationMethodCode = 1
	// Armenian
	// ISO 9985:1996
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_ARMN TransliterationMethodCode = 2
	// Cyrillic
	// ISO 9:1995
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_CYRL TransliterationMethodCode = 3
	// Devanagari & related Indic
	// ISO 15919:2001
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_DEVA TransliterationMethodCode = 4
	// Georgian
	// ISO 9984:1996
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_GEOR TransliterationMethodCode = 5
	// Greek
	// ISO 843:1997
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_GREK TransliterationMethodCode = 6
	// Han (Hanzi, Kanji, Hanja)
	// ISO 7098:2015
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_HANI TransliterationMethodCode = 7
	// Hebrew
	// ISO 259-2:1994
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_HEBR TransliterationMethodCode = 8
	// Kana
	// ISO 3602:1989
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_KANA TransliterationMethodCode = 10
	// Korean
	// Revised Romanization of Korean
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_KORE TransliterationMethodCode = 11
	// Thai
	// ISO 11940-2:2007
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_THAI TransliterationMethodCode = 12
	// Script other than those listed above
	// Unspecified Standard
	TransliterationMethodCode_TRANSLITERATION_METHOD_CODE_OTHR TransliterationMethodCode = 13
)

// Enum value maps for TransliterationMethodCode.
var (
	TransliterationMethodCode_name = map[int32]string{
		0:  "TRANSLITERATION_METHOD_CODE_ARAB",
		1:  "TRANSLITERATION_METHOD_CODE_ARAN",
		2:  "TRANSLITERATION_METHOD_CODE_ARMN",
		3:  "TRANSLITERATION_METHOD_CODE_CYRL",
		4:  "TRANSLITERATION_METHOD_CODE_DEVA",
		5:  "TRANSLITERATION_METHOD_CODE_GEOR",
		6:  "TRANSLITERATION_METHOD_CODE_GREK",
		7:  "TRANSLITERATION_METHOD_CODE_HANI",
		8:  "TRANSLITERATION_METHOD_CODE_HEBR",
		10: "TRANSLITERATION_METHOD_CODE_KANA",
		11: "TRANSLITERATION_METHOD_CODE_KORE",
		12: "TRANSLITERATION_METHOD_CODE_THAI",
		13: "TRANSLITERATION_METHOD_CODE_OTHR",
	}
	TransliterationMethodCode_value = map[string]int32{
		"TRANSLITERATION_METHOD_CODE_ARAB": 0,
		"TRANSLITERATION_METHOD_CODE_ARAN": 1,
		"TRANSLITERATION_METHOD_CODE_ARMN": 2,
		"TRANSLITERATION_METHOD_CODE_CYRL": 3,
		"TRANSLITERATION_METHOD_CODE_DEVA": 4,
		"TRANSLITERATION_METHOD_CODE_GEOR": 5,
		"TRANSLITERATION_METHOD_CODE_GREK": 6,
		"TRANSLITERATION_METHOD_CODE_HANI": 7,
		"TRANSLITERATION_METHOD_CODE_HEBR": 8,
		"TRANSLITERATION_METHOD_CODE_KANA": 10,
		"TRANSLITERATION_METHOD_CODE_KORE": 11,
		"TRANSLITERATION_METHOD_CODE_THAI": 12,
		"TRANSLITERATION_METHOD_CODE_OTHR": 13,
	}
)

func (x TransliterationMethodCode) Enum() *TransliterationMethodCode {
	p := new(TransliterationMethodCode)
	*p = x
	return p
}

func (x TransliterationMethodCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransliterationMethodCode) Descriptor() protoreflect.EnumDescriptor {
	return file_ivms101_enum_proto_enumTypes[4].Descriptor()
}

func (TransliterationMethodCode) Type() protoreflect.EnumType {
	return &file_ivms101_enum_proto_enumTypes[4]
}

func (x TransliterationMethodCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransliterationMethodCode.Descriptor instead.
func (TransliterationMethodCode) EnumDescriptor() ([]byte, []int) {
	return file_ivms101_enum_proto_rawDescGZIP(), []int{4}
}

var File_ivms101_enum_proto protoreflect.FileDescriptor

var file_ivms101_enum_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x76, 0x6d, 0x73, 0x31, 0x30, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x76, 0x6d, 0x73, 0x31, 0x30, 0x31, 0x2a, 0xe3, 0x01,
	0x0a, 0x19, 0x4e, 0x61, 0x74, 0x75, 0x72, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x4e,
	0x41, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4c, 0x49,
	0x41, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x50,
	0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x49, 0x52, 0x54, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x4e,
	0x41, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x41, 0x49,
	0x44, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x50,
	0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4c, 0x45, 0x47, 0x4c, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x4e,
	0x41, 0x54, 0x55, 0x52, 0x41, 0x4c, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x49, 0x53,
	0x43, 0x10, 0x04, 0x2a, 0x8b, 0x01, 0x0a, 0x17, 0x4c, 0x65, 0x67, 0x61, 0x6c, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x24, 0x0a, 0x20, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4c,
	0x45, 0x47, 0x4c, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x50,
	0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x48, 0x52, 0x54, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x4c,
	0x45, 0x47, 0x41, 0x4c, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x44, 0x10,
	0x02, 0x2a, 0x65, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x49, 0x5a, 0x5a, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x41, 0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x47, 0x45, 0x4f, 0x47, 0x10, 0x02, 0x2a, 0xac, 0x03, 0x0a, 0x1a, 0x4e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x52, 0x4e, 0x55, 0x10, 0x00, 0x12,
	0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x43, 0x43, 0x50, 0x54, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x41, 0x49, 0x44, 0x10, 0x02, 0x12,
	0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x44, 0x52, 0x4c, 0x43, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x49, 0x49, 0x4e, 0x10, 0x04, 0x12,
	0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x54, 0x58, 0x49, 0x44, 0x10, 0x05, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x4f, 0x43, 0x53, 0x10, 0x06, 0x12,
	0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x49, 0x44, 0x43, 0x44, 0x10, 0x07, 0x12, 0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4c, 0x45, 0x49, 0x58, 0x10, 0x08, 0x12,
	0x26, 0x0a, 0x22, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x49, 0x44, 0x45, 0x4e,
	0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x4d, 0x49, 0x53, 0x43, 0x10, 0x09, 0x2a, 0x89, 0x04, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49,
	0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x52, 0x41, 0x42, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x52, 0x41, 0x4e, 0x10,
	0x01, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x41, 0x52, 0x4d, 0x4e, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x59, 0x52, 0x4c, 0x10, 0x03, 0x12, 0x24, 0x0a,
	0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x56,
	0x41, 0x10, 0x04, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x47, 0x45, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54,
	0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x52, 0x45, 0x4b, 0x10, 0x06, 0x12,
	0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x48,
	0x41, 0x4e, 0x49, 0x10, 0x07, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49,
	0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x45, 0x42, 0x52, 0x10, 0x08, 0x12, 0x24, 0x0a, 0x20, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4b, 0x41, 0x4e, 0x41, 0x10,
	0x0a, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x4b, 0x4f, 0x52, 0x45, 0x10, 0x0b, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x48, 0x41, 0x49, 0x10, 0x0c, 0x12, 0x24, 0x0a,
	0x20, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x4c, 0x49, 0x54, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x54, 0x48,
	0x52, 0x10, 0x0d, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x72, 0x69, 0x73, 0x61, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x76, 0x6d, 0x73, 0x31, 0x30,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ivms101_enum_proto_rawDescOnce sync.Once
	file_ivms101_enum_proto_rawDescData = file_ivms101_enum_proto_rawDesc
)

func file_ivms101_enum_proto_rawDescGZIP() []byte {
	file_ivms101_enum_proto_rawDescOnce.Do(func() {
		file_ivms101_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_ivms101_enum_proto_rawDescData)
	})
	return file_ivms101_enum_proto_rawDescData
}

var file_ivms101_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_ivms101_enum_proto_goTypes = []interface{}{
	(NaturalPersonNameTypeCode)(0),  // 0: ivms101.NaturalPersonNameTypeCode
	(LegalPersonNameTypeCode)(0),    // 1: ivms101.LegalPersonNameTypeCode
	(AddressTypeCode)(0),            // 2: ivms101.AddressTypeCode
	(NationalIdentifierTypeCode)(0), // 3: ivms101.NationalIdentifierTypeCode
	(TransliterationMethodCode)(0),  // 4: ivms101.TransliterationMethodCode
}
var file_ivms101_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ivms101_enum_proto_init() }
func file_ivms101_enum_proto_init() {
	if File_ivms101_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ivms101_enum_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ivms101_enum_proto_goTypes,
		DependencyIndexes: file_ivms101_enum_proto_depIdxs,
		EnumInfos:         file_ivms101_enum_proto_enumTypes,
	}.Build()
	File_ivms101_enum_proto = out.File
	file_ivms101_enum_proto_rawDesc = nil
	file_ivms101_enum_proto_goTypes = nil
	file_ivms101_enum_proto_depIdxs = nil
}
