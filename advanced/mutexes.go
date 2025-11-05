package advanced

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	value int
}

func (c *Counter) increment() {
	// We lock the mutex when we enter into the critical section
	c.mu.Lock()
	defer c.mu.Unlock() // We unlock the mutex at the end of our goroutine

	// Critical section
	c.value++
}

func (c *Counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func MutexDemo() {
	// Since we will spawn multiple goroutines, we need wait groups to synchronize and wait for them to finish
	var wg sync.WaitGroup
	numWorkers := 5
	counter := &Counter{}

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
			fmt.Println("Counter value:", counter.getValue())
		}()
	}


	// Wait for all the goroutines to finish execution
	wg.Wait()
	fmt.Println("Final value:", counter.value)

}