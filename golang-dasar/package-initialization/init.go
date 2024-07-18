package main

import (
	"Go/package-initialization/database"
	_ "Go/package-initialization/internal"
	"fmt"
)

func main() {
	connection := database.GetDatabase()
	fmt.Println(connection)
}
