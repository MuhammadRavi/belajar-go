package main

import "fmt"

type Filter func(string) string

// Contoh Penggunaan function sebagai parameter
// tanpa menggunakan function type declaration
func sayHelloWithParameter(name string, filter func(string) string) {
	fmt.Println("Hello ", filter(name))
}

// Contoh Penggunaan function sebagai parameter
// menggunakan function type declaration
func sayHelloWithParameterWithTypeDeclaration(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func main() {
	filter := spamFilter
	sayHelloWithParameter("Anjing", filter)

	sayHelloWithParameter("Muhamad Ravi Mega Arasy", spamFilter)
}
