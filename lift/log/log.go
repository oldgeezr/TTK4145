package log

import (
	// . "fmt"
	// . "strconv"
	"time"
)

type Dict struct {
	Ip    string
	Floor int
}

type Slice struct {
	A int
}

type Jobs struct {
	Ip   string
	Dest []Slice
}

type Order struct {
	Ip    string
	Pos   int
	Floor int
}

func Pop_first(this []Slice) []Slice {

	return this[1:len(this)]
}

func Insert(that []Slice, new_order Order) []Slice {

	pos := new_order.Pos
	floor := new_order.Floor
	temp_slice := []Slice{}
	rest_slice := []Slice{}
	if pos > len(that) || pos == 0 {
		temp_slice = append(temp_slice, Slice{new_order.Floor})
	} else {
		temp_slice = append(temp_slice, that[:pos-1]...)
		rest_slice = append(rest_slice, that[pos-1:]...)
		temp_slice = append(temp_slice, Slice{floor})
		temp_slice = append(temp_slice, rest_slice...)
	}
	return temp_slice
}
func Last_queue(last_floor chan Dict, get_last_queue chan []Dict, get_last_queue_request chan bool, new_job_queue chan string) {

	last_queue := []Dict{}

	// i := 1
	// j := 1

	for {
		select {
		case msg := <-last_floor:
			missing := true
			for i, last := range last_queue {
				if msg.Ip == last.Ip {
					last_queue[i] = msg
					missing = false
					// Println("Fantes allerede:", j, "gang")
					// j++
				}
			}
			if missing {
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
}

func Job_queues(new_job_queue, master_request, master_pop chan string, master_order chan Dict, algo_out chan Order) {

	job_queue := []Jobs{}
	// job_queue = append(job_queue, Jobs{"0", []Slice{}})

	for {
		select {
		case ip := <-new_job_queue:
			// Opprett ny kø på gitt ip
			job_queue = append(job_queue, Jobs{ip, []Slice{}})
		case Do := <-algo_out:
			// Legg til beslutning fra algo i rett jobb kø
			for i, queue := range job_queue {
				if queue.Ip == Do.Ip {
					job_queue[i].Dest = Insert(queue.Dest, Do)
					master_order <- Dict{Do.Ip, Do.Floor}
					// Println(job_queue)
				}
			}
		case ip := <-master_request:
			// Send ny ordre fra riktig kø til master
			for _, queue := range job_queue {
				if queue.Ip == ip {
					master_order <- Dict{queue.Ip, queue.Dest[0].A}
				}
			}
		case ip := <-master_pop:
			// pop ordre fra kø, da den er fullførrt
			for i, queue := range job_queue {
				if queue.Ip == ip {
					job_queue[i].Dest = Pop_first(queue.Dest)
					// Println(job_queue)
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
