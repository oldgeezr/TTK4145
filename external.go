package main

import (
	. "./driver"
	// . "fmt"
	"time"
)

// Stopp heis
func emergency_stop(ch chan bool) {

	for {
		if <-ch {
			Speed(-150)
			time.Sleep(25 * time.Millisecond)
			Speed(0)
			break
		}
	}
}

func order(ch2 chan int) {

	i := 0

	for {

		if i < 3 {
			if Get_button_signal(BUTTON_CALL_UP, i) == 1 {
				ch2 <- i
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_CALL_DOWN, i) == 1 {
				ch2 <- i
				time.Sleep(300 * time.Millisecond)
			}
		}

		i++
		i = i % 4

	}
}

func main() {

	// Initialize
	Init()
	Speed(150)

	// Stop at nearest floor
	last_floor := -1

	for {

		last_floor = Get_floor_sensor()

		if last_floor != -1 {

			Speed(-150)
			time.Sleep(25 * time.Millisecond)
			Speed(0)
			break
		}
	}

	// Get external destination and initialize emergency stop
	// Initialize emergency stop
	ch := make(chan bool)
	ch2 := make(chan int)

	go emergency_stop(ch)
	go order(ch2)

	for {

		next_floor := <-ch2

		if next_floor > last_floor {

			Speed(150)
			time.Sleep(500 * time.Millisecond)

			for {

				last_floor = Get_floor_sensor()

				if last_floor != -1 {

					Speed(-150)
					time.Sleep(25 * time.Millisecond)
					Speed(0)
					break
				}
			}
		} else if next_floor < last_floor {

			Speed(-150)
			time.Sleep(500 * time.Millisecond)

			for {

				last_floor = Get_floor_sensor()

				if last_floor != -1 {

					Speed(150)
					time.Sleep(25 * time.Millisecond)
					Speed(0)
					break
				}
			}
		}
	}
}
