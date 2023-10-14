package models

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)

	if err != nil {
		validatorErrs := err.(validator.ValidationErrors)

		for _, e := range validatorErrs {
			translatedErr := fmt.Errorf(e.Translate(cv.Trans))

			return translatedErr
		}
	}

	return nil
}

func NewCustomValidator(e *echo.Echo) echo.Validator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &CustomValidator{
		Validator: validate,
		Trans:     trans,
	}
}
