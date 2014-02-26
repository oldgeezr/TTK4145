package log

// NOTE TO SELF!
// This works!
/*

	ext_queue := []slice{}
	job_queue := []jobs{}

	ext_queue = append(ext_queue, slice{2}, slice{3})
	job_queue = append(job_queue, jobs{ext_queue})

	Println(ext_queue)

*/

import (
	. "fmt"
)

const (
	N int = 3 // Antall heiser
)

type dist struct {
	ip    string
	floor int
}

type slice struct {
	a int
}

type jobs struct {
	ip string
	b  []slice
}

func Test_queue(last_floor, int_order chan dist, ext_order chan dist) {

	last_queue := []dist{}
	int_queue := []dist{}
	ext_queue := []slice{}
	job_queue := []jobs{}

	last_queue = append(last_queue, dist{"147", 2}, dist{"151", 1}, dist{"146", 3})

	for _, lifts := range last_queue {

		job_queue = append(job_queue, jobs{lifts.ip, []slice{}})
	}

	for {
		select {
		case msg := <-last_floor:
			missing := false
			for _, last := range last_queue {
				if msg.ip == last.ip {
					last.floor = msg.floor
				} else {
					missing = true
				}
			}
			if missing {
				last_queue = append(last_queue, dist{msg})
			}
		case msg := <-int_order:
			int_queue = append(int_queue, dist{msg})
		case msg := <-ext_order:
			ext_queue = append(ext_queue, dist{msg})
		}
	}

}
