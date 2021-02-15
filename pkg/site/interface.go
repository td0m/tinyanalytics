package site

import model "github.com/td0m/tinyanalytics"

// Service interface
type Service interface {
	CreateSite(domain, owner string) (*model.Site, error)
	GetConfirmationKey(domain string) (key string, err error)
}

// Store interface
type Store interface {
	CreateSite(domain, owner string) (*model.Site, error)
}

// KeyProvider crawls the key from a domain
type KeyProvider interface {
	GetKey(domain string) (key string, err error)
}

// KeyVerifier allows to generate and verify a domain key
type KeyVerifier interface {
	Generate(domain string) (key string)
	Check(domain string, key string) (err error)
}
