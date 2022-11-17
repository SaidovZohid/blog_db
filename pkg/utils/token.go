package utils

import (
	"time"
	"github.com/golang-jwt/jwt"
)

const SecretKey = "supersecretkey"

func CreateToken(username, email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, email, duration)
	if err != nil {
		return "", nil, err
	}
	
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(SecretKey))
	return token, payload,  err
}