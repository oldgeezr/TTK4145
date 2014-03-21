package main

import (
	. "./goelevator"
	. "./network"
	. "./network/udp"
	. "fmt"
	. "net"
	"os/exec"
	"time"
)

func main() {

	Println("PROGRAM STARTED")

	var master bool
	state := make(chan bool)
	b := make([]byte, 1024)

	Println("LISTENING FOR NETWORK ACTIVITY")

	// Initiate program
	saddr, _ := ResolveUDPAddr("udp", "localhost"+UDP_PORT_net)
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(300 * time.Millisecond))
	_, _, err := ln.ReadFromUDP(b)
	ln.Close()
	// Initiate program -- END

	Println("NETWORK ERROR:", err)

	Println("DONE LISTENING")

	go func() {
		if err != nil {
			Println("EVALUATE ERROR != nil")
			state <- true
			Println("BECOME THE MASTER")
		} else {
			state <- false
			Println("BECOME THE SLAVE")
		}
	}()

	for {
		Println("LISTENING FOR STATE")
		master = <-state
		Println("DONE LISTENING FOR STATE")
		switch {
		case master:
			Println("STAGE 1")
			go UDP_send_clone()
			cmd := exec.Command("mate-terminal", "-x", "go", "run", "main.go")
			cmd.Start()
			Go_elevator()
			Println("STAGE 2")
			return
		case !master:
			Println("STAGE 3")
			go UDP_listen_clone(state)
		}
	}

	Println("PROGRAM ENDED")
}
