package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func main() {
	siswa := Person{
		Name: "Muhammad Ravi Mega Arasy",
	}
	sayHello(siswa)
}
