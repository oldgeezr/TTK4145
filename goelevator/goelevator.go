package main

import (
	. ".././network"
	. "fmt"
	. "net"
	"os/exec"
	"time"
)

func UDP_send() {

	saddr, _ := ResolveUDPAddr("udp", "localhost"+UDP_PORT_net)
	conn, _ := DialUDP("udp", nil, saddr)

	for {
		conn.Write([]byte("alive"))
		time.Sleep(50 * time.Millisecond)
	}
}

func UDP_listen(state chan bool) {

	saddr, _ := ResolveUDPAddr("udp", "localhost"+UDP_PORT_net)
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 1024)
		ln.SetReadDeadline(time.Now().Add(150 * time.Millisecond))
		_, _, err := ln.ReadFromUDP(b)
		if err != nil {
			Println("MASTER DIED")
			state <- true
			Println("INITIATE NEW MASTER")
			break
		}
	}
}

func main() {

	Println("PROGRAM STARTED")

	var master bool
	state := make(chan bool)
	b := make([]byte, 1024)

	go func() {
		for {
			Println("LISTENING FOR STATE")
			master = <-state
			Println("DONE LISTENING FOR STATE")
			switch {
			case master:
				Println("STAGE 1")
				go UDP_send()
				cmd := exec.Command("mate-terminal", "-x", "go", "run", "golift.go")
				cmd2 := exec.Command("mate-terminal", "-x", "go", "run", "../main.go")
				// cmd := exec.Command("osascript", "-e", "tell", "application", "Terminal", "to", "do", "script,", "echo hello")
				cmd.Start()
				cmd2.Start()
				Println("STAGE 2")
			case !master:
				Println("STAGE 3")
				go UDP_listen(state)
			}
		}
	}()

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

	if err != nil {
		Println("EVALUATE ERROR != nil")
		state <- true
		Println("BECOME THE MASTER")
	} else {
		state <- false
		Println("BECOME THE SLAVE")
	}

	Println("PROGRAM ENDED")

	neverQuit := make(chan string)
	<-neverQuit
}
