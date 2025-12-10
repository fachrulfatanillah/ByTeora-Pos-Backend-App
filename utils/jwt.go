package utils

import (
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(uuid, email, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "uuid":  uuid,
        "email": email,
        "role":  role,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // 1 day
    })

    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}