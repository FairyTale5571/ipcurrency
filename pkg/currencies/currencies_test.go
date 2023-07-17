package currencies

import (
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"net/http"
	"testing"
)

func Test_currencies_GetExchangeUAHRates(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		want     []CurrencyRateToUAH
		want1    *errorops.Error
	}{
		{
			name:     "uah-usd",
			currency: "USD",
		},
		{
			name:     "uah-eur",
			currency: "EUR",
		},
		{
			name:     "uah-gbp",
			currency: "GBP",
		},
		{
			name:     "uah-cad",
			currency: "CAD",
		},
		{
			name:     "uah-pln",
			currency: "PLN",
		},
	}
	clientCurrency := NewCurrencies(http.DefaultClient, "74a864d459c64274d8095adc")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := clientCurrency.GetExchangeUAHRates(tt.currency)
			if got1 != nil {
				t.Errorf("GetExchangeUAHRates() got1 = %v, want %v", got1, tt.want1)
				return
			}
			t.Logf("GetExchangeUAHRates() got = %v, want %v", got, tt.want)
		})
	}
}
