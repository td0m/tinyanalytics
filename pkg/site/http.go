package site

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/td0m/tinyanalytics/pkg/jwt"
)

// HTTP contains site HTTP endpoints
type HTTP struct {
	s Service
}

// NewHTTP creates a new HTTP handler
func NewHTTP(s Service) *HTTP { return &HTTP{s} }

func (h *HTTP) Create(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	email := jwt.FromContext(r.Context())
	site, err := h.s.CreateSite(domain, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(site)
}
