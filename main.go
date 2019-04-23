package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	//bundle = &i18n.Bundle{DefaultLanguage: language.Japanese}
	//bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//bundle.MustLoadMessageFile("config/locale/ja.toml")
	//bundle.MustLoadMessageFile("config/locale/en-US.toml")

	//localizer := i18n.NewLocalizer(bundle, "ja;q=0.9, en;q=0.8")
	//localizer = i18n.NewLocalizer(bundle, "ja;q=0.8, en;q=0.9")
	//localizer = i18n.NewLocalizer(bundle, "ja;q=0.8, en;q=0.9, fr;q=0.8")
	//localizer := i18n.NewLocalizer(bundle)

	//pp.Print(bundle)
	// Hello go-i18n
	//fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "TestMessage"}))
	//fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Example1.Hello1"}))

	//err := errors.New("custom exception!")
	//if err != nil {
	//	//i18n.T("Example1.Hello1")
	//}

	bundle := &i18n.Bundle{DefaultLanguage: language.Japanese}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	//for _, tag := range SupportedLanguages {
	//	fname := fmt.Sprintf("%s.%s", tag, ftype)
	//	fpath := filepath.Join(dir, fname)
	//	log.Println(fpath)
	//}
	bundle.MustLoadMessageFile("config/locale/ja.toml")

	log.Printf("%+v\n", bundle)
	localizer := i18n.NewLocalizer(bundle, "ja")
	log.Printf("%+v\n", localizer)
	//log.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Hello"}))
	log.Println(localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "Example1.Hello1"}))
	log.Println(localizer.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{
		ID:    "Example1.Hello1",
		Other: "Hello Other",
	}}))
}
