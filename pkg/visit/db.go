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
	INSERT INTO visit(time,ip,domain,path,browser,platform)
	VALUES($1,$2,$3,$4,$5,$6) RETURNING *`
	createPage = `
	INSERT INTO page(domain,path)
	VALUES($1,$2)
	`
)

// NewDB creates a new visit store implementation in postgres
func NewDB(pool *pgxpool.Pool) *DB { return &DB{pool} }

func (db *DB) VisitOrCreatePage(v *model.Visit) error {
	_, err := db.pool.Exec(context.Background(), updateVisitors, v.Time, v.IP, v.Domain, v.Path, v.Browser, v.Platform)
	var pgxErr *pgconn.PgError
	// on foreign key violation, probably means that the page does not exist
	// TODO: is there a better way to check exactly what key is violated??
	if errors.As(err, &pgxErr) && pgxErr.Code == pgerrcode.ForeignKeyViolation {
		_, err := db.pool.Exec(context.Background(), createPage, v.Domain, v.Path)
		return err
	}
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

func (db *DB) PageViewsAllTime(page *model.Page) ([]model.ViewRow, error) {
	rows := []model.ViewRow{}
	err := pgxscan.Select(context.Background(), db.pool, &rows, selectVisits+byPage, page.Domain, page.Path)
	return rows, err
}

func (db *DB) SiteViewsInMonth(domain string) ([]model.ViewRow, error) {
	rows := []model.ViewRow{}
	err := pgxscan.Select(context.Background(), db.pool, &rows, selectVisitsInMonth+bySite+groupByMonth, domain)
	return rows, err
}

func (db *DB) PageViewsInMonth(page *model.Page) ([]model.ViewRow, error) {
	rows := []model.ViewRow{}
	err := pgxscan.Select(context.Background(), db.pool, &rows, selectVisitsInMonth+byPage+groupByMonth, page.Domain, page.Path)
	return rows, err
}
