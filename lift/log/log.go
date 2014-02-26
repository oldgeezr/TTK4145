package log

/*const (
	N int = 3 // Antall heiser
)*/

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

func Insert(this []Slice, pos Order.Pos, floor Order.Floor) []Slice {

	// HÅper dette funker
	temp_slice := []Slice{}
	temp_slice = this[:pos-1]
	temp_slice = append(temp_slice, floor)
	temp_slice = append(temp_slice, this[pos:])
	return temp_slice
}

func Last_queue(last_floor, get_last_queue chan Dict, new_job_queue chan string, algo_out chan Order) {

	last_queue := []Dict{}

	for {
		select {
		case msg := <-last_floor:
			missing := false
			for _, last := range last_queue {
				if msg.Ip == last.Ip {
					last.Floor = msg.Floor
				} else {
					missing = true
				}
			}
			if missing {
				last_queue = append(last_queue, Dict{msg})
				new_job_queue <- msg.Ip
			}
		case get_last_queue <- last_queue:
			// Må kanskje ha ein default med time sleep
		}
	}
}

func Job_queues(new_job_queue, master_request, master_pop chan string, algo_out chan Order) {

	job_queue := []Jobs{}

	for {
		select {
		case ip := <-new_job_queue:
			// Opprett ny kø på gitt ip
			job_queue = append(job_queue, Jobs{ip, []Slice{}})
		case Do := <-algo_out:
			// Legg til beslutning fra algo i rett jobb kø
			for _, queue := range job_queue {
				if queue.Ip == ip {
					queue = Insert(queue, Do.Pos, Do.Floor)
					master_request <- queue.Dest[0]
				}
			}
		case ip := <-master_request:
			// Send ny ordre fra riktig kø til master
			for _, queue := range job_queue {
				if queue.Ip == ip {
					master_request <- queue.Dest[0]
				}
			}
		case ip := <-master_pop:
			// pop ordre fra kø, da den er fullførrt
			for _, queue := range job_queue {
				if queue.Ip == ip {
					queue = Pop_first(queue)
				}
			}
		}
	}
}
