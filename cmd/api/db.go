package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func initDB(url string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return conn, err
	}

	return conn, err
}
