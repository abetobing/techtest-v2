package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func Encrypt(s string, salt string) string {
	var encryptedStr = []byte(s)
	encryptedStr = append(encryptedStr, salt...)
	var hash = sha512.New()
	_, err := hash.Write(encryptedStr)
	if err != nil {
		panic(err)
	}
	var hashedPasswordBytes = hash.Sum(nil)
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash
}

func Salt() string {
	var salt = make([]byte, 32)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	var base64EncodedSalt = base64.URLEncoding.EncodeToString(salt)
	return base64EncodedSalt
}

func CheckPassword(password string, passwordHash string, salt string) (bool, error) {
	passBytes := []byte(password)
	return passwordHash == Encrypt(string(passBytes), salt), nil
}
