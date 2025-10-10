package utils

import (
	"crypto/sha1"
	"fmt"

	"github.com/spf13/viper"
)

func GeneratePasswordHash(password string) string {
	salt := viper.GetString("password.hash_salt")
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
