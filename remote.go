package ip2locationcaddy

import (
	"github.com/ip2location/ip2location-io-go/ip2locationio"
)

func (h *IP2LocationCaddy) initRemote() error {
	config, err := ip2locationio.OpenConfiguration(h.APIKey)

	if err != nil {
		return err
	}
	ipl, err := ip2locationio.OpenIPGeolocation(config)

	if err != nil {
		return err
	}
	h.api = ipl
	return nil
}

func (h *IP2LocationCaddy) lookupRemote(ip string) (ip2locationio.IPGeolocationResult, error) {
	res, err := h.api.LookUp(ip, "") // won't support lang param

	if err != nil {
		return res, err
	}

	return res, nil
}
