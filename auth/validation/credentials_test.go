package validation

import (
	"github.com/Projector-Solutions/Pharaon-tools/validation"
	"pharaon-auth/data"
	"testing"
)

func login(t *testing.T) {
	c := &data.Credentials{Login: "nA9_!@()~"}
	r := ValidateCredentials(c, false)
	if r.Login != validation.Valid {
		t.Fatal("expected valid ", r.Login)
	}

	c.Login = "na"
	r = ValidateCredentials(c, false)
	if r.Login != validation.Short {
		t.Fatal("expected short ", r.Login)
	}

	//40 characters
	c.Login = "6sVkA7ZpRhKJV7qWfDwadY2ZcCXaU3DGF454fE3Y"
	r = ValidateCredentials(c, false)
	if r.Login != validation.Long {
		t.Fatal("expected long ", r.Login)
	}

	c.Login = "nata_[]"
	r = ValidateCredentials(c, false)
	if r.Login != validation.Incorrect {
		t.Fatal("expected invalid ", r.Login)
	}
}

func TestCredentialsValidator_Validate(t *testing.T) {
	t.Run("Login Validation", login)
}
