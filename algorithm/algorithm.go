package algorithm

import (
	//. ".././formating"
	. ".././functions"
	. "fmt"
)

func Algo(algo_queues Queues, at_floor Dict) Queues {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var best int = 100
	var elevator int = 0
	var best_IP string = "nobest"
	// var current_queue Jobs

	job_queue := algo_queues.Job_queue
	ext_queue := algo_queues.Ext_queue
	last_queue := algo_queues.Last_queue

	for _, last := range last_queue {
		if last.Ip_order == at_floor.Ip_order {
			last_dir = last.Dir
		}
	}

	current_index := -1

	for i, yours := range job_queue { // GÃ¥ gjennom alle jobbkÃ¸ene
		if yours.Ip == at_floor.Ip_order { // Finn riktig jobbkÃ¸
			// current_queue = yours
			current_index = i
		}
	}
	if at_floor.Ip_order == "ext" {

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
		for i, yours := range job_queue {
			if yours.Ip == best_IP {
				if !Someone_getting_off(job_queue[i], at_floor.Floor) {
					job_queue[i].Dest = Insert_at_pos("ip_order", job_queue[i].Dest, at_floor.Floor, 0)
				}
				// ext_queue = Remove_dict_ext_queue(ext_queue, ext_queue[0].Floor, "standby")
				break
			} else {
				if !Someone_getting_off(job_queue[i], at_floor.Floor) {
					elevator = elevator % len(last_queue)
					job_queue[elevator].Dest = append(job_queue[elevator].Dest, Dict{"ip_order", at_floor.Floor, "int"})
					elevator++
				}
			}
		}
		algo_queues = Queues{job_queue, ext_queue, last_queue}

	} else if at_floor.Dir == "standby" || at_floor.Dir == "stop" {

		if len(job_queue[current_index].Dest) == 0 {
			last_dir = "standby"
		}

		if Someone_getting_off(job_queue[current_index], at_floor.Floor) { // Noen skal av
			if len(job_queue[current_index].Dest) != 0 {
				if job_queue[current_index].Dest[0].Floor == at_floor.Floor {
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, "standby")
				}
			} else {
				// Re arrange
				job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
				job_queue[current_index].Dest = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
			}
		}
		algo_queues = Queues{job_queue, ext_queue, last_queue}

		if Someone_getting_on(ext_queue, at_floor.Floor, last_dir) { // Noen skal pÃ¥
			Println("ALGO: someone is getting on")
			Println("ALGO: ", current_index, job_queue)
			if len(job_queue[current_index].Dest) != 0 {
				Println("ALGO: have job_queue")
				if job_queue[current_index].Dest[0].Floor != at_floor.Floor {
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					job_queue[current_index].Dest = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
					Println(job_queue[current_index].Dest)
				}
			} else {
				Println("ALGO: ", current_index)
				job_queue[current_index].Dest = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
			}
			ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, last_dir)

		}

	}

	return algo_queues
}
