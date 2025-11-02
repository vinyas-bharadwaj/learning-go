package advanced

import (
	"errors"
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and execute in the background in a seperate thread
// They join the main thread once the function finishes execution
// Goroutines are non-blocking
// The go runtime handles goroutines
// It follows a many-to-many model where multiple goroutines are multiplexed to multiple kernel threads

func GoroutineDemo() {
	fmt.Println("Goroutine began!")
	go sayHello()

	// The below print statement is executed before the print in the 'sayHello' function
	// This is because goroutines are non blocking
	fmt.Println("Random text!")

	go printNumbers()
	go printNumbers()

	// err := go doSomething() // Note that this will not work
	// There is a clever work-around for this 

	go func() {
		err := doSomething()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}()

	// We do this so that the main thread doesn't exit before the goroutine
	time.Sleep(5 * time.Second)
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := range 10 {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func doSomething() error {
	// Simulate work
	time.Sleep(1 * time.Second)
	return errors.New("somethign went wrong")
}