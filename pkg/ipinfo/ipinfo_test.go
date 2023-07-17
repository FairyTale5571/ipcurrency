package ipinfo

import (
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func Test_ipInfo_RetrieveAddress(t *testing.T) {
	cases := []struct {
		name             string
		address          string
		expectedResponse IPInformation
		expectedError    *errorops.Error
	}{
		{
			name:    "success",
			address: "24.48.0.1",
			expectedResponse: IPInformation{
				Status:      "success",
				Country:     "Canada",
				CountryCode: "CA",
				Region:      "QC",
				RegionName:  "Quebec",
				City:        "Montreal",
				Zip:         "H2Y",
				Lat:         45.504,
				Lon:         -73.552,
				Timezone:    "America/Toronto",
				Isp:         "Le Groupe Videotron Ltee",
				Org:         "Videotron Ltee",
				As:          "AS5769 Videotron Telecom Ltee",
				Query:       "24.48.0.1",
			},
			expectedError: nil,
		},
		{
			name:             "invalid address",
			address:          "1231231231231",
			expectedResponse: IPInformation{},
			expectedError: &errorops.Error{
				Code:        http.StatusBadRequest,
				Description: "Failed to retrieve IP address: invalid query",
				Message:     []string{"Please check the IP address and try again..."},
				Value:       "1231231231231",
			},
		},
		{
			name:             "reserved range",
			address:          "127.0.0.1",
			expectedResponse: IPInformation{},
			expectedError: &errorops.Error{
				Code:        http.StatusBadRequest,
				Description: "Failed to retrieve IP address: reserved range",
				Message:     []string{"Please check the IP address and try again..."},
				Value:       "127.0.0.1",
			},
		},
	}
	ipClient := &ipInfo{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ipClient.RetrieveAddress(tt.address)
			require.Equal(t, err, tt.expectedError)
			require.Equal(t, tt.expectedResponse, resp)

		})
	}
}
