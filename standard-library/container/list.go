package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Muhammad")
	data.PushBack("Ravi")
	data.PushBack("Mega")
	data.PushBack("Arasy")

	// ======= Manual =======
	// var head *list.Element = data.Front()
	// fmt.Println(head.Value)
	// fmt.Println(head.Next().Value)
	// fmt.Println(head.Next().Next().Value)
	// fmt.Println(head.Next().Next().Next().Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
