package main

import "fmt"

func main() {
	sayHelloTo("Muhammad Ravi", "Mega Arasy")
}

func sayHelloTo(
	firstName string,
	lastName string,
) {
	fmt.Println("Hello ", firstName, lastName)
}
