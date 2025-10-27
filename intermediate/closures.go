package intermediate

import "fmt"

func ClosureDemo() {
	increment := adder()
	// Initial state of num is 0
	// Each subsequent call of 'increment' increases num by 1
	// The inner function may access the outer functions values (This is known as a closure)
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())


	// Creating a new instance 'increment2' will also create a new instance of num set to 0
	increment2 := adder()
	fmt.Println(increment2())

	// We can create closuers with anonymous functions as well
	subtracter := func() func(int) int {
		countdown := 100
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracter(23))
	fmt.Println(subtracter(22))
	fmt.Println(subtracter(10))
}

func adder() func() int {
	num := 0
	fmt.Println(num)
	return func() int {
		num++
		fmt.Println("added 1 to num")
		return num
	}
}