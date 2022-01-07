package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/tkircsi/simple-bank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		for _, c := range util.AppConfig.Currency {
			if c == currency {
				return true
			}
		}
	}
	return false
}
