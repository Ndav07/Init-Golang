package validation

import (
	"encoding/json"
	"errors"
	"ql/crud/src/configuration/opps"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl ut.Translator
)

func init() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *opps.Opps {
	var jsonError *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonError) {
		return opps.NewBadRequestError("Invalid field type")
	}	else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []opps.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := opps.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return opps.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return opps.NewBadRequestError("Error trying to convert fields")
	}
}