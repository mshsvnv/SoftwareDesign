package utils

import "golang.org/x/crypto/bcrypt"

func HashAndSalt(pass []byte) string {

	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	if err != nil {
		return ""
	}
	return string(hashed)
}
