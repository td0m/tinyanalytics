package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

var (
	port        string
	databaseURL string
	secret      string
)

func init() {
	godotenv.Load()
	port = get("PORT", "8080")
	databaseURL = get("DATABASE_URL")
	secret = get("JWT_SECRET")
}

func main() {
	// db, err := initDB(databaseURL)
	// check(err)

	// svc := initServices(db, secret)

	check(http.ListenAndServe(":"+port, nil))
}
