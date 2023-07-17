package countries

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fairytale5571/ipcurrency/pkg/errorops"
)

type CurrencyInfo struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type SuccessResp struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Currencies map[string]CurrencyInfo `json:"currencies"`
}

type errResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const endpoint = "https://restcountries.com/v3.1"

type currencies struct {
	client *http.Client
}

type Provider interface {
	GetCountryCurrencies(countryCode string) (map[string]CurrencyInfo, *errorops.Error)
}

func NewCurrencies(client *http.Client) Provider {
	return &currencies{
		client: client,
	}
}

func (c *currencies) GetCountryCurrencies(countryCode string) (map[string]CurrencyInfo, *errorops.Error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/alpha/%s", endpoint, countryCode), http.NoBody)
	if err != nil {
		return nil, errorops.NewError(
			http.StatusInternalServerError,
			"failed to make HTTP request",
			err.Error(),
			"try later...",
		)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errorops.NewError(
			http.StatusInternalServerError,
			"Failed to make HTTP request",
			err.Error(),
			"Try again later...",
		)
	}
	return c.handleResponse(resp)
}

func (c *currencies) handleResponse(resp *http.Response) (map[string]CurrencyInfo, *errorops.Error) {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errorops.NewError(
			http.StatusInternalServerError,
			"Failed to read response body",
			err.Error(),
			"Try again later...",
		)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp errResp
		if err := json.Unmarshal(bodyBytes, &errResp); err != nil {
			return nil, errorops.NewError(
				http.StatusInternalServerError,
				"Failed to unmarshal response body",
				err.Error(),
				"Try again later...",
			)
		}
		return nil, errorops.NewError(
			resp.StatusCode,
			"Failed to get country currency",
			"Check country code",
			errResp.Message,
		)
	}

	var successResp []SuccessResp
	if err := json.Unmarshal(bodyBytes, &successResp); err != nil {
		return nil, errorops.NewError(
			http.StatusInternalServerError,
			"Failed to unmarshal response body",
			err.Error(),
			"Try again later...",
		)
	}
	if len(successResp) == 0 {
		return nil, errorops.NewError(
			http.StatusInternalServerError,
			"Failed to get country currency",
			"Check country code",
			"Try again later...",
		)
	}

	var response map[string]CurrencyInfo
	for _, v := range successResp {
		response = v.Currencies
	}

	return response, nil
}
