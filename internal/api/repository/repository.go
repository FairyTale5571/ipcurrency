package repository

import (
	"github.com/fairytale5571/ipcurrency/pkg/countries"
	"github.com/fairytale5571/ipcurrency/pkg/currencies"
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"github.com/fairytale5571/ipcurrency/pkg/ipinfo"
)

type (
	IPInfo interface {
		GetIPInfo(ip string) (ipinfo.IPInformation, *errorops.Error)
		GetCountryTickers(countryCode string) (map[string]countries.CurrencyInfo, *errorops.Error)
		GetCurrencyRate(countryCode string) ([]currencies.CurrencyRateToUAH, *errorops.Error)
	}
)
