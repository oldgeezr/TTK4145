package main

import (
	. "./lift/log"
	. "fmt"
)

func main() {

	last_queue := []Dict{}
	last_queue = append(last_queue, Dict{"147", 2}, Dict{"151", 1}, Dict{"146", 3})
	Println(last_queue)

	pop, last_queue := Pop_first(last_queue)

	Println(pop, ":", last_queue)

}
