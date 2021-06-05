package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
)

func TestBCrypt_Encrypt(t *testing.T) {
	bc := NewBCrypt()
	hashed, err := bc.Encrypt(testPassword)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(testPassword))
	if err != nil {
		t.Fatal(err)
	}
}

func TestBCrypt_Compare(t *testing.T) {
	bc := NewBCrypt()
	ok, err := bc.Compare(testBCryptHash, testPassword)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("hash comparing incorrect")
	}
}

func TestBCrypt_CompareFailed(t *testing.T) {
	bc := NewBCrypt()
	incorrectHash := strings.ReplaceAll(testBCryptHash, "k", "j")
	ok, err := bc.Compare(incorrectHash, testPassword)
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		t.Fatal("hash comparing incorrect")
	}
}
