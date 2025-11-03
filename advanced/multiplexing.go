package advanced

import (
	"fmt"
	"time"
)

func MultiplexingDemo() {
	// Trying to receive from an empty channel when no goroutine is running leads to a deadlock (endless waiting)
	// We can prevent this by using 'select'
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	time.Sleep(2 * time.Second)

	// The below code runs gracefully without deadlock
	// This is because we have a default case which runs if neither of the channels have values to send
	// We loop over twice if we want to see the output from each of the two channels
	for range 2 {
		select {
		case msg := <-ch1:
			fmt.Println("Message from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Message from ch2:", msg)
		default:
			fmt.Println("No channels ready")
		}
	}

	// We can also implement a timeout mechanism to automatically timeout after a set interval if a message isn't received
	ch := make(chan int)

	// This goroutine takes 3 seconds to send a value to the channel
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	// We timeout after 2 seconds of waiting for the channel to receive a value
	select {
	case msg := <-ch:
		fmt.Println("Message:", msg)
	case <- time.After(2 * time.Second):
		fmt.Println("Timeout! exceeded 2 seconds")
	}

	// We can also implement proper error handling mechanisms with channels
	demo := make(chan int)

	go func() {
		time.Sleep(time.Second)
		demo <- 1
		close(demo)
	}()

	for range 2 {
		select {
		case msg, ok := <- demo:
			if !ok {
				fmt.Println("The channel has been closed!")
				// Cleanup code
				return
			}
			fmt.Println("Message:", msg)
		}
	}
}