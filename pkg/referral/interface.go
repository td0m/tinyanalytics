package referral

import model "github.com/td0m/tinyanalytics"

// Service interface
type Service interface {
	GetReferrals(domain string, path string) ([]model.Referral, error)
}

// DB interface
type DB interface {
	ReferralsForSite(domain string) ([]model.Referral, error)
	ReferralsForPage(domain string, path string) ([]model.Referral, error)
}
