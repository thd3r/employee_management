package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponseHandler struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value,omitempty"`
	Message string `json:"msg"`
}

func ValidateRequestStruct(data any) []*ErrorResponseHandler {
	var errors []*ErrorResponseHandler

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponseHandler

			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Message = ""

			if len(err.Value().(string)) != len(err.Param()) && err.Tag() == "min" {
				element.Message += fmt.Sprintf("The minimum value of the %s parameter is %v characters", err.Field(), err.Param())
			} else if len(err.Value().(string)) != len(err.Param()) && err.Tag() == "max" {
				element.Message += fmt.Sprintf("The maximum value of the %s parameter is %v characters", err.Field(), err.Param())
			} else {
				element.Message += fmt.Sprintf("The value of the %s parameter cannot be empty", err.Field())
			}

			errors = append(errors, &element)
		}
	}

	return errors
}
