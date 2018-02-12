// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/address_book/pb/address_book.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	examples/address_book/pb/address_book.proto

It has these top-level messages:
	AddressBook
	ContactNfo
	Phone
	Address
	Company
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gomeet/go-proto-gomeetfaker"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Phone_PhoneType int32

const (
	Phone_UNKNOW Phone_PhoneType = 0
	Phone_HOME   Phone_PhoneType = 1
	Phone_WORK   Phone_PhoneType = 2
	Phone_OTHER  Phone_PhoneType = 3
)

var Phone_PhoneType_name = map[int32]string{
	0: "UNKNOW",
	1: "HOME",
	2: "WORK",
	3: "OTHER",
}
var Phone_PhoneType_value = map[string]int32{
	"UNKNOW": 0,
	"HOME":   1,
	"WORK":   2,
	"OTHER":  3,
}

func (x Phone_PhoneType) String() string {
	return proto.EnumName(Phone_PhoneType_name, int32(x))
}
func (Phone_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Address_AddressType int32

const (
	Address_UNKNOW Address_AddressType = 0
	Address_HOME   Address_AddressType = 1
	Address_WORK   Address_AddressType = 2
	Address_OTHER  Address_AddressType = 3
)

var Address_AddressType_name = map[int32]string{
	0: "UNKNOW",
	1: "HOME",
	2: "WORK",
	3: "OTHER",
}
var Address_AddressType_value = map[string]int32{
	"UNKNOW": 0,
	"HOME":   1,
	"WORK":   2,
	"OTHER":  3,
}

func (x Address_AddressType) String() string {
	return proto.EnumName(Address_AddressType_name, int32(x))
}
func (Address_AddressType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type AddressBook struct {
	Uuid     string        `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Contacts []*ContactNfo `protobuf:"bytes,2,rep,name=contacts" json:"contacts,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddressBook) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *AddressBook) GetContacts() []*ContactNfo {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type ContactNfo struct {
	Uuid       string     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name       string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	AvatarsUrl []string   `protobuf:"bytes,3,rep,name=avatars_url,json=avatarsUrl" json:"avatars_url,omitempty"`
	Phones     []*Phone   `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	Addresses  []*Address `protobuf:"bytes,5,rep,name=addresses" json:"addresses,omitempty"`
	Company    *Company   `protobuf:"bytes,6,opt,name=company" json:"company,omitempty"`
	Notes      []string   `protobuf:"bytes,7,rep,name=notes" json:"notes,omitempty"`
	Birthday   string     `protobuf:"bytes,8,opt,name=birthday" json:"birthday,omitempty"`
}

func (m *ContactNfo) Reset()                    { *m = ContactNfo{} }
func (m *ContactNfo) String() string            { return proto.CompactTextString(m) }
func (*ContactNfo) ProtoMessage()               {}
func (*ContactNfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ContactNfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ContactNfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContactNfo) GetAvatarsUrl() []string {
	if m != nil {
		return m.AvatarsUrl
	}
	return nil
}

func (m *ContactNfo) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *ContactNfo) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ContactNfo) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *ContactNfo) GetNotes() []string {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *ContactNfo) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

type Phone struct {
	Uuid        string          `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Type        Phone_PhoneType `protobuf:"varint,2,opt,name=type,enum=gomeetfaker.examples.Phone_PhoneType" json:"type,omitempty"`
	PhoneNumber string          `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
}

func (m *Phone) Reset()                    { *m = Phone{} }
func (m *Phone) String() string            { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()               {}
func (*Phone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Phone) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Phone) GetType() Phone_PhoneType {
	if m != nil {
		return m.Type
	}
	return Phone_UNKNOW
}

func (m *Phone) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type Address struct {
	Uuid          string              `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Type          Address_AddressType `protobuf:"varint,2,opt,name=type,enum=gomeetfaker.examples.Address_AddressType" json:"type,omitempty"`
	City          string              `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	StreetAddress string              `protobuf:"bytes,4,opt,name=street_address,json=streetAddress" json:"street_address,omitempty"`
	Postcode      string              `protobuf:"bytes,5,opt,name=postcode" json:"postcode,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Address) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Address) GetType() Address_AddressType {
	if m != nil {
		return m.Type
	}
	return Address_UNKNOW
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetStreetAddress() string {
	if m != nil {
		return m.StreetAddress
	}
	return ""
}

func (m *Address) GetPostcode() string {
	if m != nil {
		return m.Postcode
	}
	return ""
}

type Company struct {
	Uuid        string     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name        string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Suffix      string     `protobuf:"bytes,3,opt,name=suffix" json:"suffix,omitempty"`
	CatchPhrase string     `protobuf:"bytes,4,opt,name=catch_phrase,json=catchPhrase" json:"catch_phrase,omitempty"`
	Bs          string     `protobuf:"bytes,5,opt,name=bs" json:"bs,omitempty"`
	Ein         string     `protobuf:"bytes,6,opt,name=ein" json:"ein,omitempty"`
	DunsNumber  string     `protobuf:"bytes,7,opt,name=duns_number,json=dunsNumber" json:"duns_number,omitempty"`
	Logo        string     `protobuf:"bytes,8,opt,name=logo" json:"logo,omitempty"`
	Addresses   []*Address `protobuf:"bytes,9,rep,name=addresses" json:"addresses,omitempty"`
	Note        string     `protobuf:"bytes,10,opt,name=note" json:"note,omitempty"`
}

func (m *Company) Reset()                    { *m = Company{} }
func (m *Company) String() string            { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()               {}
func (*Company) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Company) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Company) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Company) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *Company) GetCatchPhrase() string {
	if m != nil {
		return m.CatchPhrase
	}
	return ""
}

func (m *Company) GetBs() string {
	if m != nil {
		return m.Bs
	}
	return ""
}

func (m *Company) GetEin() string {
	if m != nil {
		return m.Ein
	}
	return ""
}

func (m *Company) GetDunsNumber() string {
	if m != nil {
		return m.DunsNumber
	}
	return ""
}

func (m *Company) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *Company) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Company) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func init() {
	proto.RegisterType((*AddressBook)(nil), "gomeetfaker.examples.AddressBook")
	proto.RegisterType((*ContactNfo)(nil), "gomeetfaker.examples.ContactNfo")
	proto.RegisterType((*Phone)(nil), "gomeetfaker.examples.Phone")
	proto.RegisterType((*Address)(nil), "gomeetfaker.examples.Address")
	proto.RegisterType((*Company)(nil), "gomeetfaker.examples.Company")
	proto.RegisterEnum("gomeetfaker.examples.Phone_PhoneType", Phone_PhoneType_name, Phone_PhoneType_value)
	proto.RegisterEnum("gomeetfaker.examples.Address_AddressType", Address_AddressType_name, Address_AddressType_value)
}

func init() { proto.RegisterFile("examples/address_book/pb/address_book.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0xbe, 0x33, 0x36, 0xc6, 0x1c, 0x72, 0x73, 0xc9, 0xe8, 0x5e, 0x69, 0x94, 0x9b, 0x5c, 0x59,
	0xe8, 0x56, 0x22, 0x4a, 0x03, 0x15, 0xad, 0x52, 0x16, 0x5d, 0xb4, 0xa4, 0x51, 0x23, 0x45, 0x85,
	0x68, 0x94, 0x34, 0x12, 0x1b, 0x64, 0x9b, 0xe1, 0xa7, 0x01, 0x8f, 0x65, 0x0f, 0x55, 0x78, 0x85,
	0xaa, 0x4f, 0xd0, 0x5d, 0xf7, 0x5e, 0xf4, 0x39, 0x58, 0xf6, 0x6d, 0xba, 0xab, 0x3c, 0x36, 0x60,
	0x24, 0x1a, 0x94, 0x0d, 0xd8, 0xe7, 0xfb, 0xce, 0x99, 0xf9, 0xbe, 0x73, 0x8e, 0xe1, 0x98, 0xdf,
	0xdb, 0x13, 0x7f, 0xcc, 0xc3, 0x9a, 0xdd, 0xeb, 0x05, 0x3c, 0x0c, 0xbb, 0x8e, 0x10, 0x77, 0x35,
	0xdf, 0x59, 0x7b, 0xaf, 0xfa, 0x81, 0x90, 0x82, 0xfc, 0x3d, 0x10, 0x13, 0xce, 0x65, 0xdf, 0xbe,
	0xe3, 0x41, 0x75, 0x91, 0xb8, 0xdf, 0x18, 0x8c, 0xe4, 0x70, 0xea, 0x54, 0x5d, 0x31, 0xa9, 0x25,
	0x84, 0xda, 0x40, 0x9c, 0xa8, 0x8c, 0x93, 0x4c, 0x42, 0x2d, 0x9b, 0xac, 0xd0, 0xb2, 0x84, 0xe2,
	0x9b, 0xe4, 0x94, 0xa6, 0x10, 0x77, 0xe4, 0x3f, 0xd0, 0xa7, 0xd3, 0x51, 0x8f, 0x22, 0x0b, 0x55,
	0x0a, 0x4d, 0x98, 0x47, 0x86, 0x41, 0x74, 0x82, 0x3f, 0xbc, 0x60, 0x2a, 0x4e, 0xce, 0xc1, 0x74,
	0x85, 0x27, 0x6d, 0x57, 0x86, 0x14, 0x5b, 0x5a, 0xa5, 0x58, 0xb7, 0xaa, 0x9b, 0x6e, 0x54, 0x3d,
	0x4b, 0x58, 0xad, 0xbe, 0x68, 0xe6, 0xe7, 0x91, 0xa1, 0x7d, 0x47, 0x3a, 0x5b, 0xa6, 0x96, 0xbf,
	0x68, 0x00, 0x2b, 0xc6, 0xd6, 0x53, 0x0f, 0x41, 0xf7, 0xec, 0x09, 0xa7, 0x58, 0xe1, 0x85, 0x79,
	0x64, 0xe4, 0x3e, 0x23, 0x5c, 0x42, 0x4c, 0x85, 0xc9, 0x29, 0x14, 0xed, 0x4f, 0xb6, 0xb4, 0x83,
	0xb0, 0x3b, 0x0d, 0xc6, 0x54, 0xb3, 0xb4, 0x4a, 0xa1, 0xf9, 0xcf, 0x3c, 0x32, 0xf6, 0xea, 0x7f,
	0x99, 0x3a, 0x29, 0x82, 0xf6, 0xd1, 0x1f, 0x94, 0x22, 0x4c, 0x23, 0xcc, 0x20, 0x65, 0xde, 0x04,
	0x63, 0xf2, 0x0a, 0x0c, 0x7f, 0x28, 0x3c, 0x1e, 0x52, 0x5d, 0x49, 0xf9, 0x77, 0xb3, 0x94, 0xab,
	0x98, 0xb3, 0x52, 0x91, 0xe6, 0x90, 0xb7, 0x50, 0x48, 0xfb, 0xc3, 0x43, 0x9a, 0x53, 0x05, 0x0e,
	0x37, 0x17, 0x58, 0x18, 0xbc, 0x2c, 0xb1, 0x4a, 0x24, 0x2f, 0x21, 0xef, 0x8a, 0x89, 0x6f, 0x7b,
	0x33, 0x6a, 0x58, 0xe8, 0xf7, 0x35, 0xce, 0x12, 0x12, 0x5b, 0xb0, 0x89, 0x05, 0x39, 0x4f, 0x48,
	0x1e, 0xd2, 0xbc, 0x92, 0xab, 0x4c, 0x0b, 0x74, 0x33, 0x76, 0x25, 0x01, 0x48, 0x05, 0x4c, 0x67,
	0x14, 0xc8, 0x61, 0xcf, 0x9e, 0x51, 0x53, 0x39, 0xb7, 0x33, 0x8f, 0x0c, 0xd3, 0x31, 0xea, 0xba,
	0xb9, 0x57, 0x72, 0xd9, 0x12, 0x2d, 0xff, 0x40, 0x90, 0x53, 0x2a, 0xb7, 0x76, 0xe2, 0x0c, 0x74,
	0x39, 0xf3, 0x93, 0x4e, 0xec, 0xd6, 0x9f, 0x3c, 0x60, 0x58, 0xf2, 0x7b, 0x3d, 0xf3, 0x79, 0xd3,
	0x9c, 0x47, 0x86, 0xbe, 0x8f, 0x1b, 0x88, 0xa9, 0x64, 0xf2, 0x14, 0x76, 0x94, 0x87, 0x5d, 0x6f,
	0x3a, 0x71, 0x78, 0x40, 0xb5, 0x55, 0x5b, 0xbf, 0xaa, 0xb6, 0x16, 0x15, 0xdc, 0x52, 0x68, 0xf9,
	0x14, 0x0a, 0xcb, 0x52, 0x04, 0xc0, 0xb8, 0x69, 0x5d, 0xb6, 0xda, 0xb7, 0xa5, 0x3f, 0x88, 0x09,
	0xfa, 0x45, 0xfb, 0xfd, 0x79, 0x09, 0xc5, 0x4f, 0xb7, 0x6d, 0x76, 0x59, 0xc2, 0xa4, 0x00, 0xb9,
	0xf6, 0xf5, 0xc5, 0x39, 0x2b, 0x69, 0xe5, 0x6f, 0x18, 0xf2, 0xa9, 0xf3, 0x5b, 0x65, 0xbd, 0x5b,
	0x93, 0x75, 0xf4, 0x60, 0x1b, 0x17, 0xff, 0x1b, 0xa5, 0x1d, 0x80, 0xee, 0x8e, 0xe4, 0x2c, 0x95,
	0xa4, 0xd0, 0xb2, 0x1a, 0xd4, 0x38, 0x4a, 0x6a, 0xb0, 0x1b, 0xca, 0x80, 0x73, 0xd9, 0x4d, 0x07,
	0x80, 0xea, 0x59, 0x9e, 0x85, 0xd8, 0x9f, 0x09, 0xbe, 0xb8, 0xf7, 0xff, 0x60, 0xfa, 0x22, 0x94,
	0xae, 0xe8, 0x71, 0x9a, 0xcb, 0x52, 0x1b, 0x88, 0x2d, 0x91, 0x72, 0x63, 0xb9, 0xc3, 0x8f, 0xf5,
	0xe8, 0x27, 0x86, 0x7c, 0x3a, 0x59, 0x5b, 0x3d, 0x3a, 0x58, 0x5b, 0x42, 0x75, 0x8f, 0xce, 0x6a,
	0x07, 0x2d, 0x30, 0xc2, 0x69, 0xbf, 0x3f, 0xba, 0xcf, 0x4a, 0xef, 0x60, 0x8a, 0x58, 0x1a, 0x27,
	0xc7, 0xb0, 0xe3, 0xda, 0xd2, 0x1d, 0x76, 0xfd, 0x61, 0x60, 0x87, 0x3c, 0x2b, 0xbd, 0x13, 0x4b,
	0x2f, 0x2a, 0xf4, 0x4a, 0x81, 0x84, 0x02, 0x76, 0xc2, 0xac, 0xe4, 0x0e, 0xae, 0x20, 0x86, 0x9d,
	0x90, 0xec, 0x83, 0xc6, 0x47, 0x9e, 0x5a, 0x96, 0x25, 0xf4, 0x0c, 0xb1, 0x38, 0x48, 0x8e, 0xa0,
	0xd8, 0x9b, 0x7a, 0xe1, 0x62, 0xae, 0xf2, 0x59, 0x4e, 0x03, 0x31, 0x88, 0xc1, 0x64, 0xaa, 0x62,
	0x35, 0x63, 0x31, 0x10, 0xe9, 0x62, 0xa4, 0x9c, 0xd7, 0x88, 0xa9, 0xe8, 0xfa, 0x6e, 0x17, 0x1e,
	0xb1, 0xdb, 0x38, 0xbb, 0xdb, 0xb1, 0x63, 0x42, 0x72, 0x0a, 0xab, 0x33, 0x82, 0xc4, 0x31, 0x21,
	0x79, 0x53, 0xef, 0x60, 0xdf, 0x71, 0x0c, 0xf5, 0x19, 0x7e, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0x35, 0x3b, 0x7f, 0x05, 0x06, 0x00, 0x00,
}
