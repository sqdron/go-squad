package hashers

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
	Cost int
}

func Bcrypt(cost int) hasher {
	return hasher{provider: bcryptHasher{Cost: cost}}
}

func (b bcryptHasher) Hash(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), b.Cost)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (b bcryptHasher) Check(password, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}
