package site

import model "github.com/td0m/tinyanalytics"

// Service interface
type Service interface {
	CreateSite(domain, owner string) (*model.Site, error)
}

// Store interface
type Store interface {
	CreateSite(domain, owner string) (*model.Site, error)
}
