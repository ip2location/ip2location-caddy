package ip2locationcaddy

import (
	"fmt"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/ip2location/ip2location-io-go/ip2locationio"
	"net/http"
)

func (h *IP2LocationCaddy) setDBGeoHeaders(r *http.Request, geo ip2location.IP2Locationrecord) {
	prefix := h.HeaderPrefix

	r.Header.Set(prefix+"-Country-Code", geo.Country_short)
	r.Header.Set(prefix+"-Country-Name", geo.Country_long)
	r.Header.Set(prefix+"-Region", geo.Region)
	r.Header.Set(prefix+"-City", geo.City)
	r.Header.Set(prefix+"-ISP", geo.Isp)
	r.Header.Set(prefix+"-Latitude", fmt.Sprintf("%.6f", geo.Latitude))
	r.Header.Set(prefix+"-Longitude", fmt.Sprintf("%.6f", geo.Longitude))
	r.Header.Set(prefix+"-Domain", geo.Domain)
	r.Header.Set(prefix+"-ZIP-Code", geo.Zipcode)
	r.Header.Set(prefix+"-Time-Zone", geo.Timezone)
	r.Header.Set(prefix+"-Net-Speed", geo.Netspeed)
	r.Header.Set(prefix+"-IDD-Code", geo.Iddcode)
	r.Header.Set(prefix+"-Area-Code", geo.Areacode)
	r.Header.Set(prefix+"-Weather-Station-Code", geo.Weatherstationcode)
	r.Header.Set(prefix+"-Weather-Station-Name", geo.Weatherstationname)
	r.Header.Set(prefix+"-MCC", geo.Mcc)
	r.Header.Set(prefix+"-MNC", geo.Mnc)
	r.Header.Set(prefix+"-Mobile-Brand", geo.Mobilebrand)
	r.Header.Set(prefix+"-Elevation", fmt.Sprintf("%.0f", geo.Elevation))
	r.Header.Set(prefix+"-Usage-Type", geo.Usagetype)
	r.Header.Set(prefix+"-Address-Type", geo.Addresstype)
	r.Header.Set(prefix+"-Category", geo.Category)
	r.Header.Set(prefix+"-District", geo.District)
	r.Header.Set(prefix+"-ASN", geo.Asn)
	r.Header.Set(prefix+"-AS", geo.As)
	r.Header.Set(prefix+"-AS-Domain", geo.Asdomain)
	r.Header.Set(prefix+"-AS-Usage-Type", geo.Asusagetype)
	r.Header.Set(prefix+"-AS-CIDR", geo.Ascidr)
}

func (h *IP2LocationCaddy) setAPIGeoHeaders(r *http.Request, geo ip2locationio.IPGeolocationResult) {
	prefix := h.HeaderPrefix

	// currently, we are matching the fields returned by the DB case
	r.Header.Set(prefix+"-Country-Code", geo.CountryCode)
	r.Header.Set(prefix+"-Country-Name", geo.CountryName)
	r.Header.Set(prefix+"-Region", geo.RegionName)
	r.Header.Set(prefix+"-City", geo.CityName)
	r.Header.Set(prefix+"-ISP", geo.Isp)
	r.Header.Set(prefix+"-Latitude", fmt.Sprintf("%.6f", geo.Latitude))
	r.Header.Set(prefix+"-Longitude", fmt.Sprintf("%.6f", geo.Longitude))
	r.Header.Set(prefix+"-Domain", geo.Domain)
	r.Header.Set(prefix+"-ZIP-Code", geo.ZipCode)
	r.Header.Set(prefix+"-Time-Zone", geo.TimeZone)
	r.Header.Set(prefix+"-Net-Speed", geo.NetSpeed)
	r.Header.Set(prefix+"-IDD-Code", geo.IddCode)
	r.Header.Set(prefix+"-Area-Code", geo.AreaCode)
	r.Header.Set(prefix+"-Weather-Station-Code", geo.WeatherStationCode)
	r.Header.Set(prefix+"-Weather-Station-Name", geo.WeatherStationName)
	r.Header.Set(prefix+"-MCC", geo.Mcc)
	r.Header.Set(prefix+"-MNC", geo.Mnc)
	r.Header.Set(prefix+"-Mobile-Brand", geo.MobileBrand)
	r.Header.Set(prefix+"-Elevation", fmt.Sprintf("%d", geo.Elevation))
	r.Header.Set(prefix+"-Usage-Type", geo.UsageType)
	r.Header.Set(prefix+"-Address-Type", geo.AddressType)
	r.Header.Set(prefix+"-Category", geo.AdsCategory)
	r.Header.Set(prefix+"-District", geo.District)
	r.Header.Set(prefix+"-ASN", geo.Asn)
	r.Header.Set(prefix+"-AS", geo.AS)
	r.Header.Set(prefix+"-AS-Domain", geo.ASInfo.ASDomain)
	r.Header.Set(prefix+"-AS-Usage-Type", geo.ASInfo.ASUsageType)
	r.Header.Set(prefix+"-AS-CIDR", geo.ASInfo.ASCidr)
}
