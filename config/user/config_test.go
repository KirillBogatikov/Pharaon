package user

import "testing"

func Test(t *testing.T) {
	if Service.Name != "user" {
		t.Fatal(Service.Name)
	}

	if Http == nil {
		t.Fatal("http config required")
	}

	if Database == nil {
		t.Fatal("database config required")
	}
}
