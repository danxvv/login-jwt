package v1

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"unicode"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		en := en.New()
		uni = ut.New(en, en)
		trans, _ := uni.GetTranslator("en")

		validate.RegisterTranslation("password", trans, passwordregistrationFunc, passwordtranslateFunc)
		validate.RegisterTranslation("email", trans, emailregistrationFunc, emailtranslateFunc)
		validate.RegisterTranslation("number", trans, phoneRegistrationFunc, phoneTranslateFunc)
		err := validate.RegisterValidation("password", passWordValidation)

		if err != nil {
			return nil
		}
	}
	return validate
}

func passWordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 6 || len(password) > 12 {
		return false
	}

	var (
		hasUpper, hasLower, hasNumber, hasSpecial bool
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func passwordregistrationFunc(ut ut.Translator) error {
	return ut.Add("password", "{0} must be between 6-12 characters long, include at least one upper case letter, one lower case letter, one number, and one special character", true)
}

func passwordtranslateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("password", fe.Field())
	return t
}

func emailregistrationFunc(ut ut.Translator) error {
	return ut.Add("email", "{0} must be a valid email address", true)
}

func emailtranslateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("email", fe.Field())
	return t
}

func phoneRegistrationFunc(ut ut.Translator) error {
	return ut.Add("number", "{0} must be a valid phone number and length should be 10", true)
}

func phoneTranslateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("number", fe.Field())
	return t
}
