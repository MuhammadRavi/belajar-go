package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Ravi Mega Arasy", "Ravi"))
	fmt.Println(strings.Split("Muhammad Ravi Mega Arasy", ""))
	fmt.Println(strings.ToLower("Muhammad Ravi Mega Arasy"))
	fmt.Println(strings.ToUpper("Muhammad Ravi Mega Arasy"))
	fmt.Println(strings.Trim("     Muhammad Ravi Mega Arasy    ", " "))
	fmt.Println(strings.ReplaceAll("Eko Eko Eko Eko", "Eko", "Ravi"))
}
