package advanced

import (
	"fmt"
	"time"
)

// An unbuffered channel cannot hold or retain values
// As soon as a value is added to an unbuffered channel, it also needs to be received

// Buffered channels can hold or retain values
// They do not need an immediate receiver

func BufferedChannelsDemo() {

	// A buffered channel with a capacity of 2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2


	// Buffered channels only block on send if the buffer is full
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(<- ch)
	}()
	fmt.Println("Blocking starts")
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
	// Buffered channels only block on receive if the buffer is empty
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 4
	}()
	
	fmt.Println("2nd Blocking starts")
	fmt.Println(<-ch)
}