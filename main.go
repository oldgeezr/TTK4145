package main

import (
	. "./lift"
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

	array_update := make(chan int)
	get_array := make(chan []int)
	new_master := make(chan bool)
	flush := make(chan bool)
	master := make(chan bool)

	int_button := make(chan int)
	int_order := make(chan string)
	ext_order := make(chan string)

	go IP_array(array_update, get_array, flush)
	// Println("Starter IP_array...")
	go Timer(flush)
	// Println("Starter Timer...")

	if err != nil { // MASTER
		go IMA(BROADCAST, UDP_PORT, master, get_array)
		// Println("Starter IMA...")
		master <- true
		go UDP_listen(array_update)
		// Println("Starter UDP_listen...")
	} else { // SLAVE
		go Internal(int_button, int_order, ext_order)
		// Println("slave")
		go IMA(BROADCAST, UDP_PORT, master, get_array)
		// Println("Starter IMA...")
		master <- false
		go UDP_listen(array_update)
		// Println("Starter UDP_listen...")
		go IMA_master(get_array, master, new_master)
		// Println("Starter IMA_master...")
		go Connect_to_MASTER(get_array, UDP_PORT, new_master, int_order, ext_order)
		new_master <- true
	}

	neverQuit := make(chan string)
	<-neverQuit
}
