package main

import (
	// . "fmt"
	// "runtime"
	"./driver"
	"time"
)

// Stopper heisen på nærmeste etg
func stop_nearest_floor(ch2 chan bool) {

	for {

		if <-ch2 {
			driver.Speed(-150)
			time.Sleep(25 * time.Millisecond)
			driver.Speed(0)
			break
		}
	}
}

// Stopp heis
func emergency_stop(ch1 chan bool) {

	for {
		if <-ch1 {
			driver.Speed(-150)
			time.Sleep(25 * time.Millisecond)
			driver.Speed(0)
			break
		}
	}
}

func main() {

	// Initialize variables an stuff
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	// Initialize
	driver.Init()
	driver.Speed(150)

	// Start functions
	go emergency_stop(ch1)
	go stop_nearest_floor(ch2)

	for {

		if driver.Get_stop_signal() == 1 {
			driver.Set_stop_lamp(1)
			ch1 <- true
		}
		if driver.Get_floor_sensor() != -1 {
			ch1 <- true
		}
	}

	// neverQuit := make(chan string)
	// <-neverQuit
}
