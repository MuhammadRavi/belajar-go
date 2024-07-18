package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ravi Customer

	// Manual Set Struct
	ravi.Name = "Muhammad Ravi Mega Arasy"
	ravi.Address = "Jl. Sindang Rusun Koja Pari 1 LT.1 No.5"
	ravi.Age = 27

	fmt.Println("All Object", ravi)
	fmt.Println("Name:", ravi.Name)
	fmt.Println("Address:", ravi.Address)
	fmt.Println("Age:", ravi.Age)

	// Set Struct With Literal
	fmt.Println("\n\nSet Struct With Literal")
	fmt.Println("\n====== Literal With Key =========")
	joko := Customer{
		Name:    "JOKO",
		Address: "Indonesia",
		Age:     27,
	}
	fmt.Println(joko)

	fmt.Println("\n====== Literal Without Key =========")
	/*
		Ketika kita menggunakan Set Struct Without Key Kita
		harus pastikan bahwa value'nya sesuai dengan field'nya
	*/
	budi := Customer{
		"BUDI",
		"Indonesia",
		27,
	}
	fmt.Println(budi)

}
