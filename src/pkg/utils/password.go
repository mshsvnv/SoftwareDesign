package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pass []byte) string {

	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)

	if err != nil {
		// fmt.Errorf("Failed to generate password: %v", err)
		return ""
	}

	return string(hashed)
}
