package advanced

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindowCounter struct {
	mu sync.Mutex
	count int
	limit int
	window time.Duration
	resetTime time.Time
}

func NewFixedWindowCounter(limit int, window time.Duration) *FixedWindowCounter {
	return &FixedWindowCounter{
		limit: limit,
		window: window,
	}
}

func (fwc *FixedWindowCounter) Allow() bool {
	fwc.mu.Lock()
	defer fwc.mu.Unlock()

	now := time.Now()
	if now.After(fwc.resetTime) {
		fwc.resetTime = now.Add(fwc.window)
		fwc.count = 0
	}

	// Increment counter if it is less than limit
	if fwc.count < fwc.limit {
		fwc.count++
		return true
	}

	return false
}

func FixedWindowCounterDemo() {
	// Creating a fixed window counter with a limit of 5 requests
	// The window determins the time after which the count resets to 0
	fwc := NewFixedWindowCounter(3, 1 * time.Second)
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if fwc.Allow() {
				fmt.Println("Request allowed!")
			} else {
				fmt.Println("Request denied!")
			}
			time.Sleep(300 * time.Millisecond)
		}()
	}

	// Wait for all the goroutines to finish execution
	wg.Wait()
}