package user

import "encoding/json"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) String() (string, error) {
	out, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(out), err
}
