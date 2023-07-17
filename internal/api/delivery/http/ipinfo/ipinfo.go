package ipinfo

import (
	"net/http"
	"time"

	"github.com/fairytale5571/ipcurrency/internal/api/services"
	"github.com/fairytale5571/ipcurrency/pkg/currencies"
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"github.com/fairytale5571/ipcurrency/pkg/timeops"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type requestBody struct {
	IP []string `json:"ip"`
}

type responseBody struct {
	IP          string                         `json:"ip"`
	Country     string                         `json:"country"`
	City        string                         `json:"city"`
	Latitude    float64                        `json:"latitude"`
	Longitude   float64                        `json:"longitude"`
	CurrentTime string                         `json:"currentTime"`
	Currencies  []currencies.CurrencyRateToUAH `json:"currencies"`
}

type Handler struct {
	ipInfoService services.IPInfo
	tmFn          func() time.Time
}

// GetIPInfo
// @Summary get ip info
// @Schemes
// @Description return country information and currency rate to UAH
// @Produce  	json
// @Param request body 			requestBody 		true "query params"
// @Success 	200 			{object} 			responseBody
// @Failure 	400 			{object} 			errorops.Error
// @Router /ip-info [post]
func (h Handler) GetIPInfo() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var request requestBody
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorops.NewError(
				http.StatusInternalServerError,
				"failed to bind request body",
				err,
			))
			return
		}
		if err := request.validate(); err != nil {
			ctx.JSON(err.Code, err)
			return
		}
		response, err := h.getIPInfoInternal(request)
		if err != nil {
			ctx.JSON(err.Code, err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func (h Handler) getIPInfoInternal(request requestBody) ([]responseBody, *errorops.Error) {
	var errGroup errgroup.Group

	response := make([]responseBody, len(request.IP))
	for i, ip := range request.IP {
		i, ip := i, ip
		errGroup.Go(func() error {
			ipInfo, err := h.ipInfoService.GetIPInfo(ip)
			if err != nil {
				return err
			}

			rate, err := h.ipInfoService.GetExchangeRate(ipInfo.CountryCode)
			if err != nil {
				return err
			}

			loc, e := time.LoadLocation(ipInfo.Timezone)
			if e != nil {
				return errorops.NewError(
					http.StatusInternalServerError,
					"failed to load location",
					ipInfo.Timezone,
					e.Error(),
				)
			}

			response[i] = responseBody{
				IP:          ip,
				Country:     ipInfo.Country,
				City:        ipInfo.City,
				Latitude:    ipInfo.Lat,
				Longitude:   ipInfo.Lon,
				CurrentTime: h.tmFn().In(loc).Format(timeops.DD_MM_YYYY_HH_MM),
				Currencies:  rate,
			}
			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return []responseBody{}, errorops.NewError(
			http.StatusInternalServerError,
			"failed to get IP info",
			err,
		)
	}
	return response, nil
}

func NewHandler(tmFn func() time.Time, service services.IPInfo) *Handler {
	return &Handler{
		tmFn:          tmFn,
		ipInfoService: service,
	}
}
