package advanced

import (
	"sync"
	"fmt"
)

type Human struct {
	name string 
	age int
}

func SyncPoolDemo() {
	// A pool just acts as a queue data structure
	var pool = sync.Pool{
		// 'New' function creates an instance whenever we call 'pool.Get()' method on an empty pool
		// Without the 'New' function, calling 'pool.Get()' on an empty pool would throw an error
		New: func() interface{} {
			fmt.Println("Creating a new human")
			return &Human{}
		},
	}

	// Getting an object from the pool
	human1 := pool.Get().(*Human)
	fmt.Println("Got human1")
	human1.name = "Jon"
	human1.age = 18
	fmt.Printf("Human1 - Name: %s, Age: %d\n", human1.name, human1.age)

	// Puts the instance back to the pool
	pool.Put(human1)
	fmt.Println("Returned the person back to the pool")

	// 2nd instance retains the values since we put the 1st instance back to the pool
	human2 := pool.Get().(*Human)
	fmt.Println("Got human2:", human2)

	// 3rd instance will have null values since the 2nd instance wasn't returned
	human3 := pool.Get().(*Human)
	fmt.Println("Got human3:", human3)
	
}