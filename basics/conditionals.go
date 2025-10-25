package basics

import "fmt"

func ConditionalsDemo() {
	// Conditionals basically work the same as any other language
	var temperature float32 = 25.01
	if temperature >= 25 {
		fmt.Println("Hot")
	} else {
		fmt.Println("Cold")
	}

	// Cool thing about conditionals in go is that you can perform assignment and conditional logic in the same line
	var age int
	if fmt.Scan(&age); age <= 20 {
		fmt.Println("Less than or equal to 20")
	} else {
		fmt.Println("Greater than 20")
	}

	// Switch statement is also very similar, except you do not need a break statement at the end of each case
	var choice int
	fmt.Print("Choice: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("1. Picked")
	case 2, 3:
		fmt.Println("2. or 3. Picked")
	default:
		fmt.Println("Default!!!")
	}

	// We can use the switch-case syntax to check the type of a variable as well
	CheckType(10)
	CheckType(10.1)
	CheckType("Emilia")
	CheckType('a')
}


func CheckType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It is an integer")
	case float64:
		fmt.Println("It is a floating point number")
	case string:
		fmt.Println("It is a string")
	default:
		fmt.Println("Unknown datatype")
	}
}
