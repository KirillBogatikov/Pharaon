package encryption

import "golang.org/x/crypto/bcrypt"

const (
	BCryptCost        = 10
	BCRYPT     Method = "bcrypt"
)

type BCrypt struct{}

func NewBCrypt() Encryptor {
	return &BCrypt{}
}

func (b *BCrypt) Encrypt(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), BCryptCost)
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func (b *BCrypt) Compare(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}

	return err == nil, err
}

func init() {
	implementations[BCRYPT] = NewBCrypt
}
