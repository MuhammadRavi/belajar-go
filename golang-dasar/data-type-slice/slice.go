package main

import "fmt"

func main() {
	names := [...]string{"Muhammad", "Ravi", "Mega Arasy", "Joko", "Budi", "Nugraha"}

	// Slice manual
	// var slice []string

	// Low: 4, High = 6
	slice1 := names[4:6]
	fmt.Println("Low = 4, High = 6 : ", slice1)

	// Low: 0, High = 3
	slice2 := names[:3]
	fmt.Println("Low = 0, High = 3  ", slice2)

	// Low: 1, High = index tertinggi
	slice3 := names[1:]
	fmt.Println("Low = 1, High = index tertinggi :  ", slice3)

	// Low: 0, High = index tertinggi
	slice4 := names[:]
	fmt.Println("Low = 0, High = index tertinggi : ", slice4)

}
