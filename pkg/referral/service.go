package referral

import model "github.com/td0m/tinyanalytics"

// ServiceImpl represents the referral app service
type ServiceImpl struct {
	store Store
}

// NewServiceImpl creates a new referral app service
func NewService(store Store) *ServiceImpl { return &ServiceImpl{store} }

func (s *ServiceImpl) GetReferrals(domain, path string) ([]model.Referral, error) {
	if len(path) > 0 {
		return s.store.ReferralsForPage(domain, path)
	}
	return s.store.ReferralsForSite(domain)
}
