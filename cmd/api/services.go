package main

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/td0m/tinyanalytics/pkg/cache"
	"github.com/td0m/tinyanalytics/pkg/jwt"
	"github.com/td0m/tinyanalytics/pkg/page"
	"github.com/td0m/tinyanalytics/pkg/referral"
	"github.com/td0m/tinyanalytics/pkg/site"
	"github.com/td0m/tinyanalytics/pkg/user"
	"github.com/td0m/tinyanalytics/pkg/useragent"
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

func initServices(db *pgxpool.Pool, secret, mmdb string) *services {
	jwtS := jwt.New(secret)
	ipCache := cache.NewMap(time.Hour * 24)
	uap := useragent.NewParser()
	locator, err := visit.NewGeoLocator(mmdb)
	check(err)
	return &services{
		jwt:      jwtS,
		referral: referral.NewService(referral.NewDB(db)),
		user:     user.NewService(user.NewDB(db), jwtS),
		// page:     page.NewService(),
		site:  site.NewService(site.NewDB(db)),
		visit: visit.NewService(visit.NewDB(db), ipCache, uap, locator),
	}
}
