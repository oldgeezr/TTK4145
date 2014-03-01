package main

import (
	. "./functions"
	. "fmt"
)

func main() {

	a := []Slice{Slice{1}}
	a = append(a, Slice{2})
	a = append(a, Slice{3})
	a = append(a, Slice{2})
	a = append(a, Slice{4})

	Println(a)
	a = Insert(a, 2, 2)
	Println(a)
}
