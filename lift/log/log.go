package log

import (
	. "../.././functions"
	. "fmt"
)

func Job_queues(master_order, get_at_floor chan Dict, queues, get_queues, do_first chan Queues) {

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}
	the_queue := Queues{job_queue, ext_queue}

	for {
		select {
		case msg := <-master_order:
			if msg.Dir == "int" {
				job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
				for i, job := range job_queue {
					if job.Ip == msg.Ip_order {
						job_queue[i].Dest, _ = AIM_Int(job_queue[i].Dest, msg.Floor)
					}
				}
				the_queue = Queues{job_queue, ext_queue}
				queues <- the_queue
				do_first <- the_queue
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
				the_queue = Queues{job_queue, ext_queue}
				queues <- the_queue
				do_first <- the_queue
			} else if msg.Dir == "last" {
				Println("I wass hererererer")
				var update bool 
				last_queue, update = AIM_Dict(last_queue, msg)
				Println(last_queue)
				if update {
					// get_at_floor <- msg
					Println(update)
				}
			}
		case msg := <-queues:
			the_queue = msg
			job_queue = msg.Int_queue
			ext_queue = msg.Ext_queue
			do_first <- the_queue
		case msg := <-get_queues:
			the_queue = msg
		case get_queues <- the_queue:
		}
	}
}
