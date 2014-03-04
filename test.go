package main

import (
	. "./algorithm"
	. "./functions"
	. "fmt"
)

func main() {

	// Last queue
	// last_queue := []Dict{}
	// last_queue = append(last_queue, Dict{"147", 1, "up"}, Dict{"154", 3, "down"}, Dict{"161", 0, "standby"})

	// Internal job queue
	int_queue := []Jobs{}
	floor := []Dict{}
	int_queue = append(int_queue, Jobs{"147", append(floor, Dict{"int", 1, "standby"})}, Jobs{"154", append(floor, Dict{"int", 2, "standby"})})

	// External job queue
	ext_queue := []Dict{}
	ext_queue = append(ext_queue, Dict{"ext", 3, "up"}, Dict{"ext", 1, "down"}, Dict{"ext", 1, "up"})
	// @ floor
	at_floor := Dict{"147", 1, "down"}

	get_at_floor := make(chan Dict)
	algo_out := make(chan Dict)
	get_int_queue := make(chan []Jobs)
	get_ext_queue := make(chan []Dict)

	go Algo(get_at_floor, algo_out, get_int_queue, get_ext_queue)

	Println("@floor:", at_floor)
	Println("int_queue:", int_queue)
	Println("ext_qeueu:", ext_queue)

	get_at_floor <- at_floor
	get_int_queue <- int_queue
	get_ext_queue <- ext_queue

	for {
		Println(<-algo_out)
		break
	}

}
