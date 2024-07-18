package main

import "fmt"

func main() {
	books := make(map[string]string)
	books["title"] = "Buku Go-Lang"
	books["author"] = "Muhammad Ravi"
	books["wrong"] = "Ups"

	fmt.Println("Data All Map: ", books)

	delete(books, "wrong")
	fmt.Println("Data After Delete: ", books)
}
