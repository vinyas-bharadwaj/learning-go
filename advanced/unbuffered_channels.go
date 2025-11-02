package advanced

import "fmt"

func UnbufferedChannelsDemo() {
	// An unbuffered channel is a channel without a fixed size
	// Below is an example of an unbuffered channel
	ch := make(chan int)

	// Whenever we add a value to an unbuffered channel, it immediately looks for an outlet
	// Outlet in this case is any receiver

	// The below goroutine adds values to the channel
	go func() {
		ch <- 1
	}()

	// Receiver from the 'main' goroutine accepts the channel value
	// receiver waits for all goroutines to finish
	// receiver basically blocks the execution of the 'main' thread until it receives a value
	receiver := <- ch
	fmt.Println(receiver)

	// Note that the below code will not work
	// ch <- 1
	// go func() {
	// 	temp := <- ch
	// 	fmt.Println(temp)
	// }()

}