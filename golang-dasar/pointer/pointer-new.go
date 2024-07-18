package main

import "fmt"

type Address struct {
	Country, Province, City string
}

func main() {
	// Cara manual membuat pointer
	// var alamat1 *Address = &Address{}

	// Cara membuat pointer dengan function
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1
	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
