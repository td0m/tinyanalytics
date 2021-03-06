package visit

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	model "github.com/td0m/tinyanalytics"
	"github.com/td0m/tinyanalytics/pkg/jwt"
)

// HTTP contains visit HTTP endpoints
type HTTP struct {
	s Service
}

// NewHTTP creates a new HTTP handler
func NewHTTP(s Service) *HTTP { return &HTTP{s} }

func (h *HTTP) Visit(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	ip := r.RemoteAddr
	domain := chi.URLParam(r, "domain")
	path := "/" + chi.URLParam(r, "*")
	referrer := r.URL.Query().Get("referrer")
	err := h.s.VisitPage(domain, path, ip, userAgent, referrer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *HTTP) ViewStats(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	period := query.Get("period")
	domain := query.Get("domain")
	path := query.Get("path")
	_, sitesOwned := jwt.FromContext(r.Context())
	if !contains(sitesOwned, domain) {
		http.Error(w, "no permission to view this site", http.StatusUnauthorized)
		return
	}

	rows, err := h.s.GetViews(model.Page{Domain: domain, Path: path}, "alltime" == period)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rows)
}

func contains(arr []string, a string) bool {
	for _, s := range arr {
		if s == a {
			return true
		}
	}
	return false
}
