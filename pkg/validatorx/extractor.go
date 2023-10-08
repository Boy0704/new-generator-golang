package validatorx

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ExtractError(err error) (result []string) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return
	}

	validationErrors := err.(validator.ValidationErrors)
	for _, e := range validationErrors {
		result = append(result, fmt.Sprintf("Field: %v Message: %v", e.Field(), e.Error()))
	}
	return
}
