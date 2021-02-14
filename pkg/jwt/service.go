package jwt

import (
	"context"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

const (
	expiresAfter = time.Hour * 24 * 7
)

var (
	method = jwt.SigningMethodHS256
)

type Service struct {
	secret []byte
}

func New(secret string) *Service {
	return &Service{[]byte(secret)}
}

// Generate generates a new jwt
func (s *Service) Generate(email string) (string, error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   email,
			ExpiresAt: time.Now().Add(expiresAfter).Unix(),
		},
	}
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(s.secret)
}

func (s *Service) Middleware() func(http.Handler) http.Handler {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return s.secret, nil
		},
		SigningMethod: method,
	}).Handler
}

func (s *Service) FromContext(ctx context.Context) string {
	token := ctx.Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return claims["sub"].(string)
}
