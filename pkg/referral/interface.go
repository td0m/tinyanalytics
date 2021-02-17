package referral

import model "github.com/td0m/tinyanalytics"

// Service interface
type Service interface {
	GetReferrals(domain string, path string) ([]model.Referral, error)
}

// Store interface
type Store interface {
	ReferralsForSite(domain string) ([]model.Referral, error)
	ReferralsForPage(domain string, path string) ([]model.Referral, error)
}
