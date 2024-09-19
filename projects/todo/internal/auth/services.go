package auth

import "golang.org/x/crypto/bcrypt"

func UserJWT() {

}

func HashPassword(string) (password string, err error) {

	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		password = string(bytePassword)
	}
	return
}

func ValidPassword(password string, hashedPassword string) (valid bool) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	valid = err == nil
	return
}
