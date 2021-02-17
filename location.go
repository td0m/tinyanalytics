package model

import "fmt"

// Location type (still undetermined for now)
type Location struct {
	Longitude float64
	Latitude  float64
}

// PostgisString formats the location to a postgis longitude format
func (l Location) PostgisString() string {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", l.Longitude, l.Latitude)
}
