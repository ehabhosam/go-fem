package my_types

import "fmt"

// an alias for the type
type integer = int
type sliceOfStrings = []string
type json = map[string]string

// create a new type
type inttt int
type stringgg string

// add a method to the new created type
func (i inttt) add(j inttt) inttt {
	return i + j
}

/*
same as:
	class inttt {
		add(j inttt) {...}
	}
*/

// example
type pound float64
type kilogram float64

func (p pound) toKilogram() kilogram {
	return kilogram(p * 0.453592)
}

func (k kilogram) toPound() pound {
	return pound(k / 0.453592)
}

func Test() {
	k := kilogram(5)
	fmt.Println("pound:", k.toPound())
	fmt.Println("kilogram:", k.toPound().toKilogram())
}
