package proxy

import "fmt"

type User struct {
	ID int
}

type UserFinder interface {
	Find(ID int) (User, error)
}

type UserFinderProxy struct {
	MainDB   UsersDB
	Stack    UsersStack
	Capacity int
}

func (u *UserFinderProxy) Find(ID int) (User, error) {
	user, err := u.Stack.Find(ID)
	if err == nil {
		fmt.Println("Found in stack: ", user)
		return user, nil
	}

	user, err = u.MainDB.Find(ID)
	if err != nil {
		return User{}, err
	}

	fmt.Println("Found in mainDB: ", user)
	u.AddToStack(user)
	return user, nil
}

func (u *UserFinderProxy) AddToStack(user User) error {
	fmt.Println("Adding to stack: ", user)
	if len(u.Stack) >= u.Capacity {
		u.Stack = append(u.Stack[1:], user)
	} else {
		u.Stack.Add(user)
	}
	return nil
}
