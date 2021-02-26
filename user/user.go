package user

import (
	"fmt"
)

type User struct {
	username string
	password string
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

func (u *User) String() string {
	return fmt.Sprintf("username: %s, password: %s", u.Username(), u.Password())
}
