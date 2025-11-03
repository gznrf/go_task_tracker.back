package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func GenerateToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})

	tokenKey := []byte(viper.GetString("token.key"))
	return token.SignedString(tokenKey)
}

func ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(viper.GetString("token.key")), nil
	})
	if err != nil {
		return 0, err
	}

	tclaims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return tclaims.UserId, nil
}

func SetTokenToCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}
