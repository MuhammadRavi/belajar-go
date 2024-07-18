package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	fromSlice := days[:]
	toSlice := make([]string, len(days), cap(days))

	copy(toSlice, fromSlice)
	fmt.Println("From SLice: ", fromSlice)
	fmt.Println("To SLice: ", toSlice)
}
