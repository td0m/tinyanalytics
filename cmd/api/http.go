package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/td0m/tinyanalytics/pkg/user"
)

func initHTTP(svc *services) *chi.Mux {
	r := chi.NewRouter()

	anyUser := svc.jwt.Middleware()

	userH := user.NewHTTP(svc.user)

	r.Route("/api", func(api chi.Router) {
		api.Post("/signup", userH.SignUp)
		api.Post("/login", userH.Login)

		api.With(anyUser).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			claims := svc.jwt.FromContext(r.Context())
			fmt.Fprintln(w, "You are signed in as", claims)
		})
	})

	return r
}