package ipinfo

import (
	"github.com/fairytale5571/ipcurrency/pkg/errorops"
	"net"
	"net/http"
)

func (r requestBody) validate() *errorops.Error {
	if len(r.IP) == 0 {
		return errorops.NewError(
			http.StatusBadRequest,
			"IP list is empty",
			nil,
		)
	}
	if len(r.IP) > 10 {
		return errorops.NewError(
			http.StatusBadRequest,
			"IP list is too long",
			nil,
			"max 10 IP addresses allowed",
		)
	}

	var invalidIPs []string
	for _, ip := range r.IP {
		if err := net.ParseIP(ip); err == nil {
			invalidIPs = append(invalidIPs, ip)
		}
	}
	if len(invalidIPs) > 0 {
		return errorops.NewError(
			http.StatusBadRequest,
			"IP list contains invalid IP addresses",
			invalidIPs,
		)
	}

	return nil
}
