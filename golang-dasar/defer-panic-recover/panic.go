package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	message := recover()
	fmt.Println("Terjadi Error: ", message)
}

func runApp(isError bool) {
	defer endApp()

	if isError {
		panic("Terdapat Masalah Pada Function")
	}
}

func main() {
	fmt.Println("=============== Run Application Error ===============")
	runApp(true)
	fmt.Println("Muhammad Ravi Mega Arasy")

}
