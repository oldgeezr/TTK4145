package main

import (
	. "./functions"
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
	lost_conn := make(chan bool)
	error_chan := make(chan error)

	Println("LISTENING FOR NETWORK ACTIVITY")

	go func() {
		st := <-error_chan
		if st != nil {
			Println("EVALUATE ERROR != nil")
			state <- true
			Println("BECOME THE MASTER")
		} else {
			state <- false
			Println("BECOME THE SLAVE")
		}
	}()

	// Initiate program
	saddr, _ := ResolveUDPAddr("udp", "localhost"+UDP_PORT_net+GetMyIP())
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
	_, _, err := ln.ReadFromUDP(b)
	error_chan <- err
	ln.Close()
	// Initiate program -- END

	Println("NETWORK ERROR:", err)

	Println("DONE LISTENING")

	for {
		Println("LISTENING FOR STATE")
		master = <-state
		Println("GOT STATE:", master)
		Println("DONE LISTENING FOR STATE")
		time.Sleep(100 * time.Millisecond)
		switch {
		case master:
			// --------------------------------- Start: Searching for net connection ------------------------------------
			go func() {
				for {
					connection := <-lost_conn
					if !connection {
						Println("CONNECTION ENABLED")
						return
					}
				}
			}()

			Got_net_connection(lost_conn, false)
			// --------------------------------- End: Searching for net connection --------------------------------------
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
