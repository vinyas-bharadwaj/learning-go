package advanced

import (
	"fmt"
	"os/exec"
	"time"
)


func ProcessDemo() {
	// Spawns a process which runs the following shell command
	// Overhead for creating processes is much larger than goroutines or threads
	cmd := exec.Command("echo", "Hello World")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))

	// Using sleep and wait commands
	cmd = exec.Command("sleep", "5")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// We can kill the process early after 2 seconds
	time.Sleep(2 * time.Second)
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// We need to wait and block other operations until the 5 second sleep gets over
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Process finished execution!")
}