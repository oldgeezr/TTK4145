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

func Missing_ext_job(job_queue []Dict, floor int, dir string) bool {

	for _, orders := range job_queue {
		if orders.Dir == dir && orders.Floor == floor {
			return false
		}
	}
	return true
}

func Remove_order_ext_queue(this []Dict, floor int, dir string) []Dict {

	for i, orders := range this {
		if orders.Dir == dir && orders.Floor == floor {
			this = this[:i+copy(this[i:], this[i+1:])]
		}
	}
	return this
}

func Remove_order_int_queue(this Jobs, floor int) Jobs {

	for i, orders := range this.Dest {
		if orders.Floor == floor {
			this.Dest = this.Dest[:i+copy(this.Dest[i:], this.Dest[i+1:])]
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

	// External job queue
	ext_queue := []Dict{}
	ext_queue = append(ext_queue, Dict{"ext", 3, "up"}, Dict{"ext", 1, "down"}, Dict{"ext", 1, "up"})
	// @ floor
	at_floor := Dict{"147", 1, "down"}

	// hei stopp flagg
	// stop_lift := false

	/*for _, order := range job_queue {
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

			}

		}
	}*/

	Println("@floor:", at_floor)
	Println("int_queue:", int_queue)
	Println("ext_qeueu:", ext_queue)

	for _, order := range int_queue {
		if order.Ip == at_floor.Ip_order { // Finn riktig kø
			if !Missing_int_job(order, at_floor.Floor) { // Noen skal av
				// Stopp heis
				Println("queue before remove:", order)
				order = Remove_order_int_queue(order, at_floor.Floor)
				Println("queue after remove:", order) // Slett alle interne
				ext_queue = Remove_order_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir)
				Println("ext_queue after remove:", ext_queue) // Slett alle eksterne i riktig retning
			} else { // Ingen skal av
				if !Missing_ext_job(ext_queue, at_floor.Floor, at_floor.Dir) { // Noen skal på
					// Stopp heis
					Println("I was here?")
					ext_queue = Remove_order_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir) // Slett alle eksterne i riktig retning
				}
				Println("I was here??")
			}
			break // Avslutt å gå gjennom køen fordi det er unødvendig da det kun finnes en instans av hver heis
		}
	}

}
