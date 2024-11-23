package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/go_dev/simplebank/utils"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		if utils.IsSupportedCurrency(currency) {
			return true
		}
	}
	return false
}
