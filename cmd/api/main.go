package main

import (
	"net/http"

	"github.com/joho/godotenv"
)

var (
	port        string
	databaseURL string
	jwtSecret   string
	mmdb        string
)

func init() {
	godotenv.Load()
	port = get("PORT", "8080")
	databaseURL = get("DATABASE_URL")
	jwtSecret = get("JWT_SECRET")
	mmdb = get("MMDB_FILE")
}

func main() {
	db, err := initDB(databaseURL)
	check(err)

	svc := initServices(db, jwtSecret, mmdb)
	r := initHTTP(svc)

	check(http.ListenAndServe(":"+port, r))
}
