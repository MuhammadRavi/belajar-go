package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(file)

	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func addToFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)

	return nil
}

func main() {
	// createNewFile("sample.log", "this is sample log")

	addToFile("sample.log", "\nMuhammad Ravi Mega Arasy")
	// result, _ := readFile("sample.log")
	// fmt.Println(result)
}
