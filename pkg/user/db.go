package user

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	model "github.com/td0m/tinyanalytics"
)

// DB instance
type DB struct {
	pool *pgxpool.Pool
}

var _ Store = &DB{}

// NewDB creates a new DB
func NewDB(pool *pgxpool.Pool) *DB { return &DB{pool} }

const (
	getUserByEmail  = `SELECT * FROM "user" WHERE email=$1 AND pass=crypt($2, pass)`
	getSitesByOwner = `SELECT domain FROM site WHERE owner=$1`
	createUser      = `INSERT INTO "user"(email, pass) VALUES($1,crypt($2, gen_salt('bf'))) RETURNING *`
)

// GetByEmailAndPassword gets a user by email, given that the passwords match
func (db *DB) GetByEmailAndPassword(email string, password string) (user *model.UserWithSites, err error) {
	user = &model.UserWithSites{}
	tx, _ := db.pool.Begin(context.Background())
	defer tx.Commit(context.Background())
	err = pgxscan.Get(context.Background(), tx, user, getUserByEmail, email, password)
	if err != nil {
		return
	}
	sites := []string{}
	// this can error since there can be no sites
	err = pgxscan.Select(context.Background(), tx, &sites, getSitesByOwner, email)
	user.Sites = sites
	err = tx.Commit(context.Background())
	return
}

// CreateUser creates a new user
func (db *DB) CreateUser(email string, password string) (*model.User, error) {
	fmt.Println(email, password, "g")
	user := &model.User{}
	err := pgxscan.Get(context.Background(), db.pool, user, createUser, email, password)
	return user, err
}
