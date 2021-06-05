package auth

import "testing"

func Test(t *testing.T) {
	if Service.Name != "auth" {
		t.Fatal(Service.Name)
	}

	if Database.URL != "postgres://localhost:5432/auth?user=auth&password=auth" {
		t.Fatal(Database.URL)
	}

	if Database.MaxConnections != 5 {
		t.Fatal(Database.MaxConnections)
	}
}
