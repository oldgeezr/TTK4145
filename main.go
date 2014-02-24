package main

import (
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
	flush := make(chan bool)
	master := make(chan bool)

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
		// Println("slave")
		go IMA(BROADCAST, UDP_PORT, master, get_array)
		// Println("Starter IMA...")
		master <- false
		go UDP_listen(array_update)
		// Println("Starter UDP_listen...")
		go IMA_master(get_array, master)
		// Println("Starter IMA_master...")
		go Connect_to_MASTER(get_array, UDP_PORT)
	}

	/*for {
		select {

		case msg := <-get_array:
			Println(msg)
			time.Sleep(150 * time.Millisecond)
		}
	}*/

	neverQuit := make(chan string)
	<-neverQuit
}
