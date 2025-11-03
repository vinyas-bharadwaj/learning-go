package advanced

import (
	"fmt"
	"time"
)

func ClosingChannelsDemo() {
	in := make(chan int)
	out := make(chan int)

	go prod(in)
	go filter(in, out)

	for val := range out {
		fmt.Println(val)
	}
}

func filter(in <- chan int, out chan <- int) {
	for val := range in {
		if val % 2 == 0 {
			out <- val
		}
	}
	close(out)
}

func prod(ch chan <- int) {
	for i := range 10 {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}
	close(ch)
}