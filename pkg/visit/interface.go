package visit

import (
	"net"
	"time"

	model "github.com/td0m/tinyanalytics"
)

type Service interface {
	VisitPage(IP net.HardwareAddr, userAgent string) error
	GetViews(page *model.Page, time time.Time, alltime bool) ([]model.ViewRow, error)
}

type DB interface {
	VisitOrCreatePage(*model.Visit) error

	SiteViewsAllTime(domain string) ([]model.ViewRow, error)
	PageViewsAllTime(page *model.Page) ([]model.ViewRow, error)
	SiteViewsInMonth(domain string, time time.Time) ([]model.ViewRow, error)
	PageViewsInMonth(page *model.Page, time time.Time) ([]model.ViewRow, error)
}

type CacheMap interface {
	Store(ip string) bool
}

type UserAgentParser interface {
	Parse(userAgent string) (model.Platform, model.Browser, model.OS, error)
}
