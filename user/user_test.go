package user

import "testing"

func TestString(t *testing.T) {
	u := &User{
		Username: "Sheldon Chen",
		Password: "123456",
	}
	_, err := u.String()
	if err != nil {
		t.Errorf("user object marshal to a json failed: %v", err)
	}
}
