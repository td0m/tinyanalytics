package main

import (
	"github.com/td0m/tinyanalytics/pkg/page"
	"github.com/td0m/tinyanalytics/pkg/referral"
	"github.com/td0m/tinyanalytics/pkg/site"
	"github.com/td0m/tinyanalytics/pkg/user"
	"github.com/td0m/tinyanalytics/pkg/visit"
)

type services struct {
	referral referral.Service
	user     user.Service
	page     page.Service
	site     site.Service
	visit    visit.Service
}

// func initServices(db *pgxpool.Pool, secret string) *services {
// 	jwt := jwt.NewService(secret)
// 	return &services{
// 		referral: referral.NewService(),
// 		user:     user.NewService(user.NewDB(db), jwt),
// 		page:     page.NewService(),
// 		site:     site.NewService(),
// 		visit:    visit.NewService(),
// 	}
// }
