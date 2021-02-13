package site

// Service interface
type Service interface {
	CreateSite(domain, owner string) error
}

// DB interface
type DB interface {
	CreateSite(domain, owner string) error
}
