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
	var best_IP string = "nobest"
	var current_index int = -1
	var best_elevator int = 0

	job_queue := algo_queues.Job_queue
	ext_queue := algo_queues.Ext_queue
	last_queue := algo_queues.Last_queue

	switch {

	case at_floor.Ip_order == "ext":

		for _, last := range last_queue {
			if last.Dir == "standby" {
				temp := ext_queue[0].Floor - last.Floor
				if temp < 0 {
					temp = temp * (-1)
				}
				if temp < best {
					best = temp
					best_IP = last.Ip_order
					Println("BEST IP:", best_IP)
				}
			}
		}
		for i, yours := range job_queue {
			if yours.Ip == best_IP {
				if !Someone_getting_off(job_queue[i], at_floor.Floor) {
					for i, ext := range ext_queue {
						if ext.Floor == at_floor.Floor && ext.Dir == at_floor.Dir && ext.Ip_order != "taken" {
							ext_queue[i].Ip_order = "taken"
							job_queue[i].Dest = Insert_at_pos("ip_order", job_queue[i].Dest, at_floor.Floor, 0)
							break
						}
					}
				}
				break
			} else if best_IP == "nobest" {
				if !Someone_getting_off(job_queue[i], at_floor.Floor) {

					best_elevator = best_elevator % len(last_queue)

					for i, ext := range ext_queue {
						if ext.Floor == at_floor.Floor && ext.Dir == at_floor.Dir && ext.Ip_order != "taken" {
							ext_queue[i].Ip_order = "taken"
							job_queue[best_elevator].Dest = append(job_queue[best_elevator].Dest, Dict{"ip_order", at_floor.Floor, "int"})
							Println("JOB SENDT TO JOB_QUEUE_index:", i)
							break
						}
					}

					best_elevator++
				}
			}
		}

	case at_floor.Dir == "standby" || at_floor.Dir == "stop":

		// --------------------------------- Start: Finds the best_elevator direction -----------------------------------------------------------
		for _, last := range last_queue {
			if last.Ip_order == at_floor.Ip_order {
				last_dir = last.Dir
			}
		}
		// --------------------------------- Start: Finds the best_elevator direction -----------------------------------------------------------

		// --------------------------------- Start: Finds the correct job_queue index ------------------------------------------------------
		for i, yours := range job_queue {
			if yours.Ip == at_floor.Ip_order {
				current_index = i
			}
		}
		// --------------------------------- Start: Finds the correct job_queue index ------------------------------------------------------

		// --------------------------------- Start: If best_elevator has no jobs, it must be in standby -----------------------------------------
		if len(job_queue[current_index].Dest) == 0 {
			last_dir = "standby"
		}
		// --------------------------------- End: If best_elevator has no jobs, it must be in standby -------------------------------------------

		// --------------------------------- Start: Is there a floor in job_queue that is equal to this floor ------------------------------
		if Someone_getting_off(job_queue[current_index], at_floor.Floor) {
			Println("SOMEONE IS GETTING OFF!")
			if len(job_queue[current_index].Dest) != 0 {
				if job_queue[current_index].Dest[0].Floor == at_floor.Floor {
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, "standby")
				} else {
					// Re arrange
					Println("REARRANGING!")
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					job_queue[current_index].Dest = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
				}
			}
		}
		// --------------------------------- End: Is there a floor in job_queue that is equal to this floor --------------------------------

		// --------------------------------- Start: Is there a floor in ext_queue that is equal to this floor ------------------------------
		if Someone_getting_on(ext_queue, at_floor.Floor, last_dir) {
			if len(job_queue[current_index].Dest) != 0 {
				if job_queue[current_index].Dest[0].Floor != at_floor.Floor {
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					job_queue[current_index].Dest = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
				}
			} else {
				job_queue[current_index].Dest = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
			}
			ext_queue = Remove_dict_ext_queue(ext_queue, at_floor.Floor, last_dir)

		}
		// --------------------------------- End: Is there a floor in ext_queue that is equal to this floor --------------------------------
	}

	algo_queues = Queues{job_queue, ext_queue, last_queue}

	return algo_queues
}
