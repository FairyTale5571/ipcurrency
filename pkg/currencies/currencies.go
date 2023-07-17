package currencies

import (
	"encoding/json"
	"fmt"
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"io"
	"net/http"
)

type CurrenciesResponse struct {
	Result          string             `json:"result"`
	Documentation   string             `json:"documentation"`
	TermsOfUse      string             `json:"terms_of_use"`
	BaseCode        string             `json:"base_code"`
	ConversionRates map[string]float64 `json:"conversion_rates"`
}

type ErrorResponse struct {
	Result        string `json:"result"`
	Documentation string `json:"documentation"`
	TermsOfUse    string `json:"terms-of-use"`
	ErrorType     string `json:"error-type"`
}

type CurrencyRateToUAH struct {
	Currency  string  `json:"currency"`
	RateToUAH float64 `json:"rateToUAH"`
}

type Provider interface {
	GetExchangeUAHRates(countryCode string) ([]CurrencyRateToUAH, *errorops.Error)
}

type currencies struct {
	client *http.Client
	apiKey string
}

func NewCurrencies(client *http.Client, apiKey string) Provider {
	return &currencies{
		client: client,
		apiKey: apiKey,
	}
}

const endpoint = "https://v6.exchangerate-api.com/v6"

func (c *currencies) request(path string, response any) *errorops.Error {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", endpoint, c.apiKey, path), http.NoBody)
	if err != nil {
		return errorops.NewError(
			http.StatusInternalServerError,
			"failed to make HTTP request",
			err.Error(),
		)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return errorops.NewError(
			http.StatusInternalServerError,
			"failed to make HTTP request",
			err.Error(),
		)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	case http.StatusNotFound:
		return errorops.NewError(
			http.StatusNotFound,
			"not found",
			path,
		)
	case http.StatusUnauthorized:
		return errorops.NewError(
			http.StatusUnauthorized,
			"unauthorized",
			"invalid API key",
		)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errorops.NewError(
			http.StatusInternalServerError,
			"failed to read response body",
			err.Error(),
		)
	}

	if response != nil {
		if err := json.Unmarshal(responseBody, &response); err != nil {
			return errorops.NewError(
				http.StatusInternalServerError,
				"failed to decode response body",
				err.Error(),
			)
		}
	}
	return nil
}

func (c *currencies) GetExchangeUAHRates(countryCode string) ([]CurrencyRateToUAH, *errorops.Error) {
	var successResp CurrenciesResponse
	if err := c.request(fmt.Sprintf("latest/%s", countryCode), &successResp); err != nil {
		return nil, err
	}

	rates := successResp.ConversionRates
	var ratesToUAH []CurrencyRateToUAH
	if val, ok := rates["UAH"]; ok {
		ratesToUAH = append(ratesToUAH, CurrencyRateToUAH{
			Currency:  successResp.BaseCode,
			RateToUAH: val,
		})
	}

	return ratesToUAH, nil
}
