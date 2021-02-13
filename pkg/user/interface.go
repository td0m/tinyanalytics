package user

type Service interface {
	Login(email, password string) (jwt string, err error)
	SignUp(email, password string) (jwt string, err error)

	DeleteUser(email string) error
}

type DB interface {
	CreateUser(email, password string) error
	GetByEmailAndPassword(email, password string) error
}

type JWTGenerator interface {
	Generate(string) (string, error)
}
