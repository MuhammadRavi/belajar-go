package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// Contoh function sebagai value adalah
	// kita tidak langsung set getGoodBye() tapi
	// kita set seperti dibaah
	contoh := getGoodBye

	fmt.Println(contoh("Muhammad Ravi"))
}
