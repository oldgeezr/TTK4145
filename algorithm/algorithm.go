package algorithm

import (
	. "../functions"
	// . "../lift/log"
	. "fmt"
	"time"
)

func Algo(get_at_floor, algo_out chan Dict, get_int_queue chan []Jobs, get_ext_queue chan []Dict) {

	for {
		select {
		case at_floor := <-get_at_floor:
			int_queue := <-get_int_queue
			ext_queue := <-get_ext_queue
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
					algo_out <- ext_queue[0]
					break // Avslutt å gå gjennom køen fordi det er unødvendig da det kun finnes en instans av hver heis
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
