package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/gomeet/go-proto-gomeetfaker/examples/address_book/pb"
)

var (
	locale string
)

func init() {
	flag.StringVar(&locale, "locale", "en", fmt.Sprintf("locales [%s]", allLocalesParams()))
}

func allLocalesParams() string {
	return strings.Join([]string{"de-at", "de-ch", "de", "en-au", "en-bork", "en-ca", "en-gb",
		"en-ind", "en-nep", "en-us", "en-au-ocker", "en", "es", "fa",
		"fr", "it", "ja", "ko", "nb-no", "nl", "pl", "pt-br", "ru",
		"sk", "sv", "vi", "zh-cn", "zh-tw"}, "|")
}

func main() {
	flag.Parse()
	pb.GomeetFakerSetLocale(locale)
	addressBook := pb.NewAddressBookGomeetFaker()
	fmt.Printf("Address book (%s)\n", addressBook.GetUuid())
	for i, c := range addressBook.GetContacts() {
		fmt.Printf("  Contact (%d: %s) :\n", i, c.GetUuid())
		fmt.Printf("    Name      : %s\n", c.GetName())
		fmt.Printf("    Birthday  : %s\n", c.GetBirthday())
		fmt.Printf("    Avatars   : %s\n", strings.Join(c.GetAvatarsUrl(), "\n                "))
		fmt.Printf("    Notes     : %s\n", strings.Join(c.GetNotes(), "\n                "))
		phones := []string{}
		for i, phone := range c.GetPhones() {
			phones = append(
				phones,
				fmt.Sprintf(
					"%s / %s - (%d: %s)",
					phone.GetType(),
					phone.GetPhoneNumber(),
					i,
					phone.GetUuid(),
				),
			)
		}
		fmt.Printf("    Phones    : %s\n", strings.Join(phones, "\n                "))
		addresses := []string{}
		for i, addr := range c.GetAddresses() {
			addresses = append(
				addresses,
				fmt.Sprintf(
					"%s (%d: %s) - %s %s %s",
					addr.GetType(),
					i,
					addr.GetUuid(),
					addr.GetStreetAddress(),
					addr.GetPostcode(),
					addr.GetCity(),
				),
			)
		}
		fmt.Printf("    Addresses : %s\n", strings.Join(addresses, "\n                "))
		company := c.GetCompany()
		fmt.Printf("    Company (%s) :\n", company.GetUuid())
		fmt.Printf("      Name      : %s %s\n", company.GetName(), company.GetSuffix())
		fmt.Printf("                  %s\n", company.GetCatchPhrase())
		fmt.Printf("                  %s\n", company.GetBs())
		fmt.Printf("      Notes     : %s\n", company.GetNote())
		fmt.Printf("      Logo      : %s\n", company.GetLogo())
		fmt.Printf("      Codes     : ein: %s, duns_number: %s\n", company.GetEin(), company.GetDunsNumber())
		addresses = []string{}
		for i, addr := range company.GetAddresses() {
			addresses = append(
				addresses,
				fmt.Sprintf(
					"%s (%d: %s) - %s %s %s",
					addr.GetType(),
					i,
					addr.GetUuid(),
					addr.GetStreetAddress(),
					addr.GetPostcode(),
					addr.GetCity(),
				),
			)
		}
		fmt.Printf("      Addresses : %s\n", strings.Join(addresses, "\n                  "))
	}
}
