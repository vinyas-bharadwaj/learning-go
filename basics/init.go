package basics

import "fmt"

// The init function is always executed before the main function
func init() {
	fmt.Println("Initializing stuff...")
}

// init functions in the same package are executed sequentially
func init() {
	fmt.Println("Second initialization...")
}

func InitDemo() {
	// init is executed even if it isn't called (it is executed by default)
	fmt.Println("Hello world")
}