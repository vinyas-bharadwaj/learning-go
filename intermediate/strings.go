package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func StringDemo() {
	// String conversion
	num := 21
	numString := strconv.Itoa(num)
	fmt.Println(len(numString))

	// String splitting
	fruits := "apple, orange, banana"
	cars := "Ferrari-Ford-Lamborgini"
	fruitsList := strings.Split(fruits, ",")
	carsList := strings.Split(cars, "-")
	fmt.Println(fruitsList)
	fmt.Println(carsList)


	// Joining 
	countries := []string{"India", "Japan", "USA"}
	joined := strings.Join(countries, ",")
	fmt.Println(joined)

	// Checking whether substring exists in string
	fmt.Println(strings.Contains(joined, "India"))
	
	// Replacing instances of a substring in a string
	replaced := strings.Replace(joined, ",", " <---> ", 2)
	fmt.Println(replaced)

	// Stripping whitespaces
	word := "    Hello World!        "
	fmt.Println(word)
	fmt.Println(strings.TrimSpace(word))
	fmt.Println(strings.ToLower(word))
	fmt.Println(strings.ToUpper(word))

	// Repeat
	fmt.Println(strings.Repeat("Emilia ", 3))

	// Counting
	fmt.Println(strings.Count(word, "o"))

	// Prefix
	fmt.Println(strings.HasPrefix(strings.TrimSpace(word), "Hello"))

	// Regular expressions
	message := "Hello, today is a wonderful day!!! 1234567890-abcd 1 12 3-a 4-b"
	// Note that regular expressions must be used with back ticks ``
	re := regexp.MustCompile(`\d+-a`) // Match with one or more digits
	// -1 indicates all matches, putting a normal integer determines the number of matches
	slice := re.FindAllString(message, -1)
	fmt.Println(slice)

	// We can use the 'strings.Builder' struct for better memory usage
	var builder strings.Builder 
	builder.WriteString("Hello")
	builder.WriteRune(' ')
	builder.WriteString("World")

	// Convert builder to string
	res := builder.String()
	fmt.Println(res)


}