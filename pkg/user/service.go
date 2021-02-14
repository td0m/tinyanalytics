package user

import (
	"errors"
	"log"

	model "github.com/td0m/tinyanalytics"
)

// Service represents the user app service
type ServiceImpl struct {
	db  Store
	jwt JWTGenerator
}

var _ Service = &ServiceImpl{}

// NewService creates a new service
func NewService(db Store, jwt JWTGenerator) *ServiceImpl { return &ServiceImpl{db, jwt} }

// errors
var (
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrFailedtoCreate         = errors.New("Email might already be in use, or an internal error occured")
)

func (s *ServiceImpl) Login(email string, password string) (jwt string, u *model.User, err error) {
	u = &model.User{Email: email, Pass: password}
	if err = u.Validate(); err != nil {
		return
	}
	u, err = s.db.GetByEmailAndPassword(u.Email, u.Pass)
	if err != nil {
		err = ErrInvalidEmailOrPassword
		return
	}
	jwt, err = s.jwt.Generate(u.Email)
	return
}

func (s *ServiceImpl) SignUp(email string, password string) (jwt string, u *model.User, err error) {
	u = &model.User{Email: email, Pass: password}
	if err = u.Validate(); err != nil {
		return
	}
	u, err = s.db.CreateUser(u.Email, u.Pass)
	if err != nil {
		log.Println(err)
		return "", u, ErrFailedtoCreate
	}
	jwt, err = s.jwt.Generate(u.Email)
	return
}

func (s *ServiceImpl) DeleteUser(email string) error {
	panic("not implemented") // TODO: Implement
}
