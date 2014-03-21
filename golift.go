package main

import (
	. "fmt"
	. "net"
	"os/exec"
	"time"
)

var state chan bool

func UDP_send() {

	saddr, _ := ResolveUDPAddr("udp", "localhost:40000")
	conn, _ := DialUDP("udp", nil, saddr)

	for {
		conn.Write([]byte("alive"))
		time.Sleep(100 * time.Millisecond)
	}
}

func UDP_listen() {

	saddr, _ := ResolveUDPAddr("udp", "localhost:40000")
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 1024)
		ln.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		_, _, err := ln.ReadFromUDP(b)
		if err != nil {
			state <- true
			break
		}
	}
}

func main() {

	Println("PROGRAM STARTED")

	var master bool
	b := make([]byte, 1024)

	go func() {
		for {
			Println("LISTENING FOR STATE")
			master = <-state
			Println("DONE LISTENING FOR STATE")
			switch {
			case master:
				go UDP_send()
				// cmd := exec.Command("mate-terminal", "-x", "go", "run", "pheonix.go")
				cmd := exec.Command("osascript", "-e", "tell", "application", "Terminal", "to", "do", "script,", "echo hello")
				cmd.Start()
			case !master:
				go UDP_listen()
			}
		}
	}()

	Println("LISTENING FOR NETWORK ACTIVITY")

	// Initiate program
	saddr, _ := ResolveUDPAddr("udp", "localhost:40000")
	ln, _ := ListenUDP("udp", saddr)
	ln.SetReadDeadline(time.Now().Add(300 * time.Millisecond))
	_, _, err := ln.ReadFromUDP(b)
	ln.Close()
	// Initiate program -- END

	Println("NETWORK ERROR:", err)

	Println("DONE LISTENING")

	if err != nil {
		state <- true
		Println("BECOME THE MASTER")
	} else {
		state <- false
		Println("BECOME THE SLAVE")
	}

	Println("PROGRAM ENDED")
}
