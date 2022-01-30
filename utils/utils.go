package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

const SALT = "secret-salt@#$%"

func HashPassword(password string) string {
	var passwordBytes = []byte(password)

	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, []byte(SALT)...)
	sha512Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	return hex.EncodeToString(hashedPasswordBytes)
}

func CheckPasswordHash(password, hash string) bool {
	return hash == HashPassword(password)
}
