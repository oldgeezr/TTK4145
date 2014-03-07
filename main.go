package main

import (
	. "./algorithm"
	. "./functions"
	. "./lift"
	. "./lift/log"
	. "./messages"
	. "./network"
	. "./network/tcp"
	. "./network/udp"
	. "fmt"
	. "net"
	. "strconv"
	"time"
	"os"
)

func main() {

	var err error

	Fo, err = os.Create("output.txt")
    if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
        if err := Fo.Close(); err != nil {
            panic(err)
        }
    }()

	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(250 * time.Millisecond))

	b := make([]byte, 16)

	_, _, err2 := ln.ReadFromUDP(b)
	ln.Close()

	slave := make(chan bool)
	master := make(chan bool)

	udp := make(chan bool)
	ip_array_update := make(chan int)
	get_ip_array := make(chan []int)
	new_master := make(chan bool)
	flush := make(chan bool)
	order := make(chan Dict)
	master_order := make(chan Dict)
	slave_order := make(chan Dict)
	queues := make(chan Queues)
	get_queues := make(chan Queues)
	do_first := make(chan Queues)
	get_at_floor := make(chan Dict)
	slave_queues := make(chan Dict)
	// algo_out := make(chan Order)

	go IP_array(ip_array_update, get_ip_array, flush)
	// Println("Starter IP_array...")
	go Timer(flush)
	go Job_queues(master_order, slave_order, get_at_floor, queues, get_queues, slave_queues, do_first)
	go Internal(order)
	go IMA(udp)
	go UDP_listen(ip_array_update)
	go Do_first(do_first)

	go func() {
		for {
			select {
			case <-master:
				Println("=> State: Entered master state")
				Fo.WriteString("=> State: Entered master state\n")
				udp <- true
				go TCP_master_connect(slave_order, slave_queues)
				go Algo(get_at_floor, get_queues)
				go func() {for {msg := <- order; master_order <- msg}}()
			case <-slave:
				Println("=> State: Entered slave state")
				Fo.WriteString("=> State: Entered slave state\n")
				udp <- false
				go IMA_master(get_ip_array, master, new_master)
				go func() { new_master <- true }()
			case <-new_master:
				Println("=> State: Entered new_master state")
				Fo.WriteString("=> State: Entered new_master state\n")
				ip := <-get_ip_array
				if len(ip) != 0 {
					if ip[len(ip)-1] > 255 {
						master_ip := ip[len(ip)-1] - 255
						go func() { new_master <- TCP_slave_com(Itoa(master_ip), order, queues) }()
						// Det som er litt unødvendig er at ny master har en TCP med seg selv..
					} else {
						go func() { new_master <- true }()
					}
				} else {
					go func() { new_master <- true }()
				}
			}
		}

	}()

	if err2 != nil { // MASTER
		Fo.WriteString("Ble master\n")
		master <- true
	} else {
		Fo.WriteString("Ble slave\n")
		slave <- true
	}

	neverQuit := make(chan string)
	<-neverQuit
}
