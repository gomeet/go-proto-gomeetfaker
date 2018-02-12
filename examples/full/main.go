package main

import (
	"flag"
	fmt "fmt"
	"strings"

	"github.com/gomeet/go-proto-gomeetfaker/examples/full/pb"
)

var (
	msgT       string
	locale     string
	msgPrintFn = map[string]func(){
		"without_faker":              printWithoutFakerMessageGomeetFaker,
		"value":                      printValueMessageGomeetFaker,
		"uuid":                       printUuidMessageGomeetFaker,
		"address":                    printAddressMessageGomeetFaker,
		"app":                        printAppMessageGomeetFaker,
		"avatar":                     printAvatarMessageGomeetFaker,
		"bitcoin":                    printBitcoinMessageGomeetFaker,
		"business":                   printBusinessMessageGomeetFaker,
		"code":                       printCodeMessageGomeetFaker,
		"commerce":                   printCommerceMessageGomeetFaker,
		"company":                    printCompanyMessageGomeetFaker,
		"date":                       printDateMessageGomeetFaker,
		"finance":                    printFinanceMessageGomeetFaker,
		"hacker":                     printHackerMessageGomeetFaker,
		"internet":                   printInternetMessageGomeetFaker,
		"lorem":                      printLoremMessageGomeetFaker,
		"name":                       printNameMessageGomeetFaker,
		"number":                     printNumberMessageGomeetFaker,
		"phone_number":               printPhoneNumberMessageGomeetFaker,
		"team":                       printTeamMessageGomeetFaker,
		"time":                       printTimeMessageGomeetFaker,
		"embed_all":                  printEmbedAllFakerMessageGomeetFaker,
		"embed_oneof":                printEmbedOneOfFakerMessageGomeetFaker,
		"embed_all_repeated":         printEmbedAllFakerRepeatedMessageGomeetFaker,
		"embed_all_repeated_norules": printEmbedAllFakerRepeatedNoRulesMessage,
		"complex":                    printComplexMessageGomeetFaker,
		"scalar_repeated":            printScalarRepeatedRulesOnlyMessageGomeetFaker,
		"enum":                       printEnumMessageGomeetFaker,
	}
)

func init() {
	flag.StringVar(&msgT, "message", "all", fmt.Sprintf("message type [%s]", allMessageParams()))
	flag.StringVar(&locale, "locale", "fr", fmt.Sprintf("locales [%s]", allLocalesParams()))
}

func allMessageParams() string {
	a := []string{"all"}
	for msg, _ := range msgPrintFn {
		a = append(a, msg)
	}
	return strings.Join(a, "|")
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
	printMessage(msgT)
}

func printMessage(msg string) {
	switch msg {
	case "all":
		fmt.Println("Print all gomeetfaker messages :")
		fmt.Println("")
		for msgT, _ := range msgPrintFn {
			printMessage(msgT)
		}
	default:
		if fn, ok := msgPrintFn[msg]; ok {
			fn()
		} else {
			fmt.Println("Bad arg")
		}
	}
	fmt.Println("")
}

