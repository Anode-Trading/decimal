package decimal

import (
	"github.com/shopspring/decimal"
)

type Decimal string

func (d Decimal) String() string {
	return string(d)
}

func New(v string) Decimal {
	return Decimal(v)
}

func (d Decimal) Add(v Decimal) Decimal {
	return New(decimal.RequireFromString(d.String()).
		Add(decimal.RequireFromString(v.String())).
		String())
}

func (d Decimal) Sub(v Decimal) Decimal {
	return New(decimal.RequireFromString(d.String()).
		Sub(decimal.RequireFromString(v.String())).
		String())
}

func (d Decimal) Mul(v Decimal) Decimal {
	return New(decimal.RequireFromString(d.String()).
		Mul(decimal.RequireFromString(v.String())).
		String())
}

func (d Decimal) Div(v Decimal) Decimal {
	return New(decimal.RequireFromString(d.String()).
		Div(decimal.RequireFromString(v.String())).
		String())
}

func (d Decimal) DivRound(v Decimal, precision int32) Decimal {
	return New(decimal.RequireFromString(d.String()).
		DivRound(decimal.RequireFromString(v.String()), precision).
		String())
}
