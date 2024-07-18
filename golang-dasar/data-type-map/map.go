package main

import "fmt"

func main() {
	// Harcode
	fmt.Println()
	persons := map[string]string{
		"name":    "Muhammmad Ravi Mega Arasy",
		"address": "Jl. Sindang Rusun Koja Pari 1 LT.1 No.5",
	}

	fmt.Println("Data Map: ", persons)
	fmt.Println("Data Map Name: ", persons["name"])
	fmt.Println("Data Map Address: ", persons["address"])

	// Dynamic
	fmt.Println("\n======== Dynamic =======")
	persons2 := map[string]string{}

	persons2["name"] = "Muhammad Faruq"
	fmt.Println(persons2)
}
