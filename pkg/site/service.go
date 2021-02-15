package site

import (
	model "github.com/td0m/tinyanalytics"
)

// Service represents the site app service
type ServiceImpl struct {
	store           Store
	fileKeyProvider KeyProvider
	keyVerifier     KeyVerifier
}

// NewService creates a new service
func NewService(store Store) *ServiceImpl {
	return &ServiceImpl{
		store,
		NewFileKeyCrawler(),
		NewRandomKeyService(),
	}
}

func (s ServiceImpl) CreateSite(domain string, owner string) (*model.Site, error) {
	key, err := s.fileKeyProvider.GetKey(domain)
	if err != nil {
		return nil, err
	}

	if err := s.keyVerifier.Check(domain, key); err != nil {
		return nil, err
	}

	site, err := s.store.CreateSite(domain, owner)
	return site, err
}

func (s ServiceImpl) GetConfirmationKey(domain string) (key string, err error) {
	return s.keyVerifier.Generate(domain), nil
}
