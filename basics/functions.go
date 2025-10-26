package basics

import "fmt"

func FunctionDemo() {
	res := add(5, 6)
	fmt.Println(res)

	// We can create anonymous functions by just creating a function without a name
	// Since the function doesn't have a name, we cannot call it later
	// An anonymous function must be called instantly or it has no effect
	func() {
		fmt.Println("Anonymous function is being called!")
	}()

	// We can assign a function to a variable
	// Basically just behaves as a normal function
	sayHello := func() {
		fmt.Println("Hello World!")
	}
	sayHello()

	// Functions are first class objects in go
	// Functions can return functions and functions can be passed in as function arguments
	fmt.Println(applyOperation(5, 6, add))
	fmt.Println(applyOperation(5, 6, subtract))

	multiplyBy5 := multiplyByFactor(5)
	fmt.Println(multiplyBy5(10))

	multiplyBy7 := multiplyByFactor(7)
	fmt.Println(multiplyBy7(7))

	quo, rem := divide(10, 3)
	fmt.Println(quo, rem)

	fmt.Println(greet("Jon Snow"))
}

// Go provides really neat syntactic sugar where we can specify the type of similar arguments with a single type definition
// func func_name(arg1, arg2 string) -> Here both arg1 and arg2 are strings and the go compiler understands this
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// Function that takes a function argument
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Function that returns a function
func multiplyByFactor(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

// Go also allows us to create functions with multiple return values
// The main use case of this is error handling
// We can return an optional error value associated with any function which may throw an error
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// We can also use a named return to make the code simpler
// The below function returns the 'greeting' variable by default
func greet(name string) (greeting string) {
	greeting = "hello " + name
	return
}