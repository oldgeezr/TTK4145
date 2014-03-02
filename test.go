package main

import (
	. "fmt"
)

type Dict struct {
	Ip_order string
	Floor    int
	Dir      string
}

type Slice struct {
	Floor Dict
}

type Jobs struct {
	Ip   string
	Dest []Dict
}

func Missing_int_job(job_queue Jobs, floor int) bool {

	for _, orders := range job_queue.Dest {
		if orders.Floor == floor && orders.Ip_order == "int" {
			return false
		}
	}
	return true
}

func Remove_order_int_queue(this Jobs, floor int) Jobs {

	for i, orders := range this.Dest {
		if orders.Floor == floor {
			this.Dest = this.Dest[:i+copy(this.Dest[i:], this.Dest[i+1:])]
			Println(this)
		}
	}
	return this

}

func main() {

	// Last queue
	// last_queue := []Dict{}
	// last_queue = append(last_queue, Dict{"147", 1, "up"}, Dict{"154", 3, "down"}, Dict{"161", 0, "standby"})

	// Internal job queue
	int_queue := []Jobs{}
	floor := []Dict{}
	int_queue = append(int_queue, Jobs{"147", append(floor, Dict{"int", 1, "standby"})}, Jobs{"154", append(floor, Dict{"int", 2, "standby"})})

	/*
		// External job queue
		ext_queue := []Dict{}
		ext_queue = append(ext_queue, Dict{"ext", 3, "up"}, Dict{"ext", 1, "down"}, Dict{"ext", 1, "up"})
		// @ floor
		at_floor := last_queue[0]

		// hei stopp flagg
		stop_lift := false

		for _, order := range job_queue {
			for _, int_order := range order.Dest {
				if int_order.Floor.Ip == "int" || int_order.Floor.Floor == at_floor.Floor {
					// Noen skal av
					// Stop heis
					stop_lift = true
					// Fjern alle ordre i denne etg for alle heiser
					// Pop ordre
				}
				if stop_lift {
					// fjern alle ordre i denne etg
				} else {
					if int_order.Floor.Ip ==  || int_order.Floor.Ip == "down" || int_order.Floor.Floor == at_floor.Floor {
					}
				}

			}
		}
	*/

	Println(int_queue)
	if Missing_int_job(int_queue[0], 1) {
		int_queue[0] = Remove_order_int_queue(int_queue[0], 1)
		Println(int_queue)
	}

}
