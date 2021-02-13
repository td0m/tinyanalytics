package model

import (
	"net"
	"time"
)

// Visit model
type Visit struct {
	Time   time.Time
	IP     net.HardwareAddr
	Domain string
	Path   string

	Views int

	Geo      Location
	Browser  Browser
	Platform Platform
}

// ViewRow is a subset of Visit
type ViewRow struct {
	Time  time.Time
	Views int
}
