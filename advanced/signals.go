package advanced

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func SignalsDemo() {
	// Signals are a way for us to notify goroutines about certain events
	sigs := make(chan os.Signal, 1)

	// We can get the process id of the current process
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	// Notify channel on interrupt or terminate signal
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go func() {
		sig := <- sigs
		switch sig {
		case syscall.SIGINT:
			fmt.Println("Interrupt signal received")
		case syscall.SIGTERM:
			fmt.Println("Termination signal received")
		case syscall.SIGUSR1:
			fmt.Println("User defined signal received")
			fmt.Println("executing a user defined function...")
		}
		fmt.Println("Graceful exit")
		os.Exit(0)
	}()

	// Simulate work
	fmt.Println("Working...")

	// Infinite wait
	// Only way to exit is by using ctrl + c (interrupt)
	for {
		time.Sleep(time.Second)
	}
	
}