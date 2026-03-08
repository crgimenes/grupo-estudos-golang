package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Label() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}

type Employee struct {
	Person
	Role string
}

func main() {
	e := Employee{Person: Person{Name: "Ana", Age: 30}, Role: "Engineer"}
	e.Birthday()
	fmt.Printf("%s role=%s\n", e.Label(), e.Role)
}
