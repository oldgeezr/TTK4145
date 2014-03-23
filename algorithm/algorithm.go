package algorithm

import (
	. ".././functions"
)

func Algo(algo_queues Queues, at_floor Dict) Queues {

	Fo.WriteString("Entered Algo\n")

	var last_dir string
	var best int = 100
	var best_IP string = "nobest"
	var current_index int = 0
	var appended bool

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
					job_queue[i].Dest, appended = Insert_at_pos("ip_order", job_queue[i].Dest, at_floor.Floor, 0)
					ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, best_IP)
					break
				}
			}
		} else {
			//All elevators are moving, which elevator is passing the floor
			for j, last := range last_queue {
				if len(job_queue[j].Dest) != 0 {
					if job_queue[j].Dest[0].Floor > at_floor.Floor && at_floor.Floor > last.Floor && last.Dir == at_floor.Dir {
						job_queue[j].Dest, appended = Insert_at_pos("ip_order", job_queue[j].Dest, at_floor.Floor, len(job_queue[j].Dest))
						ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, last.Ip_order)
						break
					} else if job_queue[j].Dest[0].Floor < at_floor.Floor && at_floor.Floor < last.Floor && last.Dir == at_floor.Dir {
						job_queue[j].Dest, appended = Insert_at_pos("ip_order", job_queue[j].Dest, at_floor.Floor, len(job_queue[j].Dest))
						ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, last.Ip_order)
						break
					}
				}
			}
		}

		if !appended {

			var length int = 100
			var shortest_queue int = 0

			for i, jobs := range job_queue {
				temp := len(jobs.Dest)
				if temp < length {
					length = temp
					shortest_queue = i
				}
			}

			if len(job_queue[shortest_queue].Dest) != 0 {
				job_queue[shortest_queue].Dest, _ = Insert_at_pos("ip_order", job_queue[shortest_queue].Dest, at_floor.Floor, len(job_queue[shortest_queue].Dest))
			} else {
				job_queue[shortest_queue].Dest, _ = Insert_at_pos("ip_order", job_queue[shortest_queue].Dest, at_floor.Floor, 0)
			}
			ext_queue = Mark_ext_queue(ext_queue, at_floor.Floor, at_floor.Dir, last_queue[shortest_queue].Ip_order)
		}
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
					ext_queue = Remove_from_ext_queue(ext_queue, at_floor.Floor, job_queue[current_index].Ip)
					ext_queue = Remove_from_ext_queue(ext_queue, at_floor.Floor, job_queue[current_index].Ip)
				} else {
					// Rearrange
					for _, order := range ext_queue {
						if job_queue[current_index].Ip == order.Ip_order && last_dir == order.Dir && order.Floor == at_floor.Floor {
							job_queue[current_index] = Remove_job_queue(job_queue[current_index], at_floor.Floor)
							job_queue[current_index].Dest, _ = Insert_at_pos("ip_order", job_queue[current_index].Dest, at_floor.Floor, 0)
						}
					}
				}
			}
		}
	}
	// --------------------------------- End: At new floor  ------------------------------------------------------------------------------------

	algo_queues = Queues{job_queue, ext_queue, last_queue}
	return algo_queues
}
