package log

import (
	. "../.././algorithm"
	. "../.././formating"
	. "../.././functions"
	. "../.././lift"
	. "fmt"
	"sort"
)

func Job_queues(log_order chan Dict, slave_queues, queues_to_tcp, do_first chan Queues) {

	Fo.WriteString("Entered Job_queues\n")

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}

	the_queue := Queues{job_queue, ext_queue, last_queue}

	var new_value bool

	for {
		select {
		case msg := <-log_order:
			switch {

			case msg.Dir == "int":
				//Append to Correct Job_Queue
				Print("LOG: int job received on: ")
				Println(msg.Ip_order)
				job_queue = Append_if_missing_right_queue(job_queue, msg)
			case msg.Ip_order == "ext":
				//Append if missing to Ext_queue
				ext_queue, new_value = Append_if_missing_ext_queue(ext_queue, msg.Floor, msg.Dir)
				Println("LOG: appended ext queue: ", new_value)
			case msg.Floor >= M:
				Print("LOG: Elevator: ", msg.Ip_order)
				Println(" is moving -> updating direction!")
				//Update last_queue direction
				last_queue, _ = Update_Direction(last_queue, msg)
			case msg.Dir == "standby":
				//Creating job-queues
				Println("LOG: new floor detected!")
				if len(last_queue) != 0 {
					for _, last := range last_queue {
						if last.Ip_order != msg.Ip_order {
							job_queue, _ = Append_if_missing_queue(job_queue, msg.Ip_order)
						}
					}
				} else {
					job_queue, _ = Append_if_missing_queue(job_queue, msg.Ip_order)
					Println("LOG: New elevator detected: ", msg.Ip_order)
				}
				//Update last queue
				last_queue, _ = Append_if_missing_dict(last_queue, msg)
			}

			Println("LOG: updating the_queue with algo")

			the_queue = Queues{}
			the_queue = Queues{job_queue, ext_queue, last_queue}
			the_queue = Algo(the_queue, msg)

			job_queue = the_queue.Int_queue
			ext_queue = the_queue.Ext_queue
			last_queue = the_queue.Last_queue

			Format_queues_term(the_queue)
			queues_to_tcp <- the_queue //Send the_queue to all slaves

		case msg := <-slave_queues:
			the_queue = Queues{}
			the_queue.Int_queue = msg.Int_queue
			the_queue.Ext_queue = msg.Ext_queue
			the_queue.Last_queue = msg.Last_queue
			Format_queues_term(the_queue)
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
		// Fprintln(Fo, "Running, IP_array", )
		// Println(IPaddresses)
	}
}
