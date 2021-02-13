package site

type Service interface {
	CreateSite(domain, owner string) error
}

type DB interface {
	CreateSite(domain, owner string) error
}