func printComplexMessageGomeetFaker() {
	aComplexMessage := pb.NewComplexMessageGomeetFaker()
	fmt.Println("ComplexMessage gomeetfaker : ")
	fmt.Println(aComplexMessage)
	fmt.Println("")
}
func printWithoutFakerMessageGomeetFaker() {
	aWithoutFakerMessage := pb.NewWithoutFakerMessageGomeetFaker()
	fmt.Println("WithoutFakerMessage gomeetfaker : ")
	fmt.Println(aWithoutFakerMessage)
	fmt.Println("")
}
func printValueMessageGomeetFaker() {
	aValueMessage := pb.NewValueMessageGomeetFaker()
	fmt.Println("ValueMessage gomeetfaker : ")
	fmt.Println(aValueMessage)
	fmt.Println("")
}
func printAddressMessageGomeetFaker() {
	aAddressMessage := pb.NewAddressMessageGomeetFaker()
	fmt.Println("AddressMessage gomeetfaker : ")
	fmt.Println(aAddressMessage)
	fmt.Println("")
}
func printAppMessageGomeetFaker() {
	aAppMessage := pb.NewAppMessageGomeetFaker()
	fmt.Println("AppMessage gomeetfaker : ")
	fmt.Println(aAppMessage)
	fmt.Println("")
}
func printAvatarMessageGomeetFaker() {
	aAvatarMessage := pb.NewAvatarMessageGomeetFaker()
	fmt.Println("AvatarMessage gomeetfaker : ")
	fmt.Println(aAvatarMessage)
	fmt.Println("")
}
func printBitcoinMessageGomeetFaker() {
	aBitcoinMessage := pb.NewBitcoinMessageGomeetFaker()
	fmt.Println("BitcoinMessage gomeetfaker : ")
	fmt.Println(aBitcoinMessage)
	fmt.Println("")
}
func printBusinessMessageGomeetFaker() {
	aBusinessMessage := pb.NewBusinessMessageGomeetFaker()
	fmt.Println("BusinessMessage gomeetfaker : ")
	fmt.Println(aBusinessMessage)
	fmt.Println("")
}
func printCodeMessageGomeetFaker() {
	aCodeMessage := pb.NewCodeMessageGomeetFaker()
	fmt.Println("CodeMessage gomeetfaker : ")
	fmt.Println(aCodeMessage)
	fmt.Println("")
}
func printCommerceMessageGomeetFaker() {
	aCommerceMessage := pb.NewCommerceMessageGomeetFaker()
	fmt.Println("CommerceMessage gomeetfaker : ")
	fmt.Println(aCommerceMessage)
	fmt.Println("")
}
func printCompanyMessageGomeetFaker() {
	aCompanyMessage := pb.NewCompanyMessageGomeetFaker()
	fmt.Println("CompanyMessage gomeetfaker : ")
	fmt.Println(aCompanyMessage)
	fmt.Println("")
}
func printDateMessageGomeetFaker() {
	aDateMessage := pb.NewDateMessageGomeetFaker()
	fmt.Println("DateMessage gomeetfaker : ")
	fmt.Println(aDateMessage)
	fmt.Println("")
}
func printFinanceMessageGomeetFaker() {
	aFinanceMessage := pb.NewFinanceMessageGomeetFaker()
	fmt.Println("FinanceMessage gomeetfaker : ")
	fmt.Println(aFinanceMessage)
	fmt.Println("")
}
func printHackerMessageGomeetFaker() {
	aHackerMessage := pb.NewHackerMessageGomeetFaker()
	fmt.Println("HackerMessage gomeetfaker : ")
	fmt.Println(aHackerMessage)
	fmt.Println("")
}
func printInternetMessageGomeetFaker() {
	aInternetMessage := pb.NewInternetMessageGomeetFaker()
	fmt.Println("InternetMessage gomeetfaker : ")
	fmt.Println(aInternetMessage)
	fmt.Println("")
}
func printLoremMessageGomeetFaker() {
	aLoremMessage := pb.NewLoremMessageGomeetFaker()
	fmt.Println("LoremMessage gomeetfaker : ")
	fmt.Println(aLoremMessage)
	fmt.Println("")
}
func printNameMessageGomeetFaker() {
	aNameMessage := pb.NewNameMessageGomeetFaker()
	fmt.Println("NameMessage gomeetfaker : ")
	fmt.Println(aNameMessage)
	fmt.Println(aNameMessage.GetName())
	fmt.Println("")
}
func printNumberMessageGomeetFaker() {
	aNumberMessage := pb.NewNumberMessageGomeetFaker()
	fmt.Println("NumberMessage gomeetfaker : ")
	fmt.Println(aNumberMessage)
	fmt.Println("")
}
func printPhoneNumberMessageGomeetFaker() {
	aPhoneNumberMessage := pb.NewPhoneNumberMessageGomeetFaker()
	fmt.Println("PhoneNumberMessage gomeetfaker : ")
	fmt.Println(aPhoneNumberMessage)
	fmt.Println("")
}
func printTeamMessageGomeetFaker() {
	aTeamMessage := pb.NewTeamMessageGomeetFaker()
	fmt.Println("TeamMessage gomeetfaker : ")
	fmt.Println(aTeamMessage)
	fmt.Println("")
}
func printTimeMessageGomeetFaker() {
	aTimeMessage := pb.NewTimeMessageGomeetFaker()
	fmt.Println("TimeMessage gomeetfaker : ")
	fmt.Println(aTimeMessage)
	fmt.Println("")
}
func printEmbedAllFakerMessageGomeetFaker() {
	aEmbedAllFakerMessage := pb.NewEmbedAllFakerMessageGomeetFaker()
	fmt.Println("EmbedAllFakerMessage gomeetfaker : ")
	fmt.Println(aEmbedAllFakerMessage)
	fmt.Println("")
}
func printEmbedOneOfFakerMessageGomeetFaker() {
	aEmbedOneOfFakerMessage := pb.NewEmbedOneOfFakerMessageGomeetFaker()
	fmt.Println("EmbedOneOfFakerMessage gomeetfaker : ")
	fmt.Println(aEmbedOneOfFakerMessage)
	fmt.Println("")
}
func printEmbedAllFakerRepeatedMessageGomeetFaker() {
	aEmbedAllFakerRepeatedMessage := pb.NewEmbedAllFakerRepeatedMessageGomeetFaker()
	fmt.Println("EmbedAllFakerRepeatedMessage gomeetfaker : ")
	fmt.Println(aEmbedAllFakerRepeatedMessage)
	fmt.Println("")
}
func printEmbedAllFakerRepeatedNoRulesMessage() {
	aEmbedAllFakerRepeatedNoRulesMessage := pb.NewEmbedAllFakerRepeatedNoRulesMessageGomeetFaker()
	fmt.Println("EmbedAllFakerRepeatedNoRulesMessage gomeetfaker : ")
	fmt.Println(aEmbedAllFakerRepeatedNoRulesMessage)
	fmt.Println("")
}
func printScalarRepeatedRulesOnlyMessageGomeetFaker() {
	aScalarRepeatedRulesOnlyMessage := pb.NewScalarRepeatedRulesOnlyMessageGomeetFaker()
	fmt.Println("ScalarRepeatedRulesOnlyMessage gomeetfaker : ")
	fmt.Println(aScalarRepeatedRulesOnlyMessage)
	fmt.Println("")
}
func printUuidMessageGomeetFaker() {
	aUuidMessage := pb.NewUuidMessageGomeetFaker()
	fmt.Println("UuidMessage gomeetfaker : ")
	fmt.Println(aUuidMessage)
	fmt.Println("")
}
func printEnumMessageGomeetFaker() {
	aEnumMessage := pb.NewEnumMessageGomeetFaker()
	fmt.Println("EnumMessage gomeetfaker : ")
	fmt.Println(aEnumMessage)
	fmt.Println("")
}
