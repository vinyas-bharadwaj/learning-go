package intermediate

import "fmt"

// Any structure that has all the methods defined by the interface implicitly implements it
type Geometry interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// This is an extra method not present in the interface
// however, the Circle struct still implements the Geometry interface
// Extra methods are allowed as long as the methods specified in the interface are implemented
func (c Circle) diameter() float64 {
	return 2 * c.radius
}

func (s Square) area() float64 {
	return s.length * s.length
}

func (s Square) perimeter() float64 {
	return 4 * s.length
}

// Since both Square and Circle structs have an area and a perimeter method, they implicitly implement the Geometry interface
func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("Area is", g.area())
	fmt.Println("Perimeter is", g.perimeter())
}

func InterfaceDemo() {
	c := Circle{
		radius: 10,
	}

	s := Square{
		length: 4,
	}

	measure(c)
	measure(s)

	doSomething("Emilia Clarke")
	doSomething(33)
}

// Interfaces are also used in functions which may take any type of argument
func doSomething(something ...interface{}) {
	for _, v := range something {
		fmt.Println(v)
	}
}