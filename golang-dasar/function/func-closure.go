package main

import "fmt"

func main() {
	// Note: function clouse sebenarnya baik tapi alangkah baiknya kita
	// menggunakan dengan bijak, karena jika code sudah terlalu banyak
	// akan membingungkan kita dalam men-debug code'nya.

	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println("Value:", counter)

}
