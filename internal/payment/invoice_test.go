package payment_test

import (
	"errors"
	"testing"

	"github.com/ossan-dev/go-debug-vsc/internal/payment"
	"github.com/stretchr/testify/assert"
)

func TestCalculateAmount(t *testing.T) {
	// Arrange
	testSuite := []struct {
		name           string
		itemsPurchased map[string]int
		itemsInStock   map[string]int
		itemsPrices    map[string]float64
		amount         float64
		wantErr        error
	}{
		{
			name:           "ItemNotFound",
			itemsPurchased: map[string]int{"mobile phone": 1},
			itemsInStock:   map[string]int{"TV": 20},
			amount:         0.00,
			wantErr:        errors.New("item not found"),
		},
		{
			name:           "QtyNotEnough",
			itemsPurchased: map[string]int{"mobile phone": 1},
			itemsInStock:   map[string]int{"mobile phone": 0},
			amount:         0.00,
			wantErr:        errors.New("qty in stock not enough"),
		},
		{
			name:           "ItemNotInCatalog",
			itemsPurchased: map[string]int{"mobile phone": 1},
			itemsInStock:   map[string]int{"mobile phone": 40, "TV": 20},
			itemsPrices:    map[string]float64{"TV": 500.00},
			amount:         0.00,
			wantErr:        errors.New("item no longer in catalog"),
		},
		{
			name:           "PriceMustBePositive",
			itemsPurchased: map[string]int{"mobile phone": 1},
			itemsInStock:   map[string]int{"mobile phone": 40, "TV": 20},
			itemsPrices:    map[string]float64{"mobile phone": -5.00, "TV": 500.00},
			amount:         0.00,
			wantErr:        errors.New("price cannot be zero or less"),
		},
		{
			name:           "HappyPath",
			itemsPurchased: map[string]int{"mobile phone": 2},
			itemsInStock:   map[string]int{"mobile phone": 40, "TV": 20},
			itemsPrices:    map[string]float64{"mobile phone": 350.00, "TV": 500.00},
			amount:         700.00,
			wantErr:        nil,
		},
	}
	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			amount, err := payment.CalculateAmount(tt.itemsPurchased, tt.itemsInStock, tt.itemsPrices)
			// Assert
			assert.Equal(t, tt.amount, amount)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
