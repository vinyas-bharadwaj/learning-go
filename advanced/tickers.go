package advanced

import (
	"fmt"
	"time"
)

func TickersDemo() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	count := 0

	// This loop runs infinitely and prints values after every second
	// We need to implement some breaking mechanism
	for tick := range ticker.C {
		fmt.Println("Tick at:", tick)
		count++
		if count == 5 {
			break
		}
	}

	i := 0
	for range ticker.C {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	// We can use tickers in combination with 'time.After()' in order to implement useful techniques
	stop := time.After(3 * time.Second)
	for {
		select {
		case tick := <- ticker.C:
			fmt.Println("Ticker ticked at:", tick)
		case <- stop:
			fmt.Println("Ticker stopped") 
			return
		}
	}
}