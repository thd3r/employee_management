package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorHandler struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value,omitempty"`
	Message string `json:"msg"`
}

func ValidateRequest(data any) []*ErrorHandler {
	var errors []*ErrorHandler

	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorHandler

			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Message = ""

			if len(err.Value().(string)) != len(err.Param()) {
				element.Message += fmt.Sprintf("The maximum value of the %s parameter is %v", err.Field(), err.Param())
			} else {
				element.Message += fmt.Sprintf("The value of the %s parameter cannot be empty", err.Field())
			}

			errors = append(errors, &element)
		}
	}

	return errors
}
