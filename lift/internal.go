package lift

import (
	. ".././driver"
	. ".././network"
	. "fmt"
	. "strconv"
	"time"
)

var current_floor int

func Wait_for_input(int_button chan int, int_order, ext_order chan string) {

	for {
		select {
		case floor := <-int_button:
			// current_floor = Get_floor_sensor()
			int_order <- Itoa(floor) + ":" + GetMyIP()
		default:
			time.Sleep(25 * time.Millisecond)
		}
	}

}

func Send_to_floor(floor int) {
	current_floor = Get_floor_sensor()
	if current_floor < floor {
		Println("Going up")
		for {
			Speed(150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_floor_indicator(Get_floor_sensor())
				Set_stop_lamp(1)
				Speed(0)
				break
			}
		}
	} else {
		Println("Going down")
		for {
			Speed(-150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_floor_indicator(Get_floor_sensor())
				Set_stop_lamp(1)
				Speed(0)
				break
			}
		}
	}
}

func KeyboardInput(int_button chan int) {
	var a int

	for {
		Scan(&a)
		int_button <- a
	}
}

func Order(int_button chan int) {

	i := 0

	for {

		if i < 3 {
			if Get_button_signal(BUTTON_COMMAND, i) == 1 {
				//Println("Button nr: " + Itoa(i) + " has been pressed!")
				int_button <- i
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_COMMAND, i) == 1 {
				//Println("Button nr: " + Itoa(i) + " has been pressed!")
				int_button <- i
				time.Sleep(300 * time.Millisecond)
			}
		}

		i++
		i = i % 4

	}
}

func Internal(int_button chan int, int_order, ext_order chan string) {

	// Initialize
	Init()
	Speed(0)
	Set_stop_lamp(1)

	go KeyboardInput(int_button)
	go Wait_for_input(int_button, int_order, ext_order)

	neverQuit := make(chan string)
	<-neverQuit
}
