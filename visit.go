package model

import (
	"net"
	"time"
)

// Visit model
type Visit struct {
	Time   time.Time
	IP     net.IP
	Domain string
	Path   string

	Geo      Location
	Browser  Browser
	Platform Platform
}

// TODO: implement
func (v Visit) Validate() error {
	return nil
}

// ViewRow is a subset of Visit
type ViewRow struct {
	Time  time.Time
	Views int
}
