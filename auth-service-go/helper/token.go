package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func CreateToken(username string, secretKey string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "auth",
			ExpiresAt: time.Now().Add(time.Duration(1000) * time.Hour).Unix(),
		},
		Username: username,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func parseTokenJwt(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("signing method invalid")
		}

		return []byte(secretKey), nil
	})

	return token, err
}

func ParseToken(tokenString string, secretKey string) error {
	token, err := parseTokenJwt(tokenString, secretKey)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("token invalid")
	}

	// cek if claims is empty
	if claims["exp"] == nil {
		return errors.New("token invalid, claims exp empty")
	}

	// look the containts of claims
	expires_at := int(claims["exp"].(float64))

	// convert expires_at to time.Time
	expires_at_time := time.Unix(int64(expires_at), 0)

	// cek if token expired
	if time.Now().Unix() > expires_at_time.Unix() {
		return errors.New("token expired")
	}

	return nil
}
