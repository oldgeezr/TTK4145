package lift

import (
	. ".././driver"
	. ".././network"
	. "./log"
	. "fmt"
	. "strconv"
	"time"
)

func Master_get_last_queue(get_last_queue chan []Dict, master_order chan Dict) {

	for {
		select {
		case msg := <-get_last_queue:
			Println(msg)
		case msg := <-master_order:
			Println(msg.Floor)
		default:
			time.Sleep(50 * time.Microsecond)
		}
	}
}

func Master_print_last_queue(get_last_queue_request chan bool, master_request chan string, algo_out chan Order) {

	for {

		get_last_queue_request <- true
		time.Sleep(time.Second)
		algo_out <- Order{"141", 3, 1}
		time.Sleep(time.Second)
		master_request <- "141"
		time.Sleep(time.Second)

	}
}

func Master_input(int_order, ext_order, last_floor chan Dict) {

	for {
		select {
		case msg := <-int_order:
			Print(msg)
		case msg := <-ext_order:
			Print(msg)
		case msg := <-last_floor:
			_ = msg
		default:
			time.Sleep(25 * time.Millisecond)
		}
	}
}

//Sends elevator to specified floor
func Send_to_floor(floor int, button string) {
	current_floor := Get_floor_sensor()
	Elev_set_door_open_lamp(0)
	Set_stop_lamp(0)

	if current_floor < floor {
		Println("Going up")
		for {
			Speed(150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_stop_lamp(1)
				Elev_set_door_open_lamp(1)
				Speed(-150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				if button == "int" {
					Set_button_lamp(BUTTON_COMMAND, floor, 0)
				} else {
					if button == "up" {
						Set_button_lamp(BUTTON_CALL_UP, floor, 0)
					} else {
						Set_button_lamp(BUTTON_CALL_DOWN, floor, 0)
					}
				}
				return
			}
			time.Sleep(25 * time.Millisecond)
		}
	} else {
		Println("Going down")
		for {
			Speed(-150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_stop_lamp(1)
				Elev_set_door_open_lamp(1)
				Speed(150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				if button == "int" {
					Set_button_lamp(BUTTON_COMMAND, floor, 0)
				} else {
					if button == "up" {
						Set_button_lamp(BUTTON_CALL_UP, floor, 0)
					} else {
						Set_button_lamp(BUTTON_CALL_DOWN, floor, 0)
					}
				}
				return
			}
			time.Sleep(25 * time.Millisecond)
		}
	}
}

//Keyboard terminal input (For testing)
func KeyboardInput(ch chan int) {
	var a int

	for {
		Scan(&a)
		ch <- a
	}
}

//Handles external button presses
func Ext_order(ext_order chan Dict) {

	i := 0

	for {

		if i < 3 {
			if Get_button_signal(BUTTON_CALL_UP, i) == 1 {
				Println("External call up button nr: " + Itoa(i) + " has been pressed!")
				Set_button_lamp(BUTTON_CALL_UP, i, 1)
				ext_order <- Dict{"up", i}
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_CALL_DOWN, i) == 1 {
				Println("External call down button nr: " + Itoa(i) + " has been pressed!")
				Set_button_lamp(BUTTON_CALL_DOWN, i, 1)
				ext_order <- Dict{"down", i}
				time.Sleep(300 * time.Millisecond)
			}
		}

		i++
		i = i % 4
		time.Sleep(25 * time.Millisecond)

	}
}

//Handles internal button presses
func Int_order(int_order chan Dict) {

	i := 0
	for {
		if Get_button_signal(BUTTON_COMMAND, i) == 1 {
			Println("Internal button nr: " + Itoa(i) + " has been pressed!")
			Set_button_lamp(BUTTON_COMMAND, i, 1)
			int_order <- Dict{GetMyIP(), i}
			time.Sleep(300 * time.Millisecond)
		}

		i++
		i = i % 4
		time.Sleep(25 * time.Millisecond)

	}
}

//Checks which floor the elevator is on and sets the floor-light
func Floor_indicator(last_floor chan Dict) {
	Println("executing floor indicator!")
	var floor int
	for {
		floor = Get_floor_sensor()
		if floor != -1 {
			Set_floor_indicator(floor)
			last_floor <- Dict{"X" + GetMyIP(), floor}
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func To_nearest_floor() {
	for {
		Speed(150)
		if Get_floor_sensor() != -1 {
			time.Sleep(25 * time.Millisecond)
			Speed(0)
		}
	}
}

func Internal(int_order, ext_order, last_floor chan Dict) {

	// Initialize
	Init()
	Speed(150)
	floor := -1

	go func() {
		for {

			floor = Get_floor_sensor()

			if floor != -1 {

				Speed(-150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				return
			}
		}
	}()

	go Floor_indicator(last_floor)
	go Int_order(int_order)
	go Ext_order(ext_order)
}
