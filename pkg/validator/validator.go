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
			return fmt.Errorf("%s は必須です。", err.Field())

		// 文字数
		case "min":
			return fmt.Errorf("%s の文字数が足りません。", err.Field())
		case "max":
			return fmt.Errorf("%s の文字数が多すぎます。", err.Field())
		case "gt":
			return fmt.Errorf("%s の文字数が足りません。", err.Field())
		case "lt":
			return fmt.Errorf("%s の文字数が多すぎます。", err.Field())
		case "len":
			return fmt.Errorf("%s の文字数を確認してください。", err.Field())

		default:
			return fmt.Errorf("%s に不正な値が入力されました。", err.Field())
		}
	}

	return nil
}
