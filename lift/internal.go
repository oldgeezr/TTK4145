package lift

import (
	. ".././driver"
	. ".././network"
	// . ".././formating"
	// . "./log"
	. ".././functions"
	. "fmt"
	. "strconv"
	"time"
)

func Do_first(do_first chan Queues) {

	var last_floor int 
	// var temp int
	Fo.WriteString("Entered Do_first\n")

	for {
		select {
		case msg := <-do_first:
			// Fprintln(Fo, "TRASE ORDER: Mottok hele the_queue på do_first")
			job_queue := msg.Int_queue
			last_queue := msg.Last_queue
			ext_queue := msg.Ext_queue
			/*queues := Queues{job_queue,ext_queue,last_queue}
			last_queue = Determine_dir(job_queue, last_queue)
			Format_queues(queues)*/

			for _, last := range last_queue {
				if last.Ip_order == GetMyIP() {
					last_floor = last.Floor
					break
				}
			}
			
			if len(job_queue) != 0 {
				for _, yours := range job_queue {
					if yours.Ip == GetMyIP() {
						if len(yours.Dest) != 0 {
								Println("sending")
								// Fprintln(Fo, "TRASE ORDER: Sendte int ordre til send_to_floor")
								Send_to_floor(yours.Dest[0].Floor, last_floor,  "int")
						} else {
							if len(ext_queue) != 0 {
								Send_to_floor(ext_queue[0].Floor, last_floor, ext_queue[0].Dir)
							}
						}
					}
				}
			} else {
			}
		}
	}
}

//Sends elevator to specified floor
func Send_to_floor(floor, current_floor int, button string) {

	Fo.WriteString("Entered Send_to_floor\n")
	// current_floor := Get_floor_sensor()
	Elev_set_door_open_lamp(0)
	Set_stop_lamp(0)

	if current_floor < floor {
		// Fprintln(Fo, "TRASE ORDER: int ordre var skal opp")
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
				time.Sleep(1500*time.Millisecond)
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
	} else if current_floor > floor {
		// Fprintln(Fo, "TRASE ORDER: int ordre var skal ned")
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
				time.Sleep(1500*time.Millisecond)
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
		Println("Du er allerede her")
		// Fprintln(Fo, "DU ER ALLEREDE HER!")
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
			Println("Internal button nr: " + Itoa(i) + " has been pressed!")
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
	for {
		floor = Get_floor_sensor()
		if floor != -1 {
			Set_floor_indicator(floor)
			order <- Dict{GetMyIP(), floor, "standby"}
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
	Println("UP")

	go func() {
		for {

			floor = Get_floor_sensor()

			if floor != -1 {

				Speed(-150)
				Println("DOWN")
				time.Sleep(10 * time.Millisecond)
				Println("STOP")
				Speed(0)
				return
			}
		}
	}()

	go Floor_indicator(order)
	go Int_order(order)
	go Ext_order(order)
}
