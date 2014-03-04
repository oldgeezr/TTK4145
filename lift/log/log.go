package log

import (
	. "../.././functions"
	. "fmt"
	// . "strconv"
	"time"
)

/*func Last_queue(last_floor chan Dict, get_last_queue chan []Dict, get_last_queue_request chan bool, new_job_queue chan string) {

	last_queue := []Dict{}

	// i := 1
	// j := 1

	for {
		select {
		case msg := <-last_floor:
			missing_ip := true
			for i, last := range last_queue {
				if msg.Ip[1:] == last.Ip {
					missing_ip = false
					// Println("Fantes allerede:", j, "gang")
					// j++
					if msg.Floor != last.Floor {
						last_queue[i].Floor = msg.Floor
					}
				}
			}
			if missing_ip {
				msg.Ip = msg.Ip[1:]
				last_queue = append(last_queue, msg)
				new_job_queue <- msg.Ip
				// Println("Appendet:", i, "gang")
				// i++
			}
		case msg := <-get_last_queue_request:
			if msg {
				get_last_queue <- last_queue
			}
			// Må kanskje ha ein default med time sleep
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}*/

func Job_queues(que chan []Jobs, que_request chan bool, new_job_queue, master_request, master_pop chan string, master_order chan Dict, algo_out chan Order) {

	job_queue := []Jobs{}
	ext_queue := []Dict{}
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
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
			}
			Println(job_queue)
			Println(ext_queue)
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
