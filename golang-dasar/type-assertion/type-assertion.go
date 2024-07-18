package main

import "fmt"

// Function dengan return value interface kosong.
func random() interface{} {
	return "OK"
}

func main() {

	// Cara manual tanpa
	result := random()
	// convertionString := result.(string)
	// fmt.Println(convertionString)

	// Cara dengan switch
	switch result.(type) {
	case string:
		fmt.Println("String", result)
	case int:
		fmt.Println("int", result)
	default:
		fmt.Println("Unkwon")
	}
}
