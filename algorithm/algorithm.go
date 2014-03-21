package algorithm

import (
	//. ".././formating"
	. "../functions"
	. "fmt"
)

func Algo(algo_queues Queues, at_floor Dict) Queues {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var best int = 100
	var best_IP string = "nobest"
	// var current_queue Jobs

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
			// current_queue = yours
			current_index = i
		}
	}
	if at_floor.Ip_order == "ext" {

		Println("I WAS HERE")

		for _, last := range last_queue {
			if last.Dir == "standby" {
				temp := ext_queue[0].Floor - last.Floor
				if temp < 0 {
					temp = temp * (-1)
				}
				if temp < best {
					best = temp
					best_IP = last.Ip_order
				}
			}
		}
		for i, yours := range int_queue {
			if yours.Ip == best_IP {
				int_queue[i].Dest = Insert_at_pos("ip_order", int_queue[i].Dest, at_floor.Floor, 0)
				// ext_queue = Remove_dict_ext_queue(ext_queue, ext_queue[0].Floor, "standby")
				break
			}
		}
		algo_queues = Queues{int_queue, ext_queue, last_queue}

	} else if at_floor.Dir == "standby" || at_floor.Dir == "stop" {

		if len(int_queue[current_index].Dest) == 0 {
			last_dir = "standby"
		}

		if Someone_getting_off(int_queue[current_index], at_floor.Floor) { // Noen skal av
<<<<<<< HEAD
			if int_queue[current_index].Dest[0].Floor == at_floor.Floor {
				int_queue[current_index] = Remove_int_queue(int_queue[current_index], at_floor.Floor)
				//ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, last_dir) //Tror denne kan fjernes
=======
			if len(int_queue[current_index].Dest) != 0 {
				if int_queue[current_index].Dest[0].Floor == at_floor.Floor {
					int_queue[current_index] = Remove_int_queue(int_queue[current_index], at_floor.Floor)
					ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, "standby")
				}
>>>>>>> e60ddac656d00a9c427e6d6da1daee0fde3c3347
			} else {
				// Re arrange
				int_queue[current_index] = Remove_int_queue(int_queue[current_index], at_floor.Floor)
				int_queue[current_index].Dest = Insert_at_pos("ip_order", int_queue[current_index].Dest, at_floor.Floor, 0)
			}
		}
		algo_queues = Queues{int_queue, ext_queue, last_queue}

		if Someone_getting_on(ext_queue, at_floor.Floor, last_dir) { // Noen skal på
			Println("ALGO: someone is getting on")
			Println("ALGO: ", current_index, int_queue)
			if len(int_queue[current_index].Dest) != 0 {
				Println("ALGO: have int_queue")
				if int_queue[current_index].Dest[0].Floor != at_floor.Floor {
					int_queue[current_index] = Remove_int_queue(int_queue[current_index], at_floor.Floor)
					int_queue[current_index].Dest = Insert_at_pos("ip_order", int_queue[current_index].Dest, at_floor.Floor, 0)
					Println(int_queue[current_index].Dest)
				}
			} else {
				Println("ALGO: ", current_index)
				int_queue[current_index].Dest = Insert_at_pos("ip_order", int_queue[current_index].Dest, at_floor.Floor, 0)
			}
			ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, last_dir)

		}

	}

	return algo_queues
}
