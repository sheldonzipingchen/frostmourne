package user

import "testing"

func TestSetUsername(t *testing.T) {
	u := &User{}
	u.SetUsername("Sheldon Chen")
	if u.Username() != "Sheldon Chen" {
		t.Errorf("username is not equal to Sheldon Chen: %v", u)
	}
}

func TestSetPassword(t *testing.T) {
	u := &User{}
	u.SetPassword("123456")
	if u.Password() == "" {
		t.Errorf("password is nil: %v", u)
	}
}

func TestString(t *testing.T) {
	u := &User{}
	u.SetUsername("Sheldon Chen").
		SetPassword("123456")
	_, err := u.String()
	if err != nil {
		t.Errorf("user object marshal to a json failed: %v", err)
	}
}
