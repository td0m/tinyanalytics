package referral

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	model "github.com/td0m/tinyanalytics"
)

// DB represents the referral postgres database
type DB struct {
	pool *pgxpool.Pool
}

// NewDB creates a new referral postgres database
func NewDB(pool *pgxpool.Pool) *DB { return &DB{pool} }

const (
	getReferrals = `
	SELECT
		from_domain AS "from.domain",
		from_path AS "from.path",
		SUM(count) AS "count"
	FROM referral`
	groupAndOrder = `
		GROUP BY from_domain, from_path
		ORDER BY count DESC
	`
	refsByDomain = getReferrals + ` WHERE to_domain=$1` + groupAndOrder
	refsByPage   = getReferrals + ` WHERE to_domain=$1 AND to_path=$2` + groupAndOrder
)

func (db *DB) ReferralsForSite(domain string) ([]model.Referral, error) {
	refs := []model.Referral{}
	err := pgxscan.Select(context.Background(), db.pool, &refs, refsByDomain, domain)
	return refs, err
}

func (db *DB) ReferralsForPage(domain string, path string) ([]model.Referral, error) {
	refs := []model.Referral{}
	err := pgxscan.Select(context.Background(), db.pool, &refs, refsByPage, domain, path)
	return refs, err
}
