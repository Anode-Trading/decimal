package decimal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := "123.456"
	d := New(s)

	assert.Equal(t, s, d.String())
}

func TestNewFromFloat(t *testing.T) {
	f := 123.456
	e := "123.456"

	d := NewFromFloat(f)

	assert.Equal(t, e, d.String())
}

func TestDecimal_Add(t *testing.T) {
	a := New("123.456")
	b := New("789.123")
	e := New("912.579")

	assert.True(t, e.Equal(a.Add(b)))
}

func TestDecimal_Sub(t *testing.T) {
	a := New("262.19700000")
	b := New("0.00106400")
	e := New("262.19593600")

	assert.True(t, e.Equal(a.Sub(b)))
}

func TestDecimal_Mul(t *testing.T) {
	a := New("262.19700000")
	b := New("0.00106400")
	e := New("0.278977608")

	assert.True(t, e.Equal(a.Mul(b)))
}

func TestDecimal_Div(t *testing.T) {
	a := New("154.94400000")
	b := New("37.26837740")
	e := New("4.1575193450734992")

	assert.True(t, e.Equal(a.Div(b)))
}

func TestDecimal_DivRound(t *testing.T) {
	a := New("154.94400000")
	b := New("37.26837740")
	e := New("4.1575193450735")

	assert.True(t, e.Equal(a.DivRound(b, 13)))
}

func TestDecimal_Cmp(t *testing.T) {
	a := New("154.94400000")
	b := New("37.26837740")
	e := 1

	assert.Equal(t, e, a.Cmp(b))
}

func TestDecimal_Equal(t *testing.T) {
	a := New("154.94400000")
	b := New("154.944")

	assert.True(t, a.Equal(b))
}

func TestDecimal_GreaterThan(t *testing.T) {
	a := New("154.94400000")
	b := New("37.26837740")

	assert.True(t, a.GreaterThan(b))
}

func TestDecimal_LessThan(t *testing.T) {
	a := New("154.94400000")
	b := New("37.26837740")

	assert.True(t, b.LessThan(a))
}

func TestDecimal_GreaterThanOrEqual(t *testing.T) {
	a := New("154.94400000")
	b := New("37.26837740")

	assert.True(t, a.GreaterThanOrEqual(b))

	b = New("154.944")
	assert.True(t, a.GreaterThanOrEqual(b))
}

func TestDecimal_LessThanOrEqual(t *testing.T) {
	a := New("37.26837740")
	b := New("154.94400000")

	assert.True(t, a.LessThanOrEqual(b))

	a = New("154.944")
	assert.True(t, b.LessThanOrEqual(a))
}

func TestDecimal_IsZero(t *testing.T) {
	a := New("0")
	b := New("154.944")

	assert.True(t, a.IsZero())
	assert.False(t, b.IsZero())
}

func TestMarshal(t *testing.T) {
	exp := `{"decimal":"123.456"}`
	d := s{
		Decimal: "123.456",
	}

	b, err := json.Marshal(d)

	assert.NoError(t, err)
	assert.Equal(t, exp, string(b))
}

func TestUnmarshal(t *testing.T) {
	payload := `{"decimal":"123.456"}`
	exp := s{
		Decimal: "123.456",
	}

	var d s
	err := json.Unmarshal([]byte(payload), &d)

	assert.NoError(t, err)
	assert.Equal(t, exp, d)
}

func TestSumMap(t *testing.T) {
	a := map[string]Decimal{
		"1": New("1.1"),
		"2": New("2.02"),
		"3": New("3.003"),
	}
	e := New("6.123")

	assert.Equal(t, e, SumMap(a))
}

type s struct {
	Decimal Decimal `json:"decimal"`
}
