package user

import (
	"encoding/json"
	"net/http"
)

// HTTP contains user HTTP endpoints
type HTTP struct {
	s Service
}

// NewHTTP creates a new HTTP handler
func NewHTTP(s Service) *HTTP { return &HTTP{s} }

type loginBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// Login endpoint
func (h *HTTP) Login(w http.ResponseWriter, r *http.Request) {
	body := &loginBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, user, err := h.s.Login(body.Email, body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func (h *HTTP) SignUp(w http.ResponseWriter, r *http.Request) {
	body := &loginBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, user, err := h.s.SignUp(body.Email, body.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}
