package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data interface{}) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if data != "ravi" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := saveData("", nil)
	if err != nil {
		// ============ Cara dengan kondisi IF =========
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown Error:", err.Error())
		// }

		// =========== Cara dengan kondisi Switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("validation error:", finalError.Error())
		default:
			fmt.Println("Unkown error:", finalError.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
