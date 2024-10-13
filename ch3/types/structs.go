package my_types

import "fmt"

// Structures
// - replaces classes
// - strongly types props
// - have a constructor
// - can have methods

type User struct {
	id   int
	name string
}

func (u User) String() string {
	return fmt.Sprintf("User: %d, %s", u.id, u.name)
}

func TestStructs() {
	var u1 User
	u1 = User{id: 1, name: "modeste"}
	u2 := User{id: 2, name: "mukendi"}

	fmt.Println(u1.String(), u2.String())
}
