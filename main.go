package main

import (
	//. "./algorithm"
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
	ip_array_update := make(chan int)
	get_ip_array := make(chan []int)
	new_master := make(chan bool)
	flush := make(chan bool)
	order := make(chan Dict)
	master_order := make(chan Dict)
	queues := make(chan Queues)
	// algo_out := make(chan Order)

	go IP_array(ip_array_update, get_ip_array, flush)
	// Println("Starter IP_array...")
	go Timer(flush)
	// Println("Starter Timer...")
	// go Last_queue(last_floor, get_last_queue, get_last_queue_request, new_job_queue)
	go Job_queues(master_order, queues)
	go Internal(order)
	go IMA(master)
	go UDP_listen(ip_array_update)

	if err != nil { // MASTER
		master <- true
	} else {
		slave <- true
	}

	for {
		select {
		case <-master:
			go TCP_master_connect(order, master_order, queues)
		case <-slave:
			go IMA_master(get_ip_array, master, new_master)
		case <-new_master:
			ip <- get_ip_array
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					Println("mi master_ip:", master_ip)
					go TCP_slave_com(Itoa(master_ip), order, queues)
					slave <- true
				} else {
					new_master <- true
				}
			} else {
				new_master <- true
			}
		}
	}
}
