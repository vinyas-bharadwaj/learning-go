package advanced

import "fmt"

// We can make our code more robust and explicit by mentioning the type of channel

func ChannelDirectionsDemo() {
	// A normal channel is bidirectional by default
	ch := make(chan int)

	// We can use a bidirectional channel as either send only or receive only
	go producer(ch)
	consumer(ch)

}

// Send only channel
func producer(ch chan <- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

// Receive only channel
func consumer(ch <- chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

