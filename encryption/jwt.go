package encryption

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/rarifbkn/api-rest-gorm-echo/internal/models"
)

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name":  u.Name,
	})

	return token.SignedString([]byte(key))
}

/*
func ParseLoginJWT(value string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

*/
