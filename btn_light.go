package main

import (
	. "./driver"
	// . "fmt"
)

func main() {

	// Initialize
	Init()
	Speed(0)

	for {

		signal := Get_stop_signal()

		if signal == 1 {
			Set_stop_lamp(1)
		} else {
			Set_stop_lamp(0)
		}
	}
}
