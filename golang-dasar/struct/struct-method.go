package main

import "fmt"

type CustomerRegistration struct {
	Name, Address string
	Age           int
}

func (customer CustomerRegistration) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func (customer CustomerRegistration) modifyAge(newAge int) int {
	return customer.Age + newAge
}

func main() {
	ravi := CustomerRegistration{
		Name:    "Muhammad Ravi Mega Arasy",
		Address: "Jl. Sindang Rusun Koja",
		Age:     27,
	}

	fmt.Println("============ Function Not Return Value =============")
	ravi.sayHello()

	fmt.Println("============ Function Return Value =============")
	hasil := ravi.modifyAge(30)
	fmt.Println("Umur Terbaru:", hasil)

}
