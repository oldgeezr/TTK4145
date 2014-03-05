package log

import (
	. "../.././functions"
	// . "fmt"
	// . "strconv"
)

func Last_queue(order chan Dict) {

	last_queue := []Dict{}

	// i := 1
	// j := 1

	for {
		select {
		case msg := <-order:

	}
}

func Job_queues(master_order chan Dict, queues, do_first chan Queues) {

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}
	the_queue := Queues{job_queue, ext_queue}
	// job_queue = append(job_queue, Jobs{"0", []Slice{}})

	for {
		select {
		/*case ip := <-new_job_queue:
			// Opprett ny kø på gitt ip
			job_queue = append(job_queue, Jobs{ip, []Dict{}})
		/*case Do := <-algo_out:
		// Legg til beslutning fra algo i rett jobb kø
		for i, queue := range job_queue {
			if queue.Ip == Do.Ip {
				job_queue[i].Dest = Insert_at_pos(job_queue[i].Dest, Do)
				master_order <- Dict{Do.Ip, Do.Floor}
				// Println(job_queue)
			}
		}
		case ip := <-master_request:
			// Send ny ordre fra riktig kø til master
			for _, queue := range job_queue {
				if queue.Ip == ip {
					master_order <- Dict{queue.Ip, queue.Dest[0].Floor, "care2"}
				}
			}
		case msg := <-que_request:
			if msg {
				que <- job_queue
			}
		case ip := <-master_pop:
			// pop ordre fra kø, da den er fullførrt
			for i, queue := range job_queue {
				if queue.Ip == ip {
					job_queue[i].Dest = Pop_first(queue.Dest)
					// Println(job_queue)
				}
			}*/
		case msg := <-master_order:
			if msg.Dir == "int" {
				job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
				for i, job := range job_queue {
					if job.Ip == msg.Ip_order {
						job_queue[i].Dest, _ = AIM_Dict(job_queue[i].Dest, msg.Floor)
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
				
			}
		case msg := <-queues:
			the_queue = msg
			job_queue = msg.Int_queue
			ext_queue = msg.Ext_queue
			do_first <- the_queue
		}
	}
}
