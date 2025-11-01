package intermediate

import (
	"encoding/json"
	"fmt"
)

type PersonJSON struct {
	// Note that the struct field name and the json key name may be different
	Name string `json:"name"`
	// 'omitempty' makes it so that after marshalling, if age is not present, it is omitted rather than displaying it's default value
	Age int `json:"age,omitempty"` 
	Email string `json:"email"`
	// We can also have nested structures
	Address AddressJSON `json:"address"`
}

type AddressJSON struct {
	City string `json:"city"`
	State string `json:"state"`
}

func JSONDemo() {
	person := PersonJSON{
		Name: "Emilia",
		Email: "EmiliaClarke@gmail.com",
	}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
	newPerson := PersonJSON{
		Name: "Jon",
		Age: 18,
		Email: "JonSnow@gmail.com",
		Address: AddressJSON{
			City: "Winterfell",
			State: "North",
		},
	}

	newJsonData, err := json.Marshal(newPerson)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(newJsonData))

	// Un-marshalling (json -> struct)
	jsonString := `{"name":"terion", "age":30, "email":"terionlannister@gmail.com", "address": {"city":"casterly rock", "state": "west"}}`
	var personFromJSON PersonJSON
	// We need to pass the reference to the 'personFromJSON' variable since we are changing values in place
	err = json.Unmarshal([]byte(jsonString), &personFromJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(personFromJSON)

	// Handling unknown json structures
	// We do this by using a map as opposed to a struct
	jsonString = `{"Hello": "World", "Emilia": "clarke"}`

	// We use interfaces as a substitute for 'any' datatype
	var decoded map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &decoded)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(decoded)
}