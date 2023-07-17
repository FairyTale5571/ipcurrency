package ipinfo

import (
	"github.com/fairytale5571/ipcurrency/internal/api/repository"
	"github.com/fairytale5571/ipcurrency/internal/api/services"
	"github.com/fairytale5571/ipcurrency/pkg/currencies"
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"github.com/fairytale5571/ipcurrency/pkg/ipinfo"
)

var _ services.IPInfo = &Service{}

type Service struct {
	ipInfoRepo repository.IPInfo
}

func NewService(repo repository.IPInfo) *Service {
	return &Service{
		ipInfoRepo: repo,
	}
}

func (s *Service) GetIPInfo(ip string) (ipinfo.IPInformation, *errorops.Error) {
	return s.ipInfoRepo.GetIPInfo(ip)
}

func (s *Service) GetExchangeRate(countryCode string) ([]currencies.CurrencyRateToUAH, *errorops.Error) {
	return s.ipInfoRepo.GetCurrencyRate(countryCode)
}
