package visit

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
	model "github.com/td0m/tinyanalytics"
)

// DB represents the visit store implementation in postgres
type DB struct {
	pool *pgxpool.Pool
}

const (
	updateVisitors = `
	INSERT INTO visit(time,ip,domain,path,browser,platform,geo)
	VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING *`
	createPage = `
	INSERT INTO page(domain,path)
	VALUES($1,$2)
	`
	refer = `
	INSERT INTO referral(from_domain, from_path, to_domain, to_path, count)
	VALUES($1,$2,$3,$4,1)
	ON CONFLICT(from_domain,from_path,to_domain,to_path) DO UPDATE SET count=referral.count+1
	`
)

// NewDB creates a new visit store implementation in postgres
func NewDB(pool *pgxpool.Pool) *DB { return &DB{pool} }

func (db *DB) VisitOrCreatePage(v model.Visit, from model.Page) error {
	tx, _ := db.pool.Begin(context.Background())
	defer tx.Rollback(context.Background())
	_, err := tx.Exec(context.Background(), updateVisitors, v.Time, v.IP, v.Domain, v.Path, v.Browser, v.Platform, v.Geo.PostgisString())
	var pgxErr *pgconn.PgError
	// on foreign key violation, probably means that the page does not exist
	// TODO: is there a better way to check exactly what key is violated??
	if errors.As(err, &pgxErr) && pgxErr.Code == pgerrcode.ForeignKeyViolation {
		tx, _ = db.pool.Begin(context.Background())
		_, err := db.pool.Exec(context.Background(), createPage, v.Domain, v.Path)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), refer, from.Domain, from.Path, v.Domain, v.Path)
	if err != nil {
		return err
	}
	err = tx.Commit(context.Background())
	return err
}

const (
	selectVisits        = `SELECT COUNT(*) AS views FROM visit`
	selectVisitsInMonth = `
		SELECT COUNT(*) AS views,
			time_bucket('30 days', time) AS "time_bucket"
		FROM visit`
	bySite       = ` WHERE domain=$1`
	byPage       = bySite + ` AND path=$2`
	groupByMonth = ` GROUP BY time_bucket`
)

func (db *DB) SiteViewsAllTime(domain string) ([]model.ViewRow, error) {
	rows := []model.ViewRow{}
	err := pgxscan.Select(context.Background(), db.pool, &rows, selectVisits+bySite, domain)
	return rows, err
}

func (db *DB) PageViewsAllTime(page model.Page) ([]model.ViewRow, error) {
	rows := []model.ViewRow{}
	err := pgxscan.Select(context.Background(), db.pool, &rows, selectVisits+byPage, page.Domain, page.Path)
	return rows, err
}

func (db *DB) SiteViewsInMonth(domain string) ([]model.ViewRow, error) {
	rows := []model.ViewRow{}
	err := pgxscan.Select(context.Background(), db.pool, &rows, selectVisitsInMonth+bySite+groupByMonth, domain)
	return rows, err
}

func (db *DB) PageViewsInMonth(page model.Page) ([]model.ViewRow, error) {
	rows := []model.ViewRow{}
	err := pgxscan.Select(context.Background(), db.pool, &rows, selectVisitsInMonth+byPage+groupByMonth, page.Domain, page.Path)
	return rows, err
}
