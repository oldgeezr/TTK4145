package main

import (
	. "./formating"
	. "./functions"
	// . "./lift"
	. "./lift/log"
	. "./network"
	. "./network/tcp"
	. "./network/udp"
	. "fmt"
	. "net"
	"os"
	. "strconv"
	"time"
)

func main() {

	Elevator_art()
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

	// --------------------------------- Start: Listen for network activity -------------------------------------
	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(250 * time.Millisecond))

	b := make([]byte, 16)

	_, _, err2 := ln.ReadFromUDP(b)
	ln.Close()
	// --------------------------------- End: Listen for network activity ---------------------------------------

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
	queues := make(chan Queues)
	get_queues := make(chan Queues)
	do_first := make(chan Queues)
	get_at_floor := make(chan Dict)
	slave_queues := make(chan Queues)
	kill_IMA_master := make(chan bool)
	set_queues := make(chan Queues)
	// --------------------------------- End: Create system channels --------------------------------------------

	// --------------------------------- Start: Common program threads ------------------------------------------
	go IP_array(ip_array_update, get_ip_array, flush)
	go Timer(flush)
	go Job_queues(log_order, get_at_floor, queues, get_queues, set_queues, slave_queues, do_first)
	go IMA(udp)
	go UDP_listen(ip_array_update)
	// go Internal(order)
	// go Do_first(do_first, order)
	// --------------------------------- End: Common program threads --------------------------------------------

	// --------------------------------- Start: System state maching --------------------------------------------
	go func() {
		for {
			select {
			case <-master:
				Println("=> State: Entered master state")
				Fo.WriteString("=> State: Entered master state\n")
				udp <- true
				go TCP_master_connect(log_order, slave_queues)
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
				go IMA_master(get_ip_array, master, new_master, kill_IMA_master)
				go func() { new_master <- true }()
			case <-new_master:
				ip := <-get_ip_array
				if len(ip) != 0 {
					if ip[len(ip)-1] > 255 {
						master_ip := Itoa(ip[len(ip)-1] - 255)
						if master_ip != GetMyIP() {
							go func() { new_master <- TCP_slave_com(master_ip, order, queues) }()
							// Det som er litt unødvendig er at ny master har en TCP med seg selv..
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
	// --------------------------------- End: System state maching -----------------------------------------------

	// --------------------------------- Start: Set state --------------------------------------------------------
	if err2 != nil {
		master <- true
		Fo.WriteString("I am master\n")
	} else {
		slave <- true
		Fo.WriteString("I am slave\n")
	}
	// --------------------------------- End: Set state ----------------------------------------------------------

	neverQuit := make(chan string)
	<-neverQuit
}
