package user

import model "github.com/td0m/tinyanalytics"

// Service interface
type Service interface {
	Login(email, password string) (jwt string, user *model.UserWithSites, err error)
	SignUp(email, password string) (jwt string, user *model.User, err error)

	DeleteUser(email string) error
}

// Store interface
type Store interface {
	CreateUser(email, password string) (*model.User, error)
	GetByEmailAndPassword(email, password string) (*model.UserWithSites, error)
}

// JWTGenerator is responsible for generating a JWT token from a unique identifier
type JWTGenerator interface {
	Generate(string, []string) (string, error)
}
