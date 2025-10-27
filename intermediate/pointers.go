package intermediate

import "fmt"

func PointerDemo() {
	var ptr *int
	var num int = 10
	ptr = &num

	fmt.Println(num)
	fmt.Println(ptr) // ptr holds the memory address of num
	fmt.Println(*ptr) // Dereferencing a pointer

	// Important to note that go does not support pointer arithmetic like c and c++
	// Pointer operations are restricted to referencing and dereferencing
	modifyValue(ptr) // Modifies value of the original variable
	fmt.Println(num)


	// Non-primitives are passed by reference
	// Primities are passed by value
	arr := []int{1, 2, 3}
	modifyArr(arr)
	fmt.Println(arr)
}

func modifyValue(ptr *int) {
	*ptr++
}

func modifyArr(arr []int) {
	arr[0] = 10
}