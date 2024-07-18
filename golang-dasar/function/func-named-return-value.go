package main

import "fmt"

func getCompleteName() (firstName, midleName, lastName string) {
	firstName = "Muhammad"

	return firstName, midleName, lastName
}

func main() {
	firstName, midleName, lastName := getCompleteName()

	fmt.Println("Hallo " + firstName + " " + midleName + " " + lastName)
}
