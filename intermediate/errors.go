package intermediate

import (
	"fmt"
	"errors"
)

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Math Error: Square root of a negative number")
	}

	// Assume we are computing the square root
	return 1, nil
}

func ErrorsDemo() {
	_, err := sqrt(16) // Works fine and hence 'err' is nil
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = sqrt(-1) // Throws an error and hence 'err' is not nil
	if err != nil {
		fmt.Println(err)
		// return
	}

	err = process()
	if err != nil {
		fmt.Println(err)
		// return
	}

	err = readData()
	if err != nil {
		fmt.Println(err)
	}


}

// Note that go's default package has an error interface which implements a single method called 'Error'
// Any structure that has an 'Error' method implicitly implements this interface

type MyError struct {
	message string
}

// Now, the 'MyError' struct implements the error interface
func (m *MyError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func process() error {
	return &MyError{"Custom error is being displayed!"}
}


// We can use 'fmt.Errorf' to define custom errors with the '%w' formatting verb
func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("Config Error")
}