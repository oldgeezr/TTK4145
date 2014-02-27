package main

import (
	. "./lift/log"
	"encoding/json"
	. "fmt"
)

/*
new_job_queue := make(chan string)
	master_request := make(chan string)
	master_order := make(chan Dict)
	master_pop := make(chan string)
	algo_out := make(chan Order)
	last_floor := make(chan Dict)
	get_last_queue := make(chan []Dict)
	get_last_queue_request := make(chan bool)

	go Last_queue(last_floor, get_last_queue, get_last_queue_request, new_job_queue)
	go Job_queues(new_job_queue, master_request, master_pop, master_order, algo_out)

	go func() {
		for {
			select {
			case msg := <-get_last_queue:
				Println(msg)
			case msg := <-new_job_queue:
				Println(msg)
				get_last_queue_request <- true
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	go func() {
		for {
			select {
			case msg := <-master_order:
				Println(msg)
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}

	}()

	floor := Dict{"147", 1}
	last_floor <- floor
	time.Sleep(500 * time.Millisecond)

	floor = Dict{"142", 2}
	last_floor <- floor
	time.Sleep(500 * time.Millisecond)

	floor = Dict{"186", 3}
	last_floor <- floor
	time.Sleep(500 * time.Millisecond)

	floor = Dict{"147", 1}
	last_floor <- floor
	time.Sleep(500 * time.Millisecond)

	ord1 := Order{"147", 0, 3}
	ord2 := Order{"142", 0, 2}
	ord3 := Order{"186", 0, 3}
	ord4 := Order{"142", 1, 4}
	ord5 := Order{"142", 2, 9}

	algo_out <- ord1
	time.Sleep(500 * time.Millisecond)
	algo_out <- ord2
	time.Sleep(500 * time.Millisecond)
	algo_out <- ord3
	time.Sleep(500 * time.Millisecond)
	algo_out <- ord4
	time.Sleep(500 * time.Millisecond)
	algo_out <- ord5
	time.Sleep(500 * time.Millisecond)

	master_pop <- "142"

	neverQuit := make(chan string)
	<-neverQuit
*/

func main() {

	job_queue := []Jobs{}
	// job_queue = append(job_queue, Jobs{"0", []Slice{}})

	b, err := json.Marshal(job_queue)

	if err != nil {
		Println("first:", err)
	} else {
		var c []Jobs
		err2 := json.Unmarshal(b, &c)
		if err2 != nil {
			Println("second:", err2)
		} else {
			Println(c)
		}
	}

}
