package algorithm

import (
	. "../functions"
	. "fmt"
)

func Algo(get_at_floor chan Dict, get_queues chan Queues) {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var current_dict int

	for {
		select {
		case at_floor := <-get_at_floor:
			queues := <-get_queues
			int_queue := queues.Int_queue
			ext_queue := queues.Ext_queue
			last_queue := queues.Last_queue

			for _, last := range last_queue {
				if last.Ip_order == at_floor.Ip_order {
					last_dir = last.Dir
				}
			}

			for i, order := range int_queue {
				if order.Ip == at_floor.Ip_order { // Finn riktig kø
					if !Missing_int_job(order, at_floor.Floor) { // Noen skal av
						if int_queue[i].Dest[0].Floor == at_floor.Floor {
							current_dict = i
							int_queue[i] = Remove_order_int_queue(int_queue[i], at_floor.Floor)
							ext_queue = Remove_order_ext_queue(ext_queue, at_floor.Floor, last_dir)
						} else {
							Println("QUEUE:", int_queue)
							int_queue[i].Dest = Insert_at_pos(int_queue[i].Ip, int_queue[i].Dest, at_floor.Floor, 0)
							Println("QUEUE2:", int_queue)
						}
					}
				}
			}

			if !Missing_ext_job(ext_queue, at_floor.Floor, last_dir) {
				Println("ALGO:", ext_queue, at_floor.Floor, last_dir)
				int_queue[current_dict].Dest = Insert_at_pos(int_queue[current_dict].Ip, int_queue[current_dict].Dest, at_floor.Floor, 0)
				ext_queue = Remove_order_ext_queue(ext_queue, at_floor.Floor, last_dir)
				Println("ALGO2:", ext_queue, at_floor.Floor, last_dir)
			}

			// Avslutt å gå gjennom køen fordi det er unødvendig da det kun finnes en instans av hver heis
			queues = Queues{int_queue, ext_queue, last_queue}
			Println("ALGO3:", queues)
			get_queues <- queues
		}
	}
}
