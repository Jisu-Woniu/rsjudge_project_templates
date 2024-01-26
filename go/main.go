package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/lists/arraylist"
)

func main() {
	fmt.Println("Hello World!")
	list := lists.List.Empty(arraylist.New())
	fmt.Println(list)
}
