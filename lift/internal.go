package lift

import (
	. ".././driver"
	//. ".././formating"
	. ".././functions"
	. ".././network"
	// . "./log"
	. "fmt"
	. "strconv"
	"time"
)

func Do_first(do_first chan Queues, order chan Dict) {

	var last_floor int
	myIP := GetMyIP()
	state := make(chan string)

	go Send_to_floor(state, order)

	for {
		//time.Sleep(50 * time.Millisecond)
		queues := <-do_first

		//Format_queues_term(queues, "Do_first")

		job_queue := queues.Int_queue
		ext_queue := queues.Ext_queue
		last_queue := queues.Last_queue

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
							Fprintf(Fo, "JOB:stop:\n")
						}
					} else {
						if len(ext_queue) != 0 {

							if Determine_best_elevator(ext_queue, last_queue, myIP) {
								if ext_queue[0].Floor > last_floor {
									state <- "up"
								} else if ext_queue[0].Floor < last_floor {
									state <- "down"
								} else {
									state <- "stop"
									Fprintf(Fo, "Ext:stop\n")
								}
							} else {
								state <- "standby"
								Fprintf(Fo, "Ext:standby\n")
							}
						}
					}
				}
			}
		}

	}
}

//Sends elevator to specified floor
func Send_to_floor(state chan string, order chan Dict) {

	var last_dir string
	var floor int
	myIP := GetMyIP()

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
			Println("STOP!")
			if last_dir == "up" {
				Speed(-150)
				time.Sleep(25 * time.Millisecond)
				// Set_button_lamp(BUTTON_CALL_UP, floor, 0)
			} else if last_dir == "down" {
				Speed(150)
				time.Sleep(25 * time.Millisecond)
				// Set_button_lamp(BUTTON_CALL_DOWN, floor, 0)
			}
			Speed(0)
			Elev_set_door_open_lamp(1)
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

//Keyboard terminal input (For testing)
func KeyboardInput(ch chan int) {
	var a int

	for {
		Scan(&a)
		ch <- a
	}
}

//Handles external button presses
func Ext_order(order chan Dict) {

	Fo.WriteString("Entered Ext_order\n")

	i := 0

	for {

		if i < 3 {
			if Get_button_signal(BUTTON_CALL_UP, i) == 1 {
				Println("External call up button nr: " + Itoa(i) + " has been pressed!")
				Set_button_lamp(BUTTON_CALL_UP, i, 1)
				order <- Dict{"ext", i, "up"}
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_CALL_DOWN, i) == 1 {
				Println("External call down button nr: " + Itoa(i) + " has been pressed!")
				Set_button_lamp(BUTTON_CALL_DOWN, i, 1)
				order <- Dict{"ext", i, "down"}
				time.Sleep(300 * time.Millisecond)
			}
		}

		i++
		i = i % 4
		time.Sleep(25 * time.Millisecond)

	}
}

//Handles internal button presses
func Int_order(order chan Dict) {

	Fo.WriteString("Entered Int_order\n")

	i := 0
	for {
		if Get_button_signal(BUTTON_COMMAND, i) == 1 {
			// Println("Internal button nr: " + Itoa(i) + " has been pressed!")
			Set_button_lamp(BUTTON_COMMAND, i, 1)
			order <- Dict{GetMyIP(), i, "int"}
			Fprintln(Fo, "INTERNAL: btn -> order -> tcp")
			time.Sleep(300 * time.Millisecond)
		}

		i++
		i = i % 4
		time.Sleep(25 * time.Millisecond)

	}
}

//Checks which floor the elevator is on and sets the floor-light
func Floor_indicator(order chan Dict) {

	Fo.WriteString("Entered Floor_indicator\n")

	Println("executing floor indicator!")
	var floor int
	var last_floor int = 100
	for {
		floor = Get_floor_sensor()
		if floor != -1 && floor != last_floor {
			Set_floor_indicator(floor)
			order <- Dict{GetMyIP(), floor, "standby"}
			last_floor = floor
			// Fprintln(Fo, "222: @floor -> order -> tcp")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func To_nearest_floor() {

	Fo.WriteString("Entered To_nearest_floor\n")

	for {
		Speed(150)
		if Get_floor_sensor() != -1 {
			time.Sleep(25 * time.Millisecond)
			Speed(0)
		}
	}
}

func Internal(order chan Dict) {

	Fo.WriteString("Entered Internal\n")

	// Initialize
	Init()
	Speed(150)
	floor := -1

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
	go Int_order(order)
	go Ext_order(order)
}
