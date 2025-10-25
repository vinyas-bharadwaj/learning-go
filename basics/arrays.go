package basics

import "fmt"

func ArrayDemo() {
	// The general syntax to declare an array is
	// var array_name[array_size] datatype
	var numbers[5] int
	for idx := range 5 {
		fmt.Printf("Enter number at index %d: ", idx)
		fmt.Scan(&numbers[idx])
	}

	fmt.Println(numbers)

	// We can also use the walrus operator to create an array
	arr := [5]int{3, 7, 1, 9, 8}
	for idx, val := range arr {
		fmt.Printf("Element at index %d is %d\n", idx, val)
	}

	// Every declared variable must be used in go unless we name it with a blank identifier '_'

	// Comparing arrays
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println("Array1 is equal to Array2: ", arr1 == arr2)
	// Note that the references weren't compared and the actual array values were

	// 2-D arrays
	var matrix[3][3] int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	// Creating a reference to an array
	var arrayReference *[3]int
	arrayReference = &arr1 // arrayReference is a pointer to arr1 
	// Any changes made to arrayReference will be reflected on the original arr1

	arrayReference[0] = 101 // This line also modifies the original array

	fmt.Println(arr1)

}