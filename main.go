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
	// . "fmt"
	. "net"
	"time"
)

func main() {

	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(250 * time.Millisecond))

	b := make([]byte, 16)

	_, _, err := ln.ReadFromUDP(b)
	ln.Close()

	ip_array_update := make(chan int)
	get_ip_array := make(chan []int)
	new_master := make(chan bool)
	flush := make(chan bool)
	master := make(chan bool)
	order := make(chan Dict)
	queues := make(chan Queues)
	// algo_out := make(chan Order)

	go IP_array(ip_array_update, get_ip_array, flush)
	// Println("Starter IP_array...")
	go Timer(flush)
	// Println("Starter Timer...")
	// go Last_queue(last_floor, get_last_queue, get_last_queue_request, new_job_queue)
	// go Job_queues(order, queues)
	// go Internal(order)
	go IMA(master)
	go UDP_listen(ip_array_update)

	if err != nil { // MASTER
		// go Master_input(int_order, ext_order, last_floor)

		// go Master_get_last_queue(get_last_queue, master_order)
		// go Master_print_last_queue(get_last_queue_request, master_request, algo_out)

		// Println("Starter IMA...")
		master <- true
		// Println("Starter UDP_listen...")
	} else { // SLAVE
		// Println("slave")
		// Println("Starter IMA...")
		master <- false
		// Println("Starter UDP_listen...")
		go IMA_master(get_ip_array, master, new_master)
		// Println("Starter IMA_master...")
		go Connect_to_MASTER(get_ip_array, new_master, order, queues)
		new_master <- true
		// go Do_first(que)
	}

	neverQuit := make(chan string)
	<-neverQuit
}
