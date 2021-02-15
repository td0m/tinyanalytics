package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

var (
	port        string
	databaseURL string
	jwtSecret   string
)

func init() {
	godotenv.Load()
	port = get("PORT", "8080")
	databaseURL = get("DATABASE_URL")
	jwtSecret = get("JWT_SECRET")
}

func main() {
	db, err := initDB(databaseURL)
	check(err)

	svc := initServices(db, jwtSecret)
	r := initHTTP(svc)

	check(http.ListenAndServe(":"+port, r))
}
