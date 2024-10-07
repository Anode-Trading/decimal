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

func NewFromFloat(v float64) Decimal {
	return New(decimal.NewFromFloat(v).String())
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

func (d Decimal) Cmp(v Decimal) int {
	return decimal.RequireFromString(d.String()).
		Cmp(decimal.RequireFromString(v.String()))
}

func (d Decimal) Equal(v Decimal) bool {
	return d.Cmp(v) == 0
}

func (d Decimal) GreaterThan(v Decimal) bool {
	return d.Cmp(v) == 1
}

func (d Decimal) LessThan(v Decimal) bool {
	return d.Cmp(v) == -1
}

func (d Decimal) GreaterThanOrEqual(v Decimal) bool {
	return d.Cmp(v) != -1
}

func (d Decimal) LessThanOrEqual(v Decimal) bool {
	return d.Cmp(v) != 1
}

func (d Decimal) IsZero() bool {
	return d.Cmp(New("0")) == 0
}

func SumMap[T comparable](m map[T]Decimal) Decimal {
	sum := New("0")
	for _, v := range m {
		sum = sum.Add(v)
	}
	return sum
}
