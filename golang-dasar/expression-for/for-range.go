package main

import "fmt"

func main() {
	namesArray := []string{"Muhammad", "Ravi", "Mega", "Arasy"}
	namesMap := map[string]string{"first": "Muhammad", "middle": "ravi", "last": "mega arasy"}

	fmt.Println("========= For Manual ==========")
	for i := 0; i < len(namesArray); i++ {
		fmt.Println("Name:  ", namesArray[i])
	}

	//1. index, name bisa kita ganti dengan apapun
	//   ketika data collection:
	// 		1: array --> maka balikannya terdiri atas index & value
	// 		2: map --> maka balikannya terdiri atas key & value
	// 2. ketika index tidak dipakai maka kita bisa menggantinya dengan "_"
	fmt.Println("\n========== For Range =============")
	fmt.Println("========= Data Collection Array =========")
	fmt.Println("=============== Dengan Index ==========")
	for index, name := range namesArray {
		fmt.Println("index ", index, " name ", name)
	}

	fmt.Println("\n =============== Tanpa Dengan Index ==========")
	for _, name := range namesArray {
		fmt.Println(" name ", name)
	}

	fmt.Println("\n========= Data Collection Map =========")
	for key, name := range namesMap {
		fmt.Println("key ", key, " name ", name)
	}
}
