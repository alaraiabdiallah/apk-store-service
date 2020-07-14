package helpers

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func Validate(data interface{}) error {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	v := validator.New()
	en_translations.RegisterDefaultTranslations(v, trans)
	if err := v.Struct(data); err != nil {
		// translate all error at once
		errs := err.(validator.ValidationErrors)
		errMsg := ""
		for i, e := range errs {
			if i == 1 { break }
			errMsg = e.Translate(trans)
		}
		return errors.New(errMsg)
	}
	return nil
}
