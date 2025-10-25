package basics

import "fmt"

func LoopsDemo() {
	// Simplest version of a for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Iterating over a collection
	numbers := []int{3, 7, 3, 1}
	for index, value := range numbers {
		fmt.Printf("Index: %d Number: %d\n", index, value)
	}

	// Using for loop as a while loop
	var count int = 10
	for count > 0 {
		fmt.Println(count)
		count--
	}

	// Infinite loop
	var sum int = 10
	for {
		sum += 10
		fmt.Println(sum)
		if sum >= 50 {
			break
		}
	}

}