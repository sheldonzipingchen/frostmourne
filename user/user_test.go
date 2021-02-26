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
	s := u.String()
	if s == "" {
		t.Errorf("user object marshal to a json failed")
	}
}
