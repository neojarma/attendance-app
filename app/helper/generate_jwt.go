package helper

import (
	"os"
	"presensi/model"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(user *model.EmployeesJoinRole, expires time.Time) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    user.Nip,
		"name":      user.Name,
		"role":      user.Role,
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": expires.Unix(),
	})

	tokenStr, err := claims.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
