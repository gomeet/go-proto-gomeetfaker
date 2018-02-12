// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/book/pb/book.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	examples/book/pb/book.proto

It has these top-level messages:
	Book
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
	case "de-at":
		faker.Locale = locales.De_AT
	case "en-nep":
		faker.Locale = locales.En_NEP
	case "en-us":
		faker.Locale = locales.En_US
	case "fa":
		faker.Locale = locales.Fa
	case "ja":
		faker.Locale = locales.Ja
	case "ru":
		faker.Locale = locales.Ru
	case "zh-cn":
		faker.Locale = locales.Zh_CN
	case "en-gb":
		faker.Locale = locales.En_GB
	case "en-ind":
		faker.Locale = locales.En_IND
	case "zh-tw":
		faker.Locale = locales.Zh_TW
	case "pl":
		faker.Locale = locales.Pl
	case "sv":
		faker.Locale = locales.Sv
	case "it":
		faker.Locale = locales.It
	case "vi":
		faker.Locale = locales.Vi
	case "en":
		faker.Locale = locales.En
	case "es":
		faker.Locale = locales.Es
	case "de-ch":
		faker.Locale = locales.De_CH
	case "en-bork":
		faker.Locale = locales.En_BORK
	case "nl":
		faker.Locale = locales.Nl
	case "pt-br":
		faker.Locale = locales.Pt_BR
	case "en-ca":
		faker.Locale = locales.En_CA
	case "fr":
		faker.Locale = locales.Fr
	case "en-au-ocker":
		faker.Locale = locales.En_AU_OCKER
	case "ko":
		faker.Locale = locales.Ko
	case "nb-no":
		faker.Locale = locales.Nb_NO
	case "sk":
		faker.Locale = locales.Sk
	case "de":
		faker.Locale = locales.De
	case "en-au":
		faker.Locale = locales.En_AU
	default:
		faker.Locale = locales.En
	}
}
func NewBookGomeetFaker() *Book {
	this := &Book{}
	this.Uuid = uuid.New().String()
	this.Author = faker.Name().Name()
	this.Title = faker.Lorem().String()
	this.Isbn10 = faker.Code().Isbn10()
	this.Isbn13 = faker.Code().Isbn13()
	return this
}