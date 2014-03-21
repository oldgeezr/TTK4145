package lift

import (
	. ".././driver"
	. ".././functions"
	. ".././network"
	. "./log"
	"time"
)

func Do_first(do_first chan Queues, order chan Dict) {

	Fo.WriteString("Entered Do_first\n")

	// HACK
	time.Sleep(500 * time.Millisecond)

	var last_floor int
	var myIP string = GetMyIP()

	state := make(chan string)

	go Send_to_floor(state, order)

	for {
		queues := <-do_first
		job_queue := queues.Job_queue

		if Get_floor_sensor() != -1 {
			last_floor = Get_floor_sensor()
		}
		if len(job_queue) != 0 {
			for _, yours := range job_queue {
				if yours.Ip == myIP {
					if len(yours.Dest) != 0 {
						if yours.Dest[0].Floor > last_floor {
							state <- "up"
						} else if yours.Dest[0].Floor < last_floor {
							state <- "down"
						} else {
							state <- "stop"
						}
					} else {
						state <- "standby"

					}
				}
			}
		}
	}
}

func Send_to_floor(state chan string, order chan Dict) {

	var last_dir string
	var floor int
	var myIP string = GetMyIP()

	Elev_set_door_open_lamp(0)
	Set_stop_lamp(0)

	for {
		st := <-state
		Elev_set_door_open_lamp(0)
		if Get_floor_sensor() != -1 {
			floor = Get_floor_sensor()
		}

		switch {

		case st == "up":
			Speed(150)
			if last_dir != "up" {
				order <- Dict{myIP, M + 1, "up"}
			}
			last_dir = "up"

		case st == "down":
			Speed(-150)
			if last_dir != "down" {
				order <- Dict{myIP, M + 1, "down"}
			}
			last_dir = "down"

		case st == "stop":
			if last_dir == "up" {
				Speed(-150)
			} else if last_dir == "down" {
				Speed(150)
			}
			time.Sleep(25 * time.Millisecond)
			Speed(0)
			Elev_set_door_open_lamp(1)
			Set_button_lamp(BUTTON_COMMAND, floor, 0)
			order <- Dict{myIP, floor, "stop"}
			time.Sleep(1500 * time.Millisecond)
			last_dir = "stop"

		case st == "standby":
			Speed(0)
			if last_dir != "standby" {
				order <- Dict{myIP, M + 1, "standby"}
			}
			last_dir = "standby"
		}
	}
}

func External_btn_order(order chan Dict) {

	Fo.WriteString("Entered Ext_order\n")

	var i int = 0

	for {
		if i < M-1 {
			if Get_button_signal(BUTTON_CALL_UP, i) == 1 {
				order <- Dict{"ext", i, "up"}
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_CALL_DOWN, i) == 1 {
				order <- Dict{"ext", i, "down"}
				time.Sleep(500 * time.Millisecond)
			}
		}
		i++
		i = i % M
		time.Sleep(25 * time.Millisecond)
	}
}

func Internal_btn_order(order chan Dict) {

	Fo.WriteString("Entered Int_order\n")

	var i int = 0

	for {
		if Get_button_signal(BUTTON_COMMAND, i) == 1 {
			Set_button_lamp(BUTTON_COMMAND, i, 1)
			order <- Dict{GetMyIP(), i, "int"}
			time.Sleep(500 * time.Millisecond)
		}
		i++
		i = i % M
		time.Sleep(25 * time.Millisecond)
	}
}

func Floor_indicator(order chan Dict) {

	Fo.WriteString("Entered Floor_indicator\n")

	var floor int
	var last_floor int = M + 1

	for {
		floor = Get_floor_sensor()
		if floor != -1 && floor != last_floor {
			Set_floor_indicator(floor)
			order <- Dict{GetMyIP(), floor, "standby"}
			last_floor = floor
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func External_lights(do_first chan Queues) {

	Fo.WriteString("Entered External_lights\n")

	var queue Queues
	for {
		queue = <-do_first
		for _, external := range queue.Ext_queue {
			if external.Dir == "up" {
				Set_button_lamp(BUTTON_CALL_UP, external.Floor, 1)
			} else {
				Set_button_lamp(BUTTON_CALL_DOWN, external.Floor, 1)
			}
		}
		time.Sleep(100 * time.Millisecond)

		Set_button_lamp(BUTTON_CALL_UP, 0, 0)
		Set_button_lamp(BUTTON_CALL_UP, 1, 0)
		Set_button_lamp(BUTTON_CALL_UP, 2, 0)

		Set_button_lamp(BUTTON_CALL_DOWN, 3, 0)
		Set_button_lamp(BUTTON_CALL_DOWN, 2, 0)
		Set_button_lamp(BUTTON_CALL_DOWN, 1, 0)
	}
}

func Lift_init(do_first chan Queues, order chan Dict) {

	Fo.WriteString("Entered lift\n")

	var floor int = -1

	Init()
	Speed(150)

	go func() {
		for {
			floor = Get_floor_sensor()
			if floor != -1 {
				Speed(-150)
				time.Sleep(10 * time.Millisecond)
				Speed(0)
				return
			}
		}
	}()

	go Floor_indicator(order)
	go Internal_btn_order(order)
	go External_btn_order(order)
	go External_lights(do_first)
	go Do_first(do_first, order)
}
