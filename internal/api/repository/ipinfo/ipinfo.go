package ipinfo

import (
	"net/http"
	"time"

	"github.com/fairytale5571/ipcurrency/internal/api/repository"
	"github.com/fairytale5571/ipcurrency/pkg/countries"
	"github.com/fairytale5571/ipcurrency/pkg/currencies"
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"github.com/fairytale5571/ipcurrency/pkg/ipinfo"
	"github.com/spf13/viper"
)

var _ repository.IPInfo = &Repository{}

type Repository struct {
	countryProvider    countries.Provider
	ipInfoProvider     ipinfo.Provider
	currenciesProvider currencies.Provider
}

func NewRepository() *Repository {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &Repository{
		countryProvider:    countries.NewCurrencies(client),
		ipInfoProvider:     ipinfo.NewIpInfo(client),
		currenciesProvider: currencies.NewCurrencies(client, viper.GetString("exchangeRateApiKey")),
	}
}

func (r Repository) GetIPInfo(ip string) (ipinfo.IPInformation, *errorops.Error) {
	return r.ipInfoProvider.RetrieveAddress(ip)
}

func (r Repository) GetCountryTickers(countryCode string) (map[string]countries.CurrencyInfo, *errorops.Error) {
	return r.countryProvider.GetCountryCurrencies(countryCode)
}

func (r Repository) GetCurrencyRate(countryCode string) ([]currencies.CurrencyRateToUAH, *errorops.Error) {

	tickers, err := r.GetCountryTickers(countryCode)
	if err != nil {
		return nil, err
	}

	var rates []currencies.CurrencyRateToUAH
	for k := range tickers {
		rate, err := r.currenciesProvider.GetExchangeUAHRates(k)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate...)
	}
	return rates, nil
}
