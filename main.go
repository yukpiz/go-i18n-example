package main

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func main() {
	bundle = &i18n.Bundle{DefaultLanguage: language.Japanese}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("ja-JP.toml")

	localizer := i18n.NewLocalizer(bundle, "ja-JP")
	fmt.Printf("%+v\n", localizer)

	// Hello go-i18n
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TestMessage"}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Example1.Hello1"}))

	err := errors.New("custom exception!")
	if err != nil {
		//i18n.T("Example1.Hello1")
	}
}
