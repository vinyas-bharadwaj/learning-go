package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func EnvDemo() {
	// Accessing environment variables through the 'os' package
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println(user, home)

	// Setting environment variables
	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return
	}

	for _, ele := range os.Environ() {
		kvpair := strings.SplitN(ele, "=", 2)
		fmt.Printf("Key: %s, Value: %s\n", kvpair[0], kvpair[1])
	}

	// Unsetting the environment variable
	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable", err)
	}
}