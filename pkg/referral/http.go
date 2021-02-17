package referral

import (
	"encoding/json"
	"net/http"
)

// HTTP represents the referral http endpoints
type HTTP struct {
	s Service
}

// NewHTTP creates a new referral http endpoints
func NewHTTP(s Service) *HTTP { return &HTTP{s} }

func (h *HTTP) GetReferrals(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	domain, path := query.Get("domain"), query.Get("path")
	refs, err := h.s.GetReferrals(domain, path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(refs)
}
