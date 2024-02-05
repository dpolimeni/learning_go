package auth

// Set up JWT AUTH
import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateToken(username string) (map[string]string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 24).Unix(),
		"iat":      time.Now().Unix(),
	})
	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 24 * 60).Unix(),
		"iat":      time.Now().Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	access, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refresh, err := refreshtoken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access":  access,
		"refresh": refresh,
	}, nil
}
