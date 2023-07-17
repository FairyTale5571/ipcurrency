package ipinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fairytale5571/ipcurrency/pkg/errorops"
)

const endpoint = "http://ip-api.com/json"

type ErrorResponse struct {
	Query   string `json:"query"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type IPInformation struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

type ipInfo struct {
	client *http.Client
}

type Provider interface {
	RetrieveAddress(address string) (IPInformation, *errorops.Error)
}

func NewIpInfo(client *http.Client) Provider {
	return &ipInfo{
		client: client,
	}
}

func (i *ipInfo) RetrieveAddress(address string) (IPInformation, *errorops.Error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", endpoint, address), http.NoBody)
	if err != nil {
		return IPInformation{}, errorops.NewError(
			http.StatusInternalServerError,
			"failed to make HTTP request",
			err.Error(),
			"try later...",
		)
	}
	resp, err := i.client.Do(req)
	if err != nil {
		return IPInformation{}, errorops.NewError(
			http.StatusInternalServerError,
			"Failed to make HTTP request",
			err.Error(),
			"Try again later...",
		)
	}
	return i.handleResponse(resp)
}

func (i *ipInfo) handleResponse(resp *http.Response) (IPInformation, *errorops.Error) {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return IPInformation{}, errorops.NewError(
			http.StatusInternalServerError,
			"Failed to read HTTP response",
			err.Error(),
			"Try again later...",
		)
	}

	var errResp ErrorResponse
	if err := json.Unmarshal(bodyBytes, &errResp); err == nil && errResp.Status == "fail" {
		return IPInformation{}, errorops.NewError(
			http.StatusBadRequest,
			fmt.Sprintf("Failed to retrieve IP address: %s", errResp.Message),
			errResp.Query,
			"Please check the IP address and try again...",
		)
	}

	if errResp.Status != "success" {
		return IPInformation{}, errorops.NewError(
			http.StatusBadRequest,
			fmt.Sprintf("Failed to retrieve IP address: %s", errResp.Message),
			errResp.Query,
			errResp.Message,
		)
	}

	var successResp IPInformation
	if err := json.Unmarshal(bodyBytes, &successResp); err != nil {
		return IPInformation{}, errorops.NewError(
			http.StatusInternalServerError,
			"Failed to parse HTTP response",
			err.Error(),
			"Try again later...",
		)
	}

	return successResp, nil
}
