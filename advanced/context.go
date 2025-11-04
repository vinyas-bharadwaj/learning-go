package advanced

import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	// ctx.Done() returns a channel and we are receiving a value from this channel using '<-'
	case <- ctx.Done():
		return "Timeout!"
	default:
		if num % 2 == 1 {
			return "Odd"
		} else {
			return "Even"
		}
	}
}

func ContextDemo() {
	todoContext := context.TODO()
	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, "name", "Emilia")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	ctx2 := context.WithValue(contextBkg, "city", "tokyo")
	fmt.Println(ctx2)
	fmt.Println(ctx2.Value("city"))

	ctx3 := context.Background()

	// We set a context with a timeout of 1 second
	// After the 1 second time duration, the context sends a cancellation signal which can be accessed by ctx.Done()	
	ctx3, cancel := context.WithTimeout(ctx3, time.Second)
	defer cancel()

	result := checkEvenOdd(ctx3, 10)
	fmt.Println(result)

	// The below result will be 'Timeout!' since we set the context timeout to be 1 second
	time.Sleep(3 * time.Second)
	result = checkEvenOdd(ctx3, 15)
	fmt.Println(result)


}