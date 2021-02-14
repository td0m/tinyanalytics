package user

import (
	"context"

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
	getUserByEmail = `SELECT * FROM "user" WHERE email=$1 AND pass=crypt($2, pass)`
	createUser     = `INSERT INTO "user"(email, pass) VALUES($1,crypt($2, gen_salt('bf'))) RETURNING *`
)

// GetByEmailAndPassword gets a user by email, given that the passwords match
func (db *DB) GetByEmailAndPassword(email string, password string) (user *model.User, err error) {
	user = &model.User{}
	err = pgxscan.Get(context.Background(), db.pool, user, getUserByEmail, email, password)
	return
}

// CreateUser creates a new user
func (db *DB) CreateUser(email string, password string) (*model.User, error) {
	user := &model.User{}
	err := pgxscan.Get(context.Background(), db.pool, user, createUser, email, password)
	return user, err
}
