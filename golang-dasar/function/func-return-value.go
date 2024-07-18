package main

import "fmt"

func main() {
	name := getHello("Muhammad Ravi", "Mega Arasy")
	fmt.Println(name)
}

func getHello(firstName string, lastName string) string {
	return ("Hello " + firstName + " " + lastName)
}
