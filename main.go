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
)

func main() {

	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(250 * time.Millisecond))

	b := make([]byte, 16)

	_, _, err := ln.ReadFromUDP(b)
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
	// algo_out := make(chan Order)

	go IP_array(ip_array_update, get_ip_array, flush)
	// Println("Starter IP_array...")
	go Timer(flush)
	// Println("Starter Timer...")
	// go Last_queue(last_floor, get_last_queue, get_last_queue_request, new_job_queue)
	go Job_queues(master_order, slave_order, get_at_floor, queues, get_queues, do_first)
	go Internal(order)
	go IMA(udp)
	go UDP_listen(ip_array_update)
	go Do_first(do_first)

	go func() {
		for {
			select {
			case <-master:
				Println("Entered master state")
				udp <- true
				go TCP_master_connect(slave_order, queues)
				go Algo(get_at_floor, get_queues)
				go func() {for {msg := <- order; master_order <- msg}}()
			case <-slave:
				Println("Entered slave state")
				udp <- false
				go IMA_master(get_ip_array, master, new_master)
				go func() { new_master <- true }()
			case <-new_master:
				Println("Entered new_master state")
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

	if err != nil { // MASTER
		master <- true
	} else {
		slave <- true
	}

	neverQuit := make(chan string)
	<-neverQuit
}
