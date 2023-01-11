package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateOrderANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateOrderANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ ID: "123" }
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateOrderANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ ID: "123", Price: 10 }
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValisParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{ ID: "123", Price: 10, Tax: 2.0 }
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValisParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOder("123", 10, 2.3)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.3, order.Tax)
}

func TestGivenAPrinceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOder("123", 10, 2.3)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.3, order.FinalPrice)
}