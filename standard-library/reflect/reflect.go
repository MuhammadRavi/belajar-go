package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value interface{}) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("\n Type Name:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)

		fmt.Printf("\n %s with type %s", structField.Name, structField.Type)
		fmt.Printf(" %s %s", structField.Tag.Get("required"), structField.Tag.Get("max"))
	}
}

func IsValid(value interface{}) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if !result {
				return result
			}
		}
	}

	return result
}

func main() {
	readField(Sample{"Eko"})
	fmt.Println()
	readField(Person{"Ravi", "", ""})

	person := Person{
		Name:    "Muhammad Ravi",
		Address: "Jakarta",
		Email:   "Tet",
	}
	fmt.Println("\n Is Valid: ", IsValid(person))
}
