package user

import "encoding/json"

type User struct {
	username string `json:"username"`
	password string `json:"password"`
}

func (u *User) Username() string {
	return u.username
}

func (u *User) SetUsername(username string) *User {
	u.username = username
	return u
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) *User {
	u.password = password
	return u
}

func (u *User) String() (string, error) {
	out, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(out), err
}
