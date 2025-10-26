package basics

import "fmt"

func RangeDemo() {
	message := "Hello World"

	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("Index: %d Rune: %c\n", i, v)
		// Characters are known as runes in go
	}
}