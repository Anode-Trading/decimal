package decimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := "123.456"
	d := New(s)

	assert.Equal(t, s, d.String())
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
