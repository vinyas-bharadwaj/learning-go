package basics

import "fmt"

func RecoverDemo() {
	// The 'recover' function is used to recover from a panic

	// Valid function call
	positive(10)

	// Invalid function call
	positive(-10)
}

func positive(num int) {
	defer func() {
		// The 'recover' function returns 'nil' if a panic was never called in the function
		if r := recover(); r != nil {
			// 'r' will have a value specified by the 'panic' function 
			// r = "Input is less than zero"
			fmt.Println("Recovered: ", r)
		}
	}()

	fmt.Println("Function started")
	if num < 0 {
		panic("Input is less than zero")
	} 
}