package main

import (
	. "fmt"
	. "net"
	"time"
)

func main() {

	kill_net := make(chan bool)

	go Got_net_connection(kill_net)

	for {
		select {
		case msg := <-kill_net:
			Println("KILL:", msg)
		}
	}

}

func Got_net_connection(kill_net chan bool) {
	var alive bool = true
	saddr, _ := ResolveUDPAddr("udp", "www.google.com:http")
	for {

		conn, err := DialUDP("udp", nil, saddr)
		time.Sleep(50 * time.Millisecond)

		switch {
		case err == nil && alive:
			Println("GOT NO ERROR")
			time.Sleep(50 * time.Millisecond)
			conn.Close()
		case err != nil && alive:
			kill_net <- true
			alive = false
			Println("GOT ERROR, HAVE NOT SENDT STATE")
			Println("ERROR:", err)
		case err != nil && !alive:
			Println("GOT ERROR")
			time.Sleep(50 * time.Millisecond)
		case err == nil && !alive:
			kill_net <- false
			alive = true
			Println("GOT NO ERROR, HAVE NOT SENDT STATE")
			Println("ERROR:", err)
		}
	}
}
