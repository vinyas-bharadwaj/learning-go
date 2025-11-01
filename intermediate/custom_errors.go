package intermediate

import (
	"fmt"
)

func CustomErrorDemo() {
	err := showError()
	if err != nil {
		fmt.Println(err)
	}
}


type CustomError struct {
	code int
	message string
}

// Now this struct implements the error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}

func showError() error {
	return &CustomError{
		code: 404,
		message: "Page not found!",
	}
}