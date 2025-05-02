package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	// main.go
	fmt.Println("*** Example Proxy ***")

	mainDB := UsersDB{}

	user1 := User{ID: 1}
	user2 := User{ID: 2}
	user3 := User{ID: 3}

	mainDB.Add(user1).Add(user2).Add(user3)

	proxy := UserFinderProxy{
		MainDB:   mainDB,
		Stack:    UsersStack{},
		Capacity: 2,
	}

	proxy.Find(1)
	proxy.Find(2)
	proxy.Find(3)
	proxy.Find(2)
	proxy.Find(1)

	fmt.Print("*** End of Proxy ***\n\n\n")

}
