package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ravi := Man{Name: "Muhammad Ravi Mega Arasy"}
	ravi.Married()

	fmt.Println(ravi)
}
