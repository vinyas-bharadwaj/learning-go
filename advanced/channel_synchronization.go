package advanced

import (
	"fmt"
	"time"
)

func ChannelSyncDemo() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		// Passing an empty struct to this channel
		done <- struct{}{}
	}()	

	// This mechanism ensures that the goroutine finishes execution before the main thread exits
	// The below operation blocks the main thread and waits until a value is passed into the channel
	// A value is only passed into the channel once the goroutine finishes execution
	<- done

	// Below is an example of synchronizing multiple goroutines
	numGoroutines := 5
	ch := make(chan int, numGoroutines)
	for i := range numGoroutines {
		go func(id int) {
			fmt.Printf("Goroutine %d is running...\n", id)
			time.Sleep(time.Second)
			ch <- id
		}(i)
	}

	for range numGoroutines {
		<- ch // Wait for all goroutines to finish
	}

	// Below is an example of synchronizing data exchange
	data := make(chan string)
	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("Hello %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		// We close this channel after it finishes execution
		// This prevents the error with the for loop 
		close(data)
	}()

	// Note that when we range over a channel, it acts the same as receiving values from the channel
	// Also important to note that this loop will keep running as long as values are added to the channel
	// This may also lead to an error where the loop tries to receive values from an empty channel
	// Therefore we use the 'close' function inside the goroutine to close the channel
	for val := range data {
		fmt.Println("Received:", val)
	}

	fmt.Println("Main thread exited")
}