package advanced

import (
	"fmt"
	"time"
)

// Channels are a way for goroutines to communicate with eachother and synchronize

func ChannelDemo() {
	greeting := make(chan string)
	greetingString := "Hello"

	// We cannot use channels in the main function
	// A goroutine must be present to receive from a channel

	// Seperate goroutine
	go func() {
		// Blocking because it is continously trying to receive values from the channel
		greeting <- greetingString
		greeting <- "world"
		greeting <- "Today is a wonderful day"
	}()

	go func() {
		temp := <- greeting
		fmt.Println(temp)
	}()

	// receiver is part of the 'main' goroutine
	// Thus, the channel is communicating between two goroutines
	receiver := <- greeting
	fmt.Println(receiver)
	receiver = <- greeting
	fmt.Println(receiver)

	alphabets := make(chan string)
	go func() {
		for _, ele := range "abcde" {
			alphabets <- fmt.Sprintf("Alphabet: %c", ele)
		}
	}()

	for range 5 {
		fmt.Println(<- alphabets)
	}

	time.Sleep(1 * time.Second)
}