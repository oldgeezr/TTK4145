package log

import (
	// . "../.././formating"
	. "../.././functions"
	. "../.././network"
	. "fmt"
)

func Job_queues(log_order, get_at_floor chan Dict, queues, get_queues, set_queues, slave_queues, do_first chan Queues) {

	Fo.WriteString("Entered Job_queues\n")

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}
	the_queue := Queues{job_queue, ext_queue, last_queue}

	for {
		select {
		case msg := <-log_order:
			switch {
			case msg.Dir == "int":
				job_queue = ARQ(job_queue, msg)
			case msg.Ip_order == "ext":
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
			case msg.Floor >= M:
				var update bool
				last_queue, update = AIM_Dict2(last_queue, msg)
				if update {
					get_at_floor <- msg
				}
			case msg.Dir == "standby":
				Println("UPDATE LASTFLOOR")
				if len(last_queue) != 0 {
					for _, last := range last_queue {
						if last.Ip_order != msg.Ip_order {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
						}
					}
				} else {
					job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
				}
				var update bool
				last_queue, update = AIM_Dict(last_queue, msg)
				if update {
					get_at_floor <- msg
					Println("Lastfloor update")
				}
			case msg.Dir == "remove":
				get_at_floor <- msg
				Println("Removing")
			}
			the_queue = Queues{job_queue, ext_queue, last_queue}
			slave_queues <- the_queue
		case msg := <-set_queues:
			the_queue = msg
			slave_queues <- the_queue
		case msg := <-queues:
			the_queue = msg
			Println("UPDATEING MYE VERSION OF LOG")
		case do_first <- the_queue: // DO FIRST
			Println("TRYING TO FETCH QUEUES")
		case get_queues <- the_queue: // ALGO
			the_queue = Queues{}
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
