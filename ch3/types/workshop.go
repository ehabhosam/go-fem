package my_types

import "time"

// without embedding
// type Workshop struct {
// 	Course Course
// 	Date time.Time
// }

// with embedding
type Workshop struct {
	Course
	Date time.Time
}

// embeddingg is go's way to implement inheritance
// workshop will have everything that course has
// plus what we add to it.
