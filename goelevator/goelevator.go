package goelevator

import (
	. ".././driver"
	. ".././formatting"
	. ".././functions"
	. ".././interrupts"
	. ".././lift"
	. ".././lift/log"
	. ".././network"
	. ".././network/tcp"
	. ".././network/udp"
	. "fmt"
	. "net"
	"os"
	. "strconv"
	"time"
)

func Go_elevator() {

	var err error

	// --------------------------------- Start: Create error log ------------------------------------------------
	Fo, err = os.Create("output.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := Fo.Close(); err != nil {
			panic(err)
		}
	}()
	// --------------------------------- End: Create error log --------------------------------------------------

	go Interrupts()

	// --------------------------------- Start: Create system channels ------------------------------------------
	slave := make(chan bool)
	master := make(chan bool)
	udp := make(chan bool)
	ip_array_update := make(chan int)
	get_ip_array := make(chan []int)
	new_master := make(chan bool)
	flush := make(chan bool)
	order := make(chan Dict)
	log_order := make(chan Dict)
	slave_queues := make(chan Queues)
	do_first := make(chan Queues)
	queues_to_tcp := make(chan Queues)
	kill_IMA_master := make(chan bool)
	lost_conn := make(chan bool)
	// --------------------------------- End: Create system channels --------------------------------------------

	Elevator_art()

	// --------------------------------- Start: Listen for network activity -------------------------------------
	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(250 * time.Millisecond))

	b := make([]byte, 16)

	_, _, err2 := ln.ReadFromUDP(b)
	ln.Close()
	// --------------------------------- End: Listen for network activity ---------------------------------------

	// --------------------------------- Start: Common program threads ------------------------------------------
	go IP_array(ip_array_update, get_ip_array, flush)
	go Flush_IP_array(flush)
	go Job_queues(log_order, slave_queues, queues_to_tcp, do_first)
	go IMA(udp)
	go UDP_listen(ip_array_update)
	go Lift_init(do_first, order)
	go Got_net_connection(lost_conn, true)
	// --------------------------------- End: Common program threads --------------------------------------------

	// --------------------------------- Start: System state machine --------------------------------------------
	go func() {
		for {
			select {
			case <-master:
				Println("=> State: Entered master state")
				Fo.WriteString("=> State: Entered master state\n")
				udp <- true
				go TCP_master_connect(log_order, queues_to_tcp, get_ip_array)
				go func() {
					for {
						msg := <-order
						log_order <- msg
					}
				}()
			case <-slave:
				Println("=> State: Entered slave state")
				Fo.WriteString("=> State: Entered slave state\n")
				udp <- false
				go Is_master_alive(get_ip_array, master, new_master, kill_IMA_master)
				go func() { new_master <- true }()
			case <-new_master:
				ip := <-get_ip_array
				if len(ip) != 0 {
					if ip[len(ip)-1] > 255 {
						master_ip := Itoa(ip[len(ip)-1] - 255)
						if master_ip != GetMyIP() {
							go func() { new_master <- TCP_slave_com(master_ip, order, slave_queues) }()
						} else {
							kill_IMA_master <- true
							master <- true
						}
					} else {
						go func() { new_master <- true }()
					}
				} else {
					go func() { new_master <- true }()
				}
			}
		}

	}()
	// --------------------------------- End: System state machine -----------------------------------------------

	// --------------------------------- Start: Set state --------------------------------------------------------
	if err2 != nil {
		master <- true
	} else {
		slave <- true
	}
	// --------------------------------- End: Set state ----------------------------------------------------------

	// --------------------------------- Start: Lost net connection => crash program -----------------------------
	for {
		select {
		case <-lost_conn:
			Print("LOST CONNECTION")
			Speed(0)
			return
		}
	}
	// --------------------------------- End: Lost net connection => crash program -------------------------------
}
