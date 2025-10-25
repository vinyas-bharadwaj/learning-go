package basics

import (
	"fmt"
)

// Note that we cannot use the walrus shorthand for a global variable
var middelename = "Queen of the dragons"

func VariableDemo() {
	var age int // General way to create a variable
	const PI = 3.1415 // A constant cannot be changed or reassigned
	// Important to note that constants must be assigned values that can be determined during compile time
	fmt.Println(PI)
	
	// Explicitly assigning the datatype is optional
	var name string = "Emilia"
	var name2 = "Clarke"

	// We can use the walrus operator ':=' to ask the go compiler dynamically understand the datatype
	// Note that this shorthand notation can only be used inside functions
	email := "emiliaclarke@gmail.com"

	age = 18

	fmt.Printf("Name: %s, Age: %d, Email: %s\n", name + " " + middelename + " " + name2, age, email)
}