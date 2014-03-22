package algorithm

import (
	. ".././functions"
)

func Algo(algo_queues Queues, at_floor Dict) Queues {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var best int = 100
	var best_IP string = "nobest"
	var current_index int = -1
	var best_elevator int = 0
	var appended bool

	job_queue := algo_queues.Job_queue
	ext_queue := algo_queues.Ext_queue
	last_queue := algo_queues.Last_queue

	switch {
	// --------------------------------- Start: Got external order -----------------------------------------------------------------------------
	case at_floor.Ip_order == "ext":

		// --------------------------------- Start: Determine best standby elevator  -----------------------------------------------------------
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
				for j, _ := range ext_queue {
					// if ext.Floor == at_floor.Floor && ext.Dir == at_floor.Dir && ext.Ip_order == "ext" {
					job_queue[i].Dest, appended = Insert_at_pos("ip_order", job_queue[i].Dest, at_floor.Floor, 0)
					if appended {
						ext_queue[j].Ip_order = best_IP
					}
					break
					// }
				}
				break
				// --------------------------------- End: Determine best standby elevator ------------------------------------------------------

				// --------------------------------- Start: Determine best moving elevator -----------------------------------------------------
			} else if best_IP == "nobest" {
				//for _, ext := range ext_queue {
				// if ext.Floor == at_floor.Floor && ext.Dir == at_floor.Dir && ext.Ip_order == "ext" {
				for j, last := range last_queue {
					if len(job_queue[j].Dest) != 0 {
						if job_queue[j].Dest[0].Floor > at_floor.Floor && last.Dir == "down" {
							job_queue[j].Dest, appended = Insert_at_pos("ip_order", job_queue[j].Dest, at_floor.Floor, len(job_queue[j].Dest)-1)
							if appended {
								ext_queue[j].Ip_order = best_IP
							}
						} else if job_queue[j].Dest[0].Floor < at_floor.Floor && last.Dir == "up" {
							job_queue[j].Dest, appended = Insert_at_pos("ip_order", job_queue[j].Dest, at_floor.Floor, len(job_queue[j].Dest)-1)
							if appended {
								ext_queue[j].Ip_order = best_IP
							}
						}
					}
				}
				// }
				//}
			}
			// --------------------------------- Start: Determine best moving elevator ---------------------------------------------------------
		}
		// --------------------------------- End: Got external order ---------------------------------------------------------------------------

		// --------------------------------- Start: At new floor  ------------------------------------------------------------------------------
	case at_floor.Dir == "standby" || at_floor.Dir == "stop":

		// --------------------------------- Start: Find best_elevator direction and correct index ---------------------------------------------
		for _, last := range last_queue {
			if last.Ip_order == at_floor.Ip_order {
				last_dir = last.Dir
			}
		}

		for i, yours := range job_queue {
			if yours.Ip == at_floor.Ip_order {
				current_index = i
			}
		}
		// --------------------------------- End: Find best_elevator direction and correct index -----------------------------------------------

		if len(job_queue[current_index].Dest) == 0 {
			last_dir = "standby"
		}

		// --------------------------------- Start: Someone getting off? (internal order) ------------------------------------------------------
		if Someone_getting_off(job_queue[current_index].Dest, at_floor.Floor) {
			if len(job_queue[current_index].Dest) != 0 {
				if job_queue[current_index].Dest[0].Floor == at_floor.Floor {
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					ext_queue = Remove_from_ext_queue(ext_queue, at_floor.Floor, last_dir)
				} else {
					// REARRANGE
					job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
					job_queue[current_index].Dest, _ = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
				}
			}
		}
		// --------------------------------- End: Someone getting off? (internal order) --------------------------------------------------------

		// --------------------------------- Start: Someone getting on? (external order) -------------------------------------------------------
		if Someone_getting_on(ext_queue, at_floor.Floor, last_dir) {
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
		// --------------------------------- End: Someone getting on? (external order) ---------------------------------------------------------

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
		}
		// --------------------------------- End: Send order to best elevator if all standby ---------------------------------------------------
	}
	// --------------------------------- End: At new floor  ------------------------------------------------------------------------------------

	algo_queues = Queues{job_queue, ext_queue, last_queue}
	return algo_queues
}
