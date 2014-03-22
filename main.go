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

	var master bool
	state := make(chan bool)
	b := make([]byte, 1024)
	lost_conn := make(chan bool)
	error_chan := make(chan error)

	// --------------------------------- Start: Searching for net connection --------------------------------------------
	go func() {
		st := <-error_chan
		if st != nil {
			state <- true
		} else {
			state <- false
		}
	}()
	// --------------------------------- End: Searching for net connection ----------------------------------------------

	// --------------------------------- Start: Listen for network activity ---------------------------------------------
	saddr, _ := ResolveUDPAddr("udp", "localhost"+UDP_PORT_net+GetMyIP())
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
	_, _, err := ln.ReadFromUDP(b)
	error_chan <- err
	ln.Close()
	// --------------------------------- End: Listen for network activity -----------------------------------------------

	for {
		master = <-state
		time.Sleep(100 * time.Millisecond)

		switch {
		case master:
			// --------------------------------- Start: Searching for net connection ------------------------------------
			go func() {
				for {
					connection := <-lost_conn
					if !connection {
						return
					}
				}
			}()

			Got_net_connection(lost_conn, false)
			// --------------------------------- End: Searching for net connection --------------------------------------
			go UDP_send_clone()
			cmd := exec.Command("mate-terminal", "-x", "go", "run", "main.go")
			cmd.Start()
			Go_elevator()
			return
		case !master:
			go UDP_listen_clone(state)
			Println("WATCHDOG")
		}
	}
}
