package basics

import "fmt"

func MapDemo() {
	// var map_name map[key_type]value_type
	// map_name := make(map[key_type]value_type)

	counter := make(map[int]int)
	var nums []int = []int{7, 8, 5, 5, 1, 3, 8}

	for _, num := range nums {
		counter[num] += 1
	}

	fmt.Println(counter)

	// We can use the 'delete' function to delete a key value pair
	delete(counter, 5)
	fmt.Println(counter)

	// We can use the 'clear' function to completely remove all values from the map
	clear(counter)
	fmt.Println(counter)

	// Accessing an element form the map using a key returns 2 values
	// The first value is the value associated with the key itself
	// The second value is a boolean which represents whether the value is present in the map or not
	counter[7] = 8
	value, exists := counter[7]
	fmt.Println(value, exists)

	// We can directly declare a map without using the 'make' function
	mpp := map[int]int{}
	mpp[1] = 2
	mpp[2] = 3
	fmt.Println(mpp)

	for k, v := range mpp {
		fmt.Printf("key: %d value: %d\n", k, v)
	}

	// A declared map which isn't initialized is always initialized to nil
	var demoMap map[int]int
	if demoMap == nil {
		fmt.Println("The map is initialized to nil")
	} else {
		fmt.Println("The map is not nil")
	}
}