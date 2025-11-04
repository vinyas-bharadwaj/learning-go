package advanced

import (
	"fmt"
	"sync"
	"time"
)

func process(id int, wg *sync.WaitGroup) {
	// We signal that the process finishes execution by calling 'wg.Done()' at the end of the function
	// 'wg.Done()' decrements the counter by 1
	defer wg.Done()

	fmt.Printf("Process %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Process %d finished\n", id)
}

func WaitGroupsDemo() {
	// We can define a waitgroup using the 'sync' package
	var wg sync.WaitGroup
	numProcesses := 3

	// Increments the counter by 'numProcesses'
	wg.Add(numProcesses)

	// It is also important to note that it is wrong practice to call 'wg.Add()' inside the goroutine itself	
	// This fails to work since the program already reaches 'wg.Wait()' before 'wg.Add()' is even called once

	for i := range numProcesses {
		go process(i, &wg)
	}

	// Blocks until the counter becomes 0
	wg.Wait()
	fmt.Println("All processes finished execution")
}