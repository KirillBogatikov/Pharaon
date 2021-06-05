package encryption

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

var (
	salt  = "fp8fJfGtxacfTqzYu6e8YvhVXdcgAyvR"
	index = 3
)

const (
	SHA256 Method = "sha-256"
)

type Sha256 struct{}

func NewSha256() Encryptor {
	return &Sha256{}
}

func (s *Sha256) Encrypt(password string) (string, error) {
	saltText := strings.Join([]string{password[:index], salt, password[index:]}, "")
	sha := sha256.New()
	hashBytes := sha.Sum([]byte(saltText))
	return hex.EncodeToString(hashBytes), nil
}

func (s *Sha256) Compare(hash, password string) (bool, error) {
	actualHash, err := s.Encrypt(password)
	if err != nil {
		return false, nil
	}

	return hash == actualHash, nil
}

func init() {
	implementations[SHA256] = NewSha256
}
