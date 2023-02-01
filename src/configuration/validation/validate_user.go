package validation

import (
	"encoding/json"
	"errors"

	"github.com/andrade-felipe-dev/jwt-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_br_translation "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt_BR := en.New()
		un := ut.New(pt_BR, pt_BR)
		transl, _ = un.GetTranslator("pt_BR")
		pt_br_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(err, &jsonErr) {
		return rest_err.NewBadRequestError("Campo inválido")
	} else if errors.As(err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Alguns campos são inválidos", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Erro na conversão dos campos")
	}
}
