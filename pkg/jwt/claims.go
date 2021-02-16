package jwt

import "github.com/form3tech-oss/jwt-go"

type Claims struct {
	jwt.StandardClaims
	Sites []string `json:"sites,omitempty"`
}
