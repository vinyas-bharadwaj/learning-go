package basics

import(
	"fmt"
	foo "net/http"
)

// Go includes only the required imports and libaries 
// Hence it reduces the final bundle size (This is known as tree shaking)
// we basically name the 'net/http' package as 'foo' which acts as the alias

func MakeRequest() {
	fmt.Println("Making an API request")
	resp, err := foo.Get("https://github.com")
	if err != nil {
		fmt.Println("Some error occurred while making the request")
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response status: ", resp.Status)
}