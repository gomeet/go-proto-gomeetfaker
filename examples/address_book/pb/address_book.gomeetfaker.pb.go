// Code generated by protoc-gen-gogo. DO NOT EDIT.
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

import faker "github.com/dmgk/faker"
import locales "github.com/dmgk/faker/locales"
import rand "math/rand"
import time "time"
import uuid "github.com/google/uuid"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gomeet/go-proto-gomeetfaker"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func GomeetFakerRand() *rand.Rand {
	seed := time.Now().UnixNano()
	return rand.New(rand.NewSource(seed))
}
func init() {
	GomeetFakerSetLocale("en")
}

func GomeetFakerSetLocale(l string) {
	switch l {
	case "es":
		faker.Locale = locales.Es
	case "zh-tw":
		faker.Locale = locales.Zh_TW
	case "en-gb":
		faker.Locale = locales.En_GB
	case "en-nep":
		faker.Locale = locales.En_NEP
	case "sk":
		faker.Locale = locales.Sk
	case "fa":
		faker.Locale = locales.Fa
	case "nb-no":
		faker.Locale = locales.Nb_NO
	case "de-at":
		faker.Locale = locales.De_AT
	case "de-ch":
		faker.Locale = locales.De_CH
	case "en-ind":
		faker.Locale = locales.En_IND
	case "pl":
		faker.Locale = locales.Pl
	case "pt-br":
		faker.Locale = locales.Pt_BR
	case "en-au":
		faker.Locale = locales.En_AU
	case "en-bork":
		faker.Locale = locales.En_BORK
	case "fr":
		faker.Locale = locales.Fr
	case "zh-cn":
		faker.Locale = locales.Zh_CN
	case "en-au-ocker":
		faker.Locale = locales.En_AU_OCKER
	case "en":
		faker.Locale = locales.En
	case "ja":
		faker.Locale = locales.Ja
	case "ru":
		faker.Locale = locales.Ru
	case "vi":
		faker.Locale = locales.Vi
	case "de":
		faker.Locale = locales.De
	case "en-ca":
		faker.Locale = locales.En_CA
	case "ko":
		faker.Locale = locales.Ko
	case "nl":
		faker.Locale = locales.Nl
	case "sv":
		faker.Locale = locales.Sv
	case "en-us":
		faker.Locale = locales.En_US
	case "it":
		faker.Locale = locales.It
	default:
		faker.Locale = locales.En
	}
}
func NewAddressBookGomeetFaker() *AddressBook {
	this := &AddressBook{}
	this.Uuid = uuid.New().String()
	for i := 0; i < 4; i++ {
		aCurrentContacts := NewContactNfoGomeetFaker()
		this.Contacts = append(this.Contacts, aCurrentContacts)
	}
	return this
}

func NewContactNfoGomeetFaker() *ContactNfo {
	this := &ContactNfo{}
	this.Uuid = uuid.New().String()
	this.Name = faker.Name().Name()
	for i := 0; i < 4; i++ {
		aCurrentAvatarsUrl := faker.Avatar().Url("jpg", 300, 300)
		this.AvatarsUrl = append(this.AvatarsUrl, aCurrentAvatarsUrl)
	}
	for i := 0; i < 4; i++ {
		aCurrentPhones := NewPhoneGomeetFaker()
		this.Phones = append(this.Phones, aCurrentPhones)
	}
	for i := 0; i < 4; i++ {
		aCurrentAddresses := NewAddressGomeetFaker()
		this.Addresses = append(this.Addresses, aCurrentAddresses)
	}
	this.Company = NewCompanyGomeetFaker()
	for i := 0; i < 2; i++ {
		aCurrentNotes := faker.Hacker().SaySomethingSmart()
		this.Notes = append(this.Notes, aCurrentNotes)
	}
	aBirthdayDate := faker.Date().Birthday(17, 99)
	this.Birthday = aBirthdayDate.Format("2006-01-02")
	return this
}

func NewPhoneGomeetFaker() *Phone {
	this := &Phone{}
	this.Uuid = uuid.New().String()
	this.Type = Phone_PhoneType([]int32{1, 2, 3}[GomeetFakerRand().Intn(3)])
	this.PhoneNumber = faker.PhoneNumber().PhoneNumber()
	return this
}

func NewAddressGomeetFaker() *Address {
	this := &Address{}
	this.Uuid = uuid.New().String()
	this.Type = Address_AddressType([]int32{1, 2, 3}[GomeetFakerRand().Intn(3)])
	this.City = faker.Address().City()
	this.StreetAddress = faker.Address().StreetAddress()
	this.Postcode = faker.Address().Postcode()
	return this
}

func NewCompanyGomeetFaker() *Company {
	this := &Company{}
	this.Uuid = uuid.New().String()
	this.Name = faker.Company().Name()
	this.Suffix = faker.Company().Suffix()
	this.CatchPhrase = faker.Company().CatchPhrase()
	this.Bs = faker.Company().Bs()
	this.Ein = faker.Company().Ein()
	this.DunsNumber = faker.Company().DunsNumber()
	this.Logo = faker.Company().Logo()
	for i := 0; i < 2; i++ {
		aCurrentAddresses := NewAddressGomeetFaker()
		this.Addresses = append(this.Addresses, aCurrentAddresses)
	}
	this.Note = faker.Hacker().SaySomethingSmart()
	return this
}
