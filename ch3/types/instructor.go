package my_types

type Instructor struct {
	// private props
	id    int
	score int
	// public (aka exported) props
	FirstName string
	LastName  string
}

func (i Instructor) String() string {
	return i.FirstName + " " + i.LastName
}

func (i Instructor) GetScore() int {
	return i.score
}

// factory to create a new Instructor
func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName}
}
