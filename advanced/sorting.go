package advanced

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a ByAge) Swap(i, j int)  {
	a[i], a[j] = a[j], a[i]
}

// We can define a more custom comparator to encapsulate basically any value in the struct
type By func(p1, p2 *Person) bool

type PersonSorter struct {
	people []Person
	by func(p1, p2 *Person) bool
}

func (s *PersonSorter) Len() int {
	return len(s.people)
}

func (s *PersonSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *PersonSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

// We define a sort method
func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by: by,
	}
	sort.Sort(ps)
}

func SortingDemo() {
	strings := []string{"Emilia", "Jon", "Tyrion", "Twyin", "Sansa"}
	sort.Strings(strings)
	fmt.Println("Sorted:", strings)

	people := []Person{
		{"Emilia", 20},
		{"Jon", 22},
		{"Tyrion", 23},
		{"Sansa", 16},
	}

	// We can use the 'sort' package on a custom datatype
	// Below function sorts by age since we defined a custom 'Less' function for the 'ByAge' type
	sort.Sort(ByAge(people))
	fmt.Println("Sorted: ", people)

	// We can make use of our custom 'By' type to sort by any value
	age := func(p1, p2 *Person) bool {
		return p1.age < p2.age
	}

	By(age).Sort(people)
	fmt.Println("Sorted by age:", people)

	name := func(p1, p2 *Person) bool {
		return p1.name < p2.name
	}
	By(name).Sort(people)
	fmt.Println("Sorted by name:", people)
}