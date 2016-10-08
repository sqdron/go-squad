package hashers

type hasher struct {
	provider interface{}
}

type IHasher interface {
	Hash(password string) string
	Check(password, hashed string) bool
}

func (h *hasher) Hash(password string) string {
	return h.provider.(IHasher).Hash(password)
}

func (h *hasher) Check(password, hashed string) bool {
	return h.provider.(IHasher).Check(password, hashed)
}
