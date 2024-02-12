package auth

// Set up JWT AUTH
import (
	"context"
	"os"
	"time"

	"github.com/dpolimeni/fiber_app/ent"
	"github.com/dpolimeni/fiber_app/ent/user"
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
	secret := os.Getenv("SECRET_KEY")
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

func CheckSuperuser(DbClient *ent.Client, username string) bool {
	user, err := DbClient.User.Query().Where(user.UsernameEQ(username)).First(context.Background())
	if err != nil {
		return false
	}
	if user.IsAdmin {
		return true
	} else {
		return false
	}
}
