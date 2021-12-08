package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// função de teste do cartão de crédito
func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Jose da Silva", 12, 2024, 123)
	assert.Equal(t, "Invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}

// função de teste do mês de expiração do cartão de crédito
func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 11, 2024, 123)
	assert.Nil(t, err)
}

// função de teste do ano de expiração do cartão de crédito
func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
}
