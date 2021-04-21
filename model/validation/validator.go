package validation

import (
	"reflect"
	"strings"

	"my/v1/errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validation container validator library and transaltor
type Validation struct {
	V *validator.Validate
	T *ut.Translator
}

// NewValidation create new Validation struct instance
func NewValidation() *Validation {
	validation := &Validation{
		V: validator.New(),
	}
	trans := initializeTranslation(validation.V)
	validation.T = trans
	registerFunc(validation.V)
	return validation
}

// Initialize initializes and returns the UniversalTranslator instance for the application
func initializeTranslation(validate *validator.Validate) *ut.Translator {

	// initialize translator
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	// initialize translations
	en_translations.RegisterDefaultTranslations(validate, trans)
	return &trans
}

func registerFunc(validate *validator.Validate) {
	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Validate performs validation on a form
func (v *Validation) Validate(form interface{}) []error {
	var errResp []error
	if err := v.V.Struct(form); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			err := errors.InvalidField.New(e.Translate(*v.T))
			key := strings.SplitAfterN(e.Namespace(), ".", 2)
			err = errors.AddErrorContext(err, key[1], e.Translate(*v.T))
			errResp = append(errResp, err)
		}
	}
	return errResp
}
