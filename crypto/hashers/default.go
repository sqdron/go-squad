package hashers

import (
	"crypto/rand"
	"crypto/subtle"
	"hash"
)

type defaultHasher struct {
	secret string
	Cost   int
	Hasher hash.Hash
}

func DefaultHash(h hash.Hash, cost int) hasher {
	return hasher{provider: &defaultHasher{Hasher: h, Cost: cost}}
}

func (h *defaultHasher) Hash(password string) string {
	salt := make([]byte, h.Cost)
	if _, err := rand.Read(salt); err != nil {
		panic("failed to generate random salt: " + err.Error())
	}
	h.Hasher.Reset()
	h.Hasher.Write(salt)
	h.Hasher.Write([]byte(password))
	return string(h.Hasher.Sum(salt))
}

func (h *defaultHasher) Check(password, hashed string) bool {
	saltLen := len(hashed) - h.Hasher.Size()
	salt, answer := hashed[:saltLen], hashed[saltLen:]
	h.Hasher.Reset()
	h.Hasher.Write([]byte(salt))
	h.Hasher.Write([]byte(password))
	var testHash []byte = h.Hasher.Sum(nil)
	return subtle.ConstantTimeCompare(testHash, []byte(answer)) == 1
}
