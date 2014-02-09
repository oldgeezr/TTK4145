package main

import (
	. "fmt"
	//"runtime"
	. "./driver"
	"time"
)

func main() {

	// Initialiser
	Init()
	Speed(150)
	last_floor := -1

	// Gå opp og ned
	for {

		floor := Get_floor_sensor()
		Println(floor)

		switch floor {
		case 0:
			Speed(150)
			time.Sleep(25 * time.Millisecond)
			Speed(0)
			time.Sleep(1000 * time.Millisecond)
			Speed(150)
			time.Sleep(1000 * time.Millisecond)
			last_floor = floor

		case 1:
			if last_floor > floor {
				Speed(150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(-150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			} else {
				Speed(-150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			}
		case 2:
			if last_floor > floor {
				Speed(150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(-150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			} else {
				Speed(-150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			}
		case 3:
			Speed(-150)
			time.Sleep(25 * time.Millisecond)
			Speed(0)
			time.Sleep(1000 * time.Millisecond)
			Speed(-150)
			time.Sleep(1000 * time.Millisecond)
			last_floor = floor
		}
	}
}
