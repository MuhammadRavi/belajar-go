package main

import "fmt"

type Address struct {
	Country, Province, City string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := &Address{
		Country:  "",
		Province: "DKI Jakarta",
		City:     "Jakarta Utara",
	}
	changeAddressToIndonesia(address)
	fmt.Println(address)
}
