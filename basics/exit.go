package basics

import (
	"fmt"
	"os"
)

func ExitDemo() {
	// This deferred print is never executed
	defer fmt.Println("This is a deferred statement!")

	fmt.Println("Function starting")

	// Once exit is called, deferred functions are not called
	// The program flow immediately exits from the function
	// 0 is for success, non zero exit value is for error conditions
	os.Exit(1)
} 