package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/santhoshvempali/simplebank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		ans := util.IsSupportedCurrency(currency)
		if ans {
			return true
		}
	}
	return false
}
