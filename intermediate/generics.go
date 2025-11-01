package intermediate

import "fmt"

// We define generics by using square brackets
// The variable 'T' can be of any type
func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // This will be set to the default/zero value of that particular datatype
		return zero, false
	}

	element := s.elements[len(s.elements) - 1]
	s.elements = s.elements[:len(s.elements) - 1]
	return element, true
}

func (s *Stack[T]) display() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty!")
		return
	}
	fmt.Println("Stack elements: ", s.elements)
}

func GenericsDemo() {
	// We can now see that the swap function can be used for any datatype	
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	name1, name2 := "Emilia", "Jon"
	name1, name2 = swap(name1, name2)
	fmt.Println(name1, name2)

	stk := Stack[int]{} // We created an integer stack
	stk.push(11)
	stk.push(33)
	stk.push(-4)

	stk.display()

	ele, _ := stk.pop()
	ele2, _ := stk.pop()
	ele3, _ := stk.pop()
	ele4, ok := stk.pop()

	fmt.Println(ele, ok)
	fmt.Println(ele2, ok)
	fmt.Println(ele3, ok)
	fmt.Println(ele4, ok)

	// We can see that our stack data structure works for strings as well
	// This stack will work for basically any datatype, including custom datatypes
	stringStk := Stack[string]{}
	stringStk.push("Emilia")
	stringStk.push("Jon")
	stringStk.push("Terion")
	
	stringStk.display()

	name, ok := stringStk.pop()
	fmt.Println(name, ok)
}