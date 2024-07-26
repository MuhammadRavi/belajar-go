package main

import (
	"errors"
	"fmt"
)

var (
	ValidationErrors = errors.New("Validation Error")
	NotFoundErrors   = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationErrors
	}

	if id != "ravi" {
		return NotFoundErrors
	}

	return nil
}

func main() {
	err := GetById("Eko")
	if err != nil {
		if errors.Is(err, ValidationErrors) {
			fmt.Println(ValidationErrors)
		} else if errors.Is(err, NotFoundErrors) {
			fmt.Println(NotFoundErrors)
		}
	}
}
