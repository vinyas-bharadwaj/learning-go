package intermediate

import (
	"bufio"
	"fmt"
	"os"
)



func FileDemo() {

	// Creating a file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error while creating a file!")
		return
	}
	defer file.Close() // Closes only once the function execution ends

	// Write data into a file
	data := []byte("Hello World!\nToday is a wonderful day!")
	_, err = file.Write(data) // 'Write' method takes a byte slice as an argument
	if err != nil {
		fmt.Println("Error writing into the file!")
		return
	}

	fmt.Println("Data has been written into output.txt successfully!")

	file, err = os.Create("WritingString.txt")
	if err != nil {
		fmt.Println("Error while creating a file!")
		return
	}
	defer file.Close()

	// Write string data into the file
	_, err = file.WriteString("You know nothing jon snow")
	if err != nil {
		fmt.Println("Error writing into the file!")
		return
	}

	fmt.Println("Data has been written into WritingString.txt successfully!")

	// Opening and reading a file
	file, err = os.Open("output.txt")
	if err != nil {
		fmt.Println("Error while opening the file!")
		return
	}
	defer func() {
		fmt.Println("Closing Open file")
		file.Close()
	}()
	
	fmt.Println("File opened successfully!")

	// Reading contents of the file
	data = make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error reading contents of the file")
		return
	}
	fmt.Println("Data read from output.txt is:", string(data))

	// If we want to read the contents of all the lines, we can use the 'bufio' package
	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file!")
		return
	}
}