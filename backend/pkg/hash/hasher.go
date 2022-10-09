package hash

type Hasher interface {
	HashSha256(s string) string
}

type hasher struct {
	salt string
}

func NewHasher(salt string) Hasher {
	return &hasher{
		salt: salt,
	}
}
