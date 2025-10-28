package intermediate

import "fmt"

type Human struct {
	name string
	age int
}

type Employee struct {
	Human // Struct embedding
	empId string
	salary float64
}

func (h Human) introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old\n", h.name, h.age)
}

// This is an example of method overriding
func (e Employee) introduce() {
	fmt.Printf("Hello, my name os %s and I am %d years old and my ID is %s\n", e.name, e.age, e.empId)
}

func EmbeddingDemo() {
	emp := Employee{
		Human: Human{
			name: "Jon snow",
			age: 16,
		},
		empId: "1",
		salary: 45000.00,
	}

	// Since the struct Human is embedded, we can access it's properties directly from Employee
	fmt.Println(emp.name) 
	fmt.Println(emp.age)

	// We can also call any methods associated with Human directly from an instance of Employee since Human is embedded
	// Since we overwrote the 'introduce' method, the one associated with Employee is called
	emp.introduce()
}