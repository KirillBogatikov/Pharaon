package proxy

import "testing"

func Test(t *testing.T) {
	if Service.Name != "proxy" {
		t.Fatal(Service.Name)
	}

	if Service.Database != nil {
		t.Fatal(Service.Database)
	}

	if Env.MongoURL != "mongodb://proxy:BGZendgysf9m2jr6@mongodb:27017/proxy" {
		t.Fatal(Env.MongoURL)
	}
}
