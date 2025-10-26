package basics

import "fmt"

func DeferDemo() {
	// 'defer' is generally used for resource cleanup
	// After getting a response from an API request, we generally call defer for cleanup
	// 'defer response.Body.Close()' -> You will run into this pattern quite often 
	process()
}

// The 'defer' keywords always takes a function or a method as it's next value
// This deferred function always executes at the end of the function call
// If there are multiple differed statements, they are executed in LIFO order
func process() {
	defer fmt.Println("This is the first deferred statement being executed")
	defer fmt.Println("This is the second deferred statement being executed")
	defer fmt.Println("This is the third deferred statement being executed")
	fmt.Println("Normal statement being executed")
}