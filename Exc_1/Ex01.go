package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Method with receiver type to print struct contents
func (p Person) Print() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	// Create an instance of Person
	person := Person{Name: "Alice", Age: 30}
	// Print the contents using the method
	person.Print()
}
