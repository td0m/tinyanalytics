package page

import model "github.com/td0m/tinyanalytics"

// Service interface
type Service interface {
	GetPage(domain, path string) (*model.Page, error)
}

// DB interface
type DB interface {
	GetPage(domain, path string) (model.Page, error)
}
