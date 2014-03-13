package algorithm

import (
	. ".././formating"
	. "../functions"
	. "fmt"
	"time"
)

func Algo(get_at_floor chan Dict, get_queues chan Queues) {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var current_queue Jobs

	for {
		at_floor := <-get_at_floor
		queues := <-get_queues

		int_queue := queues.Int_queue
		ext_queue := queues.Ext_queue
		last_queue := queues.Last_queue

		for _, last := range last_queue {
			if last.Ip_order == at_floor.Ip_order {
				last_dir = last.Dir
			}
		}

		current_index := -1

		for i, yours := range int_queue { // Gå gjennom alle jobbkøene
			if yours.Ip == at_floor.Ip_order { // Finn riktig jobbkø
				current_queue = yours
				current_index = i
			}
		}

		Format_queues_term(queues)

		if !Missing_ext_job(ext_queue, at_floor.Floor, last_dir) { // Noen skal på
			Println("STAGE 1")
			if len(int_queue[current_index].Dest) != 0 {
				Println("STAGE 2")
				if int_queue[current_index].Dest[0].Floor != at_floor.Floor {
					Println("STAGE 3")
					int_queue[current_index] = Remove_order_int_queue(int_queue[current_index], at_floor.Floor)
					Println("STAGE 4")
				}
			} else {
				Println("STAGE 5:", int_queue[current_index].Dest, at_floor.Floor)
				int_queue[current_index].Dest = Insert_at_pos("ip_order", int_queue[current_index].Dest, at_floor.Floor, 0)
				Println("STAGE 6", int_queue[current_index].Dest)
			}
			ext_queue = Remove_order_ext_queue(ext_queue, at_floor.Floor, last_dir)
		}

		queues = Queues{int_queue, ext_queue, last_queue}
		Println("STAGE 6.5:")
		Format_queues_term(queues)

		if !Missing_int_job(current_queue, at_floor.Floor) { // Noen skal av
			Println("STAGE 7")
			if len(current_queue.Dest) != 0 {
				Println("STAGE 8")
				if current_queue.Dest[0].Floor == at_floor.Floor {
					Println("STAGE 9")
					// Skal fjerne etg fra internjobbkø
					int_queue[current_index] = Remove_order_int_queue(int_queue[current_index], at_floor.Floor)
					Println("STAGE 10")
					ext_queue = Remove_order_ext_queue(ext_queue, at_floor.Floor, last_dir)
					Println("STAGE 11")
				}
			} else {
				// Re arrange
				int_queue[current_index] = Remove_order_int_queue(int_queue[current_index], at_floor.Floor)
				Println("STAGE 12")
				int_queue[current_index].Dest = Insert_at_pos("ip_order", int_queue[current_index].Dest, at_floor.Floor, 0)
				Println("STAGE 13")
			}
		}

		queues = Queues{int_queue, ext_queue, last_queue}
		get_queues <- queues
		time.Sleep(50 * time.Millisecond)
		Println("STAGE 13.5:")
		Format_queues_term(queues)
	}
}
