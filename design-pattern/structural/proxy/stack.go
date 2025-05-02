package proxy

import "fmt"

type UsersStack []User

func (u *UsersStack) Find(ID int) (User, error) {
	for _, user := range *u {
		if user.ID == ID {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

func (u *UsersStack) Add(user User) *UsersStack {
	*u = append(*u, user)
	return u
}
