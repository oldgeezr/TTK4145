package algorithm

import (
	. "../functions"
	. "../lift/log"
	"time"
)

/*func Algo(int_order chan Dict, get_last_queue chan []Dict, que_request chan bool, que chan []Jobs) {

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
}*/

func Algo(int_order, at_floor chan Dict, get_last_queue chan []Dict, que_request chan bool, que chan []Jobs) {

	for {
		select {
		case last_floor := <-at_floor: // får inn siste etg fra ip
			job_queue := <-que
			for _, find_lift := range job_queue {
				if find_lift.Ip == last_floor.Ip {
					for _, floor := range find_lift.Dest {
						if last_floor.Floor == floor {
							// Noen skal av
							// Stop heis
							// Fjern alle etg i heis kø som er likt siste etgs

						} else {
							// Ingen skal av
							// spør om
						}
					}
				}
			}

		}
	}
}
