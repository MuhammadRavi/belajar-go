package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	// Replace Value
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println("All Days Replace: ", days)

	// Append Slice
	// Karena kita mengunakan daySlice1 dengan capacity 2 maka ketika kita append
	// dia akan membuat array baru. jadi tidak akan mempengaruhi array aslinya.
	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println("daySlice2: ", daySlice2)
	fmt.Println("Days: ", days)
}
