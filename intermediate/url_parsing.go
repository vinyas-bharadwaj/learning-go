package intermediate

import (
	"fmt"
	"net/url"
)

func URLParsingDemo() {
	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing the URL:", err)
		return
	}
	fmt.Println(parsedURL.Scheme)
	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Port())
	fmt.Println(parsedURL.Path)
	fmt.Println(parsedURL.RawQuery)
	fmt.Println(parsedURL.Fragment)

	rawURL2 := "https://example.com/path?name=emilia&age=18"
	parsedURL2, err := url.Parse(rawURL2)

	if err != nil {
		fmt.Println("Error parsing the URL:", err)
		return
	}

	queryParams := parsedURL2.Query()
	fmt.Println(queryParams.Get("name"))
	fmt.Println(queryParams.Get("age"))

	// Building a URL
	baseURL := &url.URL{
		Scheme: "https",
		Host: "example.com",
		Path: "/path",
	}
	query := baseURL.Query()
	query.Set("name", "Emilia")
	query.Set("Age", "18")
	baseURL.RawQuery = query.Encode()
	fmt.Println(baseURL.String())

	values := url.Values{}
	values.Add("name", "jon")
	values.Add("age", "22")
	values.Add("country", "north")

	// Encode
	encodedQuery := values.Encode()

	// Build URL
	baseURL2 := "https://example.com/search"
	fullURL := baseURL2 + "?" + encodedQuery
	fmt.Println("Full URL is:", fullURL)
}