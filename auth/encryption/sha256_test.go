package encryption

import (
	"testing"
)

func TestSha256_Encrypt(t *testing.T) {
	sha := NewSha256()
	result, err := sha.Encrypt(testPassword)
	if err != nil {
		t.Fatal(err)
	}

	if result != testSha256Hash {
		t.Fatal("hash incorrect")
	}
}

func TestSha256_Compare(t *testing.T) {
	sha := NewSha256()
	ok, err := sha.Compare(testSha256Hash, testPassword)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("hash comparing incorrect")
	}
}

func TestSha256_CompareFailed(t *testing.T) {
	sha := NewSha256()

	ok, err := sha.Compare(testSha256HashIncorrect, testPassword)
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		t.Fatal("hash comparing incorrect")
	}
}
