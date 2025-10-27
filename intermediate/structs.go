package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
	address Address // Example of a nested structure
	Phone // Anonymous struct field
}

type Address struct {
	city string
	country string
}

type Phone struct {
	number string
	pincode string
}

// We can also have methods associated with structs
// Structs and methods must be defined at package level 
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Even structs as passed by value
// Therefore, any method that modifies a struct attribute must take a pointer to the struct
func (p *Person) passTime() {
	p.age++
}

func StructDemo() {
	p := Person{
		firstName: "Emilia",
		lastName: "Clarke",
		age: 18,
		address: Address{
			city: "London",
			country: "U.K",
		},
		Phone: Phone{
			number: "1234567890",
			pincode: "1234567890",
		},
	}

	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	fmt.Println(p.age)
	fmt.Println(p.address.city, p.address.country)
	// We can directly access anonymous struct attributes without passing through the nested struct
	fmt.Println(p.number, p.pincode) 

	// Anonymous structs
	user := struct {
		username string
		password string
	} {
		username: "Emilia",
		password: "password123",
	}

	fmt.Printf("Username: %s Password: %s\n", user.username, user.password)
	fmt.Println(p.fullName())
	p.passTime()
	fmt.Println(p.age)
}

