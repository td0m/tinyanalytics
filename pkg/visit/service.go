package visit

import (
	"net"
	"time"

	model "github.com/td0m/tinyanalytics"
)

// ServiceImpl represents the visit app service
type ServiceImpl struct {
	store Store
}

// NewService creates a new visit app service
func NewService(store Store) *ServiceImpl { return &ServiceImpl{store} }

func (s *ServiceImpl) VisitPage(domain, path, ip, userAgent string) error {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		parsedIP = []byte{0, 0, 0, 0}
	}
	visit := &model.Visit{
		Time:     time.Now().Round(time.Second),
		Domain:   domain,
		Path:     path,
		IP:       parsedIP,
		Platform: model.PlatformUnknown,
		Browser:  model.BrowserUnknown,
	}
	if err := visit.Validate(); err != nil {
		return err
	}
	return s.store.VisitOrCreatePage(visit)
}

func (s *ServiceImpl) GetViews(page *model.Page, alltime bool) ([]model.ViewRow, error) {
	var (
		rows []model.ViewRow
		err  error
	)
	if err := page.Validate(); err != nil {
		return rows, err
	}

	// TODO: what about the index page? that will always have a length of 0
	// maybe use * or - symbol?
	switch alltime {
	case true:
		switch len(page.Path) {
		case 0:
			rows, err = s.store.SiteViewsAllTime(page.Domain)
		default:
			rows, err = s.store.PageViewsAllTime(page)
		}
	default:
		switch len(page.Path) {
		case 0:
			rows, err = s.store.SiteViewsInMonth(page.Domain)
		default:
			rows, err = s.store.PageViewsInMonth(page)
		}
	}

	return rows, err
}
