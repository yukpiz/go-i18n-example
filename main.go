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
	bundle.MustLoadMessageFile("ja.toml")
	bundle.MustLoadMessageFile("en-US.toml")

	localizer := i18n.NewLocalizer(bundle, "ja;q=0.9, en;q=0.8")
	localizer = i18n.NewLocalizer(bundle, "ja;q=0.8, en;q=0.9")
	localizer = i18n.NewLocalizer(bundle, "ja;q=0.9, en;q=0.9, fr;q=0.9")

	// Hello go-i18n
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TestMessage"}))
	fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Example1.Hello1"}))

	err := errors.New("custom exception!")
	if err != nil {
		//i18n.T("Example1.Hello1")
	}
}
