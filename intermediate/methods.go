package intermediate

import "fmt"

type Shape struct {
	Rectangle // Anonymous structure embedding
}

type Rectangle struct {
	length float64
	width float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func (r *Rectangle) scale(factor float64) {
	r.length *= factor
	r.width *= factor
}


// We can also define methods on user defined datatypesb 
type CustomType int
func (c CustomType) isPositive() bool {
	return c > 0
}

func MethodDemo() {
	r := Rectangle{
		length: 10,
		width: 9,
	}
	fmt.Println("Area of the rectangle is: ", r.area())
	r.scale(2)
	fmt.Println("Area of the rectangle is: ", r.area())

	var num CustomType = 10
	var num2 CustomType = -19
	fmt.Println(num.isPositive())
	fmt.Println(num2.isPositive())

	s := Shape{
		Rectangle: Rectangle{
			length: 12,
			width: 9,	
		},
	}

	fmt.Println(s.area()) // We can call the method associated with the rectangle struct since we embedded it in shape
	fmt.Println(s.Rectangle.area()) // This also works the same
}