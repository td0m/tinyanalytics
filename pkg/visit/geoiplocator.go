package visit

import (
	"net"

	"github.com/oschwald/geoip2-golang"
	model "github.com/td0m/tinyanalytics"
)

type GeoLocator struct {
	*geoip2.Reader
}

func NewGeoLocator(file string) (*GeoLocator, error) {
	db, err := geoip2.Open(file)
	if err != nil {
		return nil, err
	}
	return &GeoLocator{db}, nil
}

// Locate returns appriximate latitude and logitude location
func (g *GeoLocator) Locate(ip net.IP) (*model.Location, error) {
	record, err := g.City(ip)
	if err != nil {
		return nil, err
	}
	return &model.Location{Latitude: record.Location.Latitude, Longitude: record.Location.Longitude}, nil
}
