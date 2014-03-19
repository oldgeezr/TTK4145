package algorithm

import (
	. ".././formating"
	. "../functions"
	. "fmt"
)

func Algo(algo_queues Queues, at_floor Dict) Queues {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var current_queue Jobs

	int_queue := algo_queues.Int_queue
	ext_queue := algo_queues.Ext_queue
	last_queue := algo_queues.Last_queue

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

	if at_floor.Dir == "standby" || at_floor.Dir == "stop" {
		Println("Before algorithm:")
		Format_queues_term(algo_queues)
		if Someone_getting_on(ext_queue, at_floor.Floor, last_dir) { // Noen skal på

			Println("Stage: 1")
			if len(int_queue[current_index].Dest) != 0 {
				Println("Stage: 2")
				if int_queue[current_index].Dest[0].Floor != at_floor.Floor {
					int_queue[current_index] = Remove_int_queue(int_queue[current_index], at_floor.Floor)
					Println("Stage: 3", current_index, int_queue)
				}
			} else {
				int_queue[current_index].Dest = Insert_at_pos("ip_order", int_queue[current_index].Dest, at_floor.Floor, 0)
				Println("Stage: 4", int_queue)
			}
			ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, last_dir)
			Println("Stage: 5", ext_queue)

		}

		if Someone_getting_off(current_queue, at_floor.Floor) { // Noen skal av
			if len(current_queue.Dest) != 0 {
				if current_queue.Dest[0].Floor == at_floor.Floor {
					int_queue[current_index] = Remove_int_queue(int_queue[current_index], at_floor.Floor)
					ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, last_dir)
				}
			} else {
				// Re arrange
				int_queue[current_index] = Remove_int_queue(int_queue[current_index], at_floor.Floor)
				int_queue[current_index].Dest = Insert_at_pos("ip_order", int_queue[current_index].Dest, at_floor.Floor, 0)
			}
		}
		algo_queues = Queues{int_queue, ext_queue, last_queue}
		Println("After algorithm:")
		Format_queues_term(algo_queues)
	}

	return algo_queues
}
