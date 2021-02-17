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
	email, _ := jwt.FromContext(r.Context())
	site, err := h.s.CreateSite(domain, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(site)
}

func (h *HTTP) GetConfirmationKey(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	key, err := h.s.GetConfirmationKey(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"key": key,
	})
}
