package user

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	model "github.com/td0m/tinyanalytics"
)

type DBImpl struct {
	pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *DBImpl { return &DBImpl{pool} }

const (
	getUserByEmail = `SELECT * FROM "user" WHERE email=$1 AND pass=crypt($2, pass)`
	createUser     = `INSERT INTO "user"(email, pass) VALUES($1,crypt($2, gen_salt('bf')))`
)

// GetUser gets a user by email, given that the passwords match
func (db *DBImpl) GetUser(email string, password string) (user *model.User, err error) {
	user = &model.User{}
	err = pgxscan.Get(context.Background(), db.pool, user, getUserByEmail, email, password)
	return
}

// CreateUser creates a new user
func (db *DBImpl) CreateUser(email string, password string) error {
	_, err := db.pool.Exec(context.Background(), createUser, email, password)
	return err
}
