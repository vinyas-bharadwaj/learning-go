package advanced

import (
	"fmt"
	"sync"
	"time"
)


func DeadlockDemo() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	// The below two goroutines will cause a deadlock
	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mutex 1")
		time.Sleep(time.Second)
		// Waits for 'mu2' to be unlocked before locking it
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mutex 2")
		// 'mu1' is never unlocked since we perpetually wait for 'mu2' to be unlocked before moving forward
		mu1.Unlock()
		mu2.Unlock()
	}()

	go func() {
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mutex 2")
		time.Sleep(time.Second)
		// Waits for 'mu1' to be unlocked before locking it
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mutex 1")
		// 'mu2' is never unlocked since we perpetually wait for 'mu1' to be unlocked before moving forward
		mu2.Unlock()
		mu1.Unlock()
	}()
	// The above two goroutines perpetually wait for eachother to unlock their resources before moving forward
	// This leads to them never being able to finish execution and they wait infinitely
	// This causes the deadlock

	// Wait for both goroutines to finish execution
	wg.Wait()
}