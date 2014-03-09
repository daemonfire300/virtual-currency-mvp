package main

import (
	"fmt"
)

type WrongCurrencyError struct {
	CurrencyA *Currency
	CurrencyB *Currency
}

func (e *WrongCurrencyError) Error() string {
	return fmt.Sprintf("Currency of type %v missmatches Currency of type %v", e.CurrencyA, e.CurrencyB)
}
