package site

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	model "github.com/td0m/tinyanalytics"
)

// DB struct
type DB struct {
	pool *pgxpool.Pool
}

// NewDB creates a new DB
func NewDB(pool *pgxpool.Pool) *DB { return &DB{pool} }

const (
	insertSite = `INSERT INTO site(domain, owner) VALUES($1,$2) RETURNING *`
)

func (db *DB) CreateSite(domain string, owner string) (*model.Site, error) {
	site := &model.Site{}
	err := pgxscan.Get(context.Background(), db.pool, site, insertSite, domain, owner)
	return site, err
}
