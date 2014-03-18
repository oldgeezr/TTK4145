package log

import (
	. "../.././formating"
	. "../.././functions"
	. "../.././network"
	. "fmt"
)

func Job_queues(log_order, get_at_floor chan Dict, queues, get_queues, set_queues, slave_queues, do_first chan Queues) {

	Fo.WriteString("Entered Job_queues\n")

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}

	var the_queue Queues
	var algo_queue Queues

	for {
		select {
		case msg := <-log_order:
			switch {

			case msg.Dir == "int":
				//Append to Correct Job_Queue
				job_queue = ARQ(job_queue, msg)
			case msg.Ip_order == "ext":
				//Append if missing to Ext_queue
				ext_queue, _ = AIM_Ext(ext_queue, msg.Floor, msg.Dir)
			case msg.Floor >= M:
				//Update last_queue direction
				last_queue, _ = Update_Direction(last_queue, msg)
			case msg.Dir == "standby":
				//Creating job-queues
				if len(last_queue) != 0 {
					for _, last := range last_queue {
						if last.Ip_order != msg.Ip_order {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
						}
					}
				} else {
					job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
				}
				//Update last queue
				last_queue, _ = AIM_Dict(last_queue, msg)
			}

			if msg.Dir == "standby" || msg.Dir == "stop" {
				the_queue = Algo(the_queue, msg)
			} else {
				the_queue = Queues{job_queue, ext_queue, last_queue}
			}
			slave_queues <- the_queue //Send the_queue to all slaves

		case msg := <-queues:
			the_queue = Queues{msg.Int_queue, msg.Ext_queue, msg.Last_queue}
		case do_first <- the_queue: // DO FIRST
		}
	}
}

func ARQ(queue []Jobs, msg Dict) []Jobs {
	for i, job := range queue {
		if job.Ip == msg.Ip_order {
			queue[i].Dest, _ = AIM_Int(queue[i].Dest, msg.Floor)
		}
	}
	return queue
}

func Determine_dir(job_queue []Jobs, last Dict) string {
	for _, job := range job_queue {
		if last.Ip_order == job.Ip {
			if len(job.Dest) != 0 {
				if job.Dest[0].Floor-last.Floor > 0 {
					return "up"
				} else if job.Dest[0].Floor-last.Floor < 0 {
					return "down"
				} else {
					return "standby"
				}
			} else {
				return "standby"
			}
		}
	}
	return "standby"
}
