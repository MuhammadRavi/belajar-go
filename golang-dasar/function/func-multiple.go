package main

import "fmt"

func main() {
	firstName, lastName := getFullName()
	fmt.Println("Hallo " + firstName + " " + lastName)

	// Ignore one value using _
	firstNameOnly, _ := getFullName()
	fmt.Println("Hallo " + firstNameOnly)
}

func getFullName() (string, string) {
	return "muhammad", "ravi"
}
