package ip2locationcaddy

import (
	"github.com/ip2location/ip2location-go/v9"
)

func (h *IP2LocationCaddy) initLocal() error {
	// Open IP2Location BIN database here.
	db, err := ip2location.OpenDB(h.BinPath)
	if err != nil {
		return err
	}
	h.db = db

	return nil
}

func (h *IP2LocationCaddy) lookupLocal(ip string) (ip2location.IP2Locationrecord, error) {
	// Query local BIN here.
	results, err := h.db.Get_all(ip)

	if err != nil {
		return results, err
	}
	return results, nil
}
