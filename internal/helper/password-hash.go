package helper

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) (string, error) {
	salt := 12
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePass(h string, p string) bool {
	pass, hash := []byte(p), []byte(h)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}