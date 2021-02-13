package visit

import (
	"net"
	"time"

	model "github.com/td0m/tinyanalytics"
)

// Service interface
type Service interface {
	VisitPage(IP net.HardwareAddr, userAgent string) error
	GetViews(page *model.Page, time time.Time, alltime bool) ([]model.ViewRow, error)
}

// DB interface
type DB interface {
	VisitOrCreatePage(*model.Visit) error

	SiteViewsAllTime(domain string) ([]model.ViewRow, error)
	PageViewsAllTime(page *model.Page) ([]model.ViewRow, error)
	SiteViewsInMonth(domain string, time time.Time) ([]model.ViewRow, error)
	PageViewsInMonth(page *model.Page, time time.Time) ([]model.ViewRow, error)
}

// CacheMap is used to store ip addresses in the short term
// used to prevent a single ip from simulating too many views
// this could also be done with session storage/cookies, but that's easier to bypass
type CacheMap interface {
	Store(ip string) bool
}

// UserAgentParser parses a HTTP user agent
type UserAgentParser interface {
	Parse(userAgent string) (model.Platform, model.Browser, error)
}
