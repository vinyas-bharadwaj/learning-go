package advanced

import (
	"context"
	"fmt"
)

func ContextDemo() {
	todoContext := context.TODO()
	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, "name", "Emilia")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	ctx2 := context.WithValue(contextBkg, "city", "tokyo")
	fmt.Println(ctx2)
	fmt.Println(ctx2.Value("city"))
}