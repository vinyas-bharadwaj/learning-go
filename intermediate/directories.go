package intermediate

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func DirectoryDemo() {

	// Creates a directory with the specified name
	err := os.Mkdir("subdir", 0755)
	checkError(err)
	defer os.RemoveAll("subdir") // Removes all directories with specified name

	// Read files and folders inside a directory
	result, err := os.ReadDir("basics")
	checkError(err)
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir())
	}

	// Change the current working directory to the specified one
	err = os.Chdir("basics")
	checkError(err)

	result, err = os.ReadDir(".")
	checkError(err)
	for _, entry := range result {
		fmt.Println(entry.Name())
	}

	// Get the current working directory
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// Creating temporary files
	tempFile, err := os.CreateTemp("", "temporaryFile")
	checkError(err)
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())


	fmt.Println("Temporary file created:", tempFile.Name())

}