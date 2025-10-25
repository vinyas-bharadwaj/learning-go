package basics

import "fmt"


func SliceDemo() {
	// var slice_name[]slice_type
	// A slice is essentially an array having variable size
	var slice []int

	// We can use the 'make' function to declare a slice as well
	slice = make([]int, 5)
	fmt.Println(slice)

	// We can also create a slice from an array
	arr := [5]int{5, 4, 3, 2, 1}
	arrSlice := arr[1:4] // starting index is inclusive and ending index is exclusive
	fmt.Println(arrSlice)

	// We use the 'append' function to add values to the slice
	// We can add multiple values at the same time with the append function
	arrSlice = append(arrSlice, 11, 12)
	fmt.Println(arrSlice)

	// We use the 'copy' function to copy a slice
	sliceCopy := make([]int, len(arrSlice))
	copy(sliceCopy, arrSlice)
	fmt.Println(sliceCopy)
	
	// Traversing a slice is the same as an array
	for i, v := range arrSlice {
		fmt.Printf("At index %d the value is %d\n", i, v)
	}
}