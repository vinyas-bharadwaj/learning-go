package basics

import "fmt"

func VariadicDemo() {
	// A variadic function takes a variable number of arguments
	fmt.Println(summation(1, 5, 2, 1))
	fmt.Println(summation(8, 7, 1))

	// We can give the values of a list as the input for the variadic function as well
	numbers := []int{7, 4, 1, 9}
	fmt.Println(summation(numbers...))
}

// This function may accept a variable number of integer arguments
func summation(nums ...int) int {
	total := 0

	// nums is treated as a list
	for _, num := range nums {
		total += num
	}

	return total
}