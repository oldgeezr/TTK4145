package algorithm

import (
	. "../lift/log"
	"time"
)

func Algo(que chan []Jobs, que_request chan bool, int_order, ext_order chan Dict, algo_out chan Order) {

	for {
		select {
		case msg := <-int_order:
			go func() {
				job_queue := <-que
				for _, queue := range job_queue {
					if queue.Ip == msg.Ip {
						for i := range queue.Dest {
							if 
						}
					}
				}
			}()
			que_request <- true
		case msg := <-ext_order:

		default:
			time.Sleep(50 * time.Millisecond)
		}
	}

}
