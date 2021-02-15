package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/td0m/tinyanalytics/pkg/jwt"
	"github.com/td0m/tinyanalytics/pkg/site"
	"github.com/td0m/tinyanalytics/pkg/user"
	"github.com/td0m/tinyanalytics/pkg/visit"
)

func initHTTP(svc *services) *chi.Mux {
	r := chi.NewRouter()

	anyUser := svc.jwt.Middleware()

	userH := user.NewHTTP(svc.user)
	siteH := site.NewHTTP(svc.site)
	visitH := visit.NewHTTP(svc.visit)

	r.Route("/api", func(api chi.Router) {
		r.Use(middleware.DefaultLogger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.SetHeader("Content-Type", "application/json"))

		api.Post("/signup", userH.SignUp)
		api.Post("/login", userH.Login)

		api.With(anyUser).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
			claims := jwt.FromContext(r.Context())
			fmt.Fprintln(w, "You are signed in as", claims)
		})

		api.Get("/verification-code", siteH.GetConfirmationKey)
		api.With(anyUser).Put("/sites/{domain}", siteH.Create)
		api.With(middleware.RealIP).Post("/visit/{domain}/*", visitH.Visit)
		// TODO: make this "owner" not "anyUser"
		api.With(anyUser).Get("/views", visitH.ViewStats)
	})

	return r
}
