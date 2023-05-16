package util

import "fmt"

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	fmt.Println(currency)
	switch currency {
	case USD:
		return true
	case EUR:
		return true
	case CAD:
		return true
	}
	return false
}
