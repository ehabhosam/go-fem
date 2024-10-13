package my_types

// Interfaces
// - definition of methods
// - emulates polymorphism

type Printable interface {
	// any struct that has a String method
	// will be considered as a Printable
	// even if it has more methods
	String() string
}

// so this is how golang implements polymorphism
// you can pass any struct that has a String method
// so you don't build for a specific struct
// but for a specific functionality
