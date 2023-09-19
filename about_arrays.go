package go_koans

import "fmt"

func aboutArrays() {
	fruits := [4]string{"apple", "orange", "mango"}
	// looks like we need to declare size of array first

	assert(fruits[0] == "apple") // indexes begin at 0
	assert(fruits[1] == "orange") // one is indeed the loneliest number
	assert(fruits[2] == "mango") // it takes two to ...tango?
	assert(fruits[3] == "") // there is no spoon, only an empty value
	//because of type string, must be empty

	assert(len(fruits) == 4) // the length is what the length is
	assert(cap(fruits) == 4) // it can hold no more

	assert(fruits == [4]string{"apple", "orange", "mango"}) // comparing arrays is not like comparing apples and oranges
	// == comparison of values

	tasty_fruits := fruits[1:3]                           // defining oneself as a variation of another
	assert(fmt.Sprintf("%T", tasty_fruits) == "[]string") //and get not a simple array as a result
	//fyi these %x are called verbs. 
	//%v is any variable type (non-prod), 
	//%s is for strings
	//%q is for quotes
	//%d for integer
	//%t for boolean value
	//%T for type
	//here for more https://pkg.go.dev/fmt

	assert(tasty_fruits[0] == "orange")                 // slices of arrays share some data
	assert(tasty_fruits[1] == "mango")                 // albeit slightly askewed

	assert(len(tasty_fruits) == 2) // its length is manifest
	assert(cap(tasty_fruits) == 3) // but its capacity is surprising!
	//why 3? slices update the capacity but what capacity is set?

	tasty_fruits[0] = "lemon" // are their shared roots truly identical?

	assert(fruits[0] == "apple") // has this element remained the same?
	assert(fruits[1] == "lemon") // how about the second?
	assert(fruits[2] == "mango") // surely one of these must have changed
	assert(fruits[3] == "") // but who can know these things
	//OHHHHHH wow okay so slices still ref original array

	veggies := [...]string{"carrot", "pea"}
	//ah interesting, don't need to declare size

	assert(len(veggies) == 2) // array literals need not repeat an obvious length
}
