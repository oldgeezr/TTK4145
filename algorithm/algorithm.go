package algorithm

import (
	. "../lift/log"
	"time"
)

func Algo(int_order chan Dict, get_last_queue chan []Dict, que_request chan bool, que chan []Jobs) {

	for {
		select {
		case last_queue := <-get_last_queue:
			que_request <- true
			time.Sleep(25 * time.Millisecond)
			int_queue := <-que
			for _, last := range last_queue {
				ip := last.Ip
				floor := last.Floor
				for _, inter := range int_queue {
					if ip == inter.Ip {
						for _, ord := range inter.Dest {
							if floor == ord {

							}
						}
					}
				}
			}
		}
	}
}
