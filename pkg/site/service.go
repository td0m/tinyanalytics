package site

import model "github.com/td0m/tinyanalytics"

// Service represents the site app service
type ServiceImpl struct {
	store Store
}

// NewService creates a new service
func NewService(store Store) *ServiceImpl { return &ServiceImpl{store} }

func (s ServiceImpl) CreateSite(domain string, owner string) (*model.Site, error) {
	site, err := s.store.CreateSite(domain, owner)
	return site, err
}
