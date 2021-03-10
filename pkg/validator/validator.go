package validator

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// CustomValidator -
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator - init
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			return fmt.Errorf("%s is required.", err.Field())
		default:
			return fmt.Errorf("%s is invalid.", err.Field())
		}
	}

	return nil
}
