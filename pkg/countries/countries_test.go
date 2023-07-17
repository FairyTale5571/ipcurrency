package countries

import (
	"net/http"
	"testing"
	"time"

	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"github.com/stretchr/testify/require"
)

func Test_currencies_GetCountryCurrency(t *testing.T) {
	tests := []struct {
		name             string
		countryCode      string
		expectedResponse map[string]CurrencyInfo
		expectedError    *errorops.Error
	}{
		{
			name:             "success UA",
			countryCode:      "UA",
			expectedResponse: map[string]CurrencyInfo{"UAH": {Name: "Ukrainian hryvnia", Symbol: "₴"}},
			expectedError:    nil,
		},
		{
			name:             "success US",
			countryCode:      "US",
			expectedResponse: map[string]CurrencyInfo{"USD": {Name: "United States dollar", Symbol: "$"}},
			expectedError:    nil,
		},
		{
			name:             "success AU",
			countryCode:      "AU",
			expectedResponse: map[string]CurrencyInfo{"AUD": {Name: "Australian dollar", Symbol: "$"}},
			expectedError:    nil,
		},
		{
			name:             "empty",
			countryCode:      "",
			expectedResponse: nil,
			expectedError: &errorops.Error{
				Code:        http.StatusBadRequest,
				Description: "Failed to get country currency",
				Message:     []string{"Required argument [String codes] not specified"},
				Value:       "Check country code",
			},
		},
		{
			name:             "not found",
			countryCode:      "mos",
			expectedResponse: nil,
			expectedError: &errorops.Error{
				Code:        http.StatusNotFound,
				Description: "Failed to get country currency",
				Message:     []string{"Not Found"},
				Value:       "Check country code",
			},
		},
		{
			name:             "invalid country code",
			countryCode:      "moskoviya",
			expectedResponse: nil,
			expectedError: &errorops.Error{
				Code:        http.StatusBadRequest,
				Description: "Failed to get country currency",
				Message:     []string{"Bad Request"},
				Value:       "Check country code",
			},
		},
	}
	clientCurrencies := &currencies{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := clientCurrencies.GetCountryCurrencies(tt.countryCode)
			require.Equal(t, err, tt.expectedError)
			require.Equal(t, tt.expectedResponse, resp)
		})
	}
}
