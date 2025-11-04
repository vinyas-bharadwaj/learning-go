package advanced

import (
	"fmt"
	"time"
)

func TimersDemo() {

	timer := time.NewTimer(2 * time.Second)

	stopped := timer.Stop() // We use timer.Stop() if we want to early stop the timer
	if stopped {
		fmt.Println("Timer stopped early")
	}

	timer.Reset(time.Second) // Restarts a stopped timer
	<- timer.C // Blocking task (Blocks for 2 seconds)
	fmt.Println("Timer expired")

	timer.Reset(2 * time.Second)

	// We can use the timer to execute goroutines after a specified delay
	go func() {
		<- timer.C
		fmt.Println("goroutine executed after delay")
	}()

	time.Sleep(3 * time.Second)
}