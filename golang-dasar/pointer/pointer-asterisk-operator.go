package main

import "fmt"

type Address struct {
	Country, Province, City string
}

func main() {
	var Address1 Address = Address{
		Country:  "Indonesia",
		Province: "DKI Jakarta",
		City:     "Jakarta Utara",
	}
	var Address2 *Address = &Address1

	Address2.City = "Jakarta Selatan"
	fmt.Println(Address1)
	fmt.Println(Address2)

	fmt.Println("\n=========== Asterisk Operator ==========")
	// ===== Cara dibawah tidak akan memperngaruhi Address1'nya ====
	// Address2 = &Address{
	// 	Country:  "Indonesia",
	// 	Province: "DKI Jakarta",
	// 	City:     "Jakarta Barat",
	// }

	/*
		ketika kita ingin membuat datar baru tapi untuk poiner Address2
		dan kita ingin Address1'nya terdampak juga maka kita bisa menambahkan
		asterisk(*) pada address2. kenapa kita nambahkan asterisk, karena Address2
		tipe data'nya merupakan pointer.
	*/
	*Address2 = Address{
		Country:  "Indonesia",
		Province: "DKI Jakarta",
		City:     "Jakarta Barat",
	}
	fmt.Println(Address1)
	fmt.Println(Address2)
}
