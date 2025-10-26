package basics

import "fmt"

func PanicDemo() {
	// 'panic' causes the function to immediately exit it's execution
	// All the deferred methods are called (if any) after the function called 'panic'

	// Valid case
	evaluate(10)

	// Invalid case
	evaluate(-1)
}

func evaluate(num int) {
	// defers are called after the panic
	defer fmt.Println("The function finished execution") 
	if num < 0 {
		panic("The number is less than 0")
	}
	fmt.Println("Evaluating the number: ", num)
}