package page

import model "github.com/td0m/tinyanalytics"

type Service interface {
	GetPage(domain, path string) (*model.Page, error)
}

type DB interface {
	GetPage(domain, path string) (model.Page, error)
}
