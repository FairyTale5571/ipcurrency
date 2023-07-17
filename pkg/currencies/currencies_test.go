package currencies

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"

	"github.com/fairytale5571/ipcurrency/pkg/errorops"
)

func Test_currencies_GetExchangeUAHRates(t *testing.T) {
	tests := []struct {
		name          string
		currency      string
		expected      []CurrencyRateToUAH
		expectedError *errorops.Error
	}{
		{
			name:     "uah-usd",
			currency: "USD",
			expected: []CurrencyRateToUAH{{Currency: "USD"}},
		},
		{
			name:     "uah-eur",
			currency: "EUR",
			expected: []CurrencyRateToUAH{{Currency: "EUR"}},
		},
		{
			name:     "uah-gbp",
			currency: "GBP",
			expected: []CurrencyRateToUAH{{Currency: "GBP"}},
		},
		{
			name:     "uah-cad",
			currency: "CAD",
			expected: []CurrencyRateToUAH{{Currency: "CAD"}},
		},
		{
			name:     "uah-pln",
			currency: "PLN",
			expected: []CurrencyRateToUAH{{Currency: "PLN"}},
		},
		{
			name:     "uah-jpy",
			currency: "JPY",
			expected: []CurrencyRateToUAH{{Currency: "JPY"}},
		},
		{
			name: "uah-empty",
			expectedError: &errorops.Error{
				Code:        http.StatusNotFound,
				Description: "not found",
				Value:       "latest/",
			},
		},
		{
			name:     "uah-unknown",
			currency: "unknown",
			expectedError: &errorops.Error{
				Code:        http.StatusNotFound,
				Description: "not found",
				Value:       "latest/unknown",
			},
		},
	}
	clientCurrency := NewCurrencies(http.DefaultClient, "74a864d459c64274d8095adc")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rates, err := clientCurrency.GetExchangeUAHRates(tt.currency)
			require.Equal(t, err, tt.expectedError)
			if tt.expectedError != nil {
				return
			}
			require.Equal(t, rates[0].Currency, tt.expected[0].Currency)
		})
	}
}
