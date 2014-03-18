package log

import (
	// . "../.././formating"
	. "../.././algorithm"
	. "../.././functions"
	. "../.././lift"
	"sort"
	// . "fmt"
)

func Job_queues(log_order, get_at_floor chan Dict, queues, get_queues, set_queues, slave_queues, do_first chan Queues) {

	Fo.WriteString("Entered Job_queues\n")

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}

	var the_queue Queues

	for {
		select {
		case msg := <-log_order:
			switch {

			case msg.Dir == "int":
				//Append to Correct Job_Queue
				job_queue = Append_if_missing_right_queue(job_queue, msg)
			case msg.Ip_order == "ext":
				//Append if missing to Ext_queue
				ext_queue, _ = Append_if_missing_ext_queue(ext_queue, msg.Floor, msg.Dir)
			case msg.Floor >= M:
				//Update last_queue direction
				last_queue, _ = Update_Direction(last_queue, msg)
			case msg.Dir == "standby":
				//Creating job-queues
				if len(last_queue) != 0 {
					for _, last := range last_queue {
						if last.Ip_order != msg.Ip_order {
							job_queue, _ = Append_if_missing_queue(job_queue, msg.Ip_order)
						}
					}
				} else {
					job_queue, _ = Append_if_missing_queue(job_queue, msg.Ip_order)
				}
				//Update last queue
				last_queue, _ = Append_if_missing_dict(last_queue, msg)
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

func IP_array(ip_array_update chan int, get_ip_array chan []int, flush chan bool) {

	Fo.WriteString("Entered IP_array\n")

	IPaddresses := []int{}

	for {
		select {
		case ip := <-ip_array_update:
			IPaddresses = Append_if_missing_ip(IPaddresses, ip)
			sort.Ints(IPaddresses)
		case get_ip_array <- IPaddresses:
		case msg := <-flush:
			_ = msg
			IPaddresses = IPaddresses[:0]
		}
	}
}
