package main

import "fmt"

type Address struct {
	Country, Province, City string
}

func main() {
	// Contoh Tanpa Pointer
	fmt.Println("\n============= Tanpa Pointer ================")
	address1 := Address{
		Country:  "Indonesia",
		Province: "DKI Jakarta",
		City:     "Jakarta Utara",
	}
	address2 := address1
	address2.City = "Jakarta Selatan"
	fmt.Println(address1)
	fmt.Println(address2)

	// Contoh Dengan Pointer
	fmt.Println("\n============= Dengan Pointer ================")
	var Address1 Address = Address{
		Country:  "Indonesia",
		Province: "DKI Jakarta",
		City:     "Jakarta Utara",
	}
	var Address2 *Address = &Address1

	Address2.City = "Jakarta Selatan"
	fmt.Println(Address1)
	fmt.Println(Address2)
}
