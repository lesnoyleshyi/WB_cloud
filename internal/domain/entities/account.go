package entities

import (
	"github.com/shopspring/decimal"
)

type Account struct {
	Id      string          `json:"id"`
	Balance decimal.Decimal `json:"balance"`
}
