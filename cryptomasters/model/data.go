package model

import "fmt"

type Price float64

type Rate struct {
	Currency string
	Price    Price
}

func (r *Rate) String() string {
	return fmt.Sprintf("1 %s = %f USD", r.Currency, r.Price)
}
