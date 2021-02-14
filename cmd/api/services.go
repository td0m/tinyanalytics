package main

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/td0m/tinyanalytics/pkg/jwt"
	"github.com/td0m/tinyanalytics/pkg/page"
	"github.com/td0m/tinyanalytics/pkg/referral"
	"github.com/td0m/tinyanalytics/pkg/site"
	"github.com/td0m/tinyanalytics/pkg/user"
	"github.com/td0m/tinyanalytics/pkg/visit"
)

type services struct {
	jwt      *jwt.Service
	referral referral.Service
	user     user.Service
	page     page.Service
	site     site.Service
	visit    visit.Service
}

func initServices(db *pgxpool.Pool, secret string) *services {
	jwtS := jwt.New(secret)
	return &services{
		jwt: jwtS,
		// referral: referral.NewService(),
		user: user.NewService(user.NewDB(db), jwtS),
		// page:     page.NewService(),
		// site:     site.NewService(),
		// visit:    visit.NewService(),
	}
}
