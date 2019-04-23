package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	bundle := &i18n.Bundle{DefaultLanguage: language.Japanese}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("ja-JP.toml")

	localizer := i18n.NewLocalizer(bundle, "ja-JP")
	fmt.Printf("%+v\n", localizer)

	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TestMessage"}))
}
