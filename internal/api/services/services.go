package services

import (
	"github.com/fairytale5571/ipcurrency/pkg/currencies"
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"github.com/fairytale5571/ipcurrency/pkg/ipinfo"
)

type (
	IPInfo interface {
		GetIPInfo(ip string) (ipinfo.IPInformation, *errorops.Error)
		GetExchangeRate(countryCode string) ([]currencies.CurrencyRateToUAH, *errorops.Error)
	}
)
