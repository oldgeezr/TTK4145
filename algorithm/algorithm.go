package algorithm

import (
	. ".././functions"
	// . "fmt"
)

func Algo(algo_queues Queues, at_floor Dict) Queues {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var best int = 100
	var best_IP string = "nobest"
	var current_index int = 0
	var best_elevator int = 0

	job_queue := algo_queues.Job_queue
	ext_queue := algo_queues.Ext_queue
	last_queue := algo_queues.Last_queue

	// --------------------------------- Start: Got external order -----------------------------------------------------------------------------
	if at_floor.Ip_order == "ext" {
		//Finds which elevator is the closest
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
		//Found the best elevator in standby
		if best_IP != "nobest" {
			for i, yours := range job_queue {
				if yours.Ip == best_IP {
					job_queue[i].Dest, _ = Insert_at_pos("ip_order", job_queue[i].Dest, at_floor.Floor, 0)
					ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, best_IP)
					break
				}
			}
		} else {
			//All elevators are moving, which elevator is passing the floor
			for j, last := range last_queue {
				if len(job_queue[j].Dest) != 0 {
					// Println("ALGO:", job_queue[j].Dest[0].Floor, at_floor.Floor, at_floor.Floor, last.Floor, last.Dir, at_floor.Dir)
					if job_queue[j].Dest[0].Floor > at_floor.Floor && at_floor.Floor > last.Floor && last.Dir == at_floor.Dir {
						job_queue[j].Dest, _ = Insert_at_pos("ip_order", job_queue[j].Dest, at_floor.Floor, len(job_queue[j].Dest)-1)
						ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, last.Ip_order)
						break
					} else if job_queue[j].Dest[0].Floor < at_floor.Floor && at_floor.Floor < last.Floor && last.Dir == at_floor.Dir {
						job_queue[j].Dest, _ = Insert_at_pos("ip_order", job_queue[j].Dest, at_floor.Floor, len(job_queue[j].Dest)-1)
						ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, last.Ip_order)
						break
					}
				}
			}
		}
		//Algorithm safeguard
		/* Println("Appended: ", appended)
		if !appended {
			Println("Appended: ", appended)
			job_queue[0].Dest, _ = Insert_at_pos("ip_order", job_queue[0].Dest, ext_queue[0].Floor, len(job_queue[0].Dest)-1)
		}*/
	}
	// --------------------------------- End: Got external order ---------------------------------------------------------------------------

	// --------------------------------- Start: At new floor  ------------------------------------------------------------------------------
	if at_floor.Dir == "standby" || at_floor.Dir == "stop" {

		//Finds direction of elevator
		for _, last := range last_queue {
			if last.Ip_order == at_floor.Ip_order {
				last_dir = last.Dir
			}
		}
		//Finds correct queue-index
		for i, yours := range job_queue {
			if yours.Ip == at_floor.Ip_order {
				current_index = i
			}
		}
		//Safeguard
		if len(job_queue[current_index].Dest) == 0 {
			last_dir = "standby"
		}

		if Someone_getting_off(job_queue[current_index].Dest, at_floor.Floor) {
			if len(job_queue[current_index].Dest) != 0 {
				if job_queue[current_index].Dest[0].Floor == at_floor.Floor {
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					ext_queue = Remove_from_ext_queue(ext_queue, at_floor.Floor, last_dir)
				} else {
					// Rearrange
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					job_queue[current_index].Dest, _ = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
				}
			}
		}

		if Someone_getting_on(ext_queue, at_floor) {
			if len(job_queue[current_index].Dest) != 0 {
				if job_queue[current_index].Dest[0].Floor != at_floor.Floor {
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					job_queue[current_index].Dest, _ = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
				}
			} else {
				job_queue[current_index].Dest, _ = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
			}
			ext_queue = Remove_from_ext_queue(ext_queue, at_floor.Floor, last_dir)
		}

		// --------------------------------- Start: Send order to best elevator if all standby -------------------------------------------------
		var all_standby int = 0
		if len(ext_queue) != 0 {
			for i, last := range last_queue {
				if last.Dir == "standby" {
					temp := ext_queue[0].Floor - last.Floor
					if temp < 0 {
						temp = temp * (-1)
					}
					if temp < best {
						best = temp
						best_elevator = i
					}
					all_standby++
				}
			}
		}
		if all_standby == len(last_queue) && len(ext_queue) != 0 {
			job_queue[best_elevator].Dest, _ = Insert_at_pos("ip_order", job_queue[best_elevator].Dest, ext_queue[0].Floor, 0)
			ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, last.Ip_order)
		}
		// --------------------------------- End: Send order to best elevator if all standby ---------------------------------------------------

	}
	// --------------------------------- End: At new floor  ------------------------------------------------------------------------------------

	algo_queues = Queues{job_queue, ext_queue, last_queue}
	return algo_queues
}
