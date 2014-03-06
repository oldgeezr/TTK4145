package tcp

import (
	. "../.././functions"
	. "../.././network"
	"encoding/json"
	. "fmt"
	. "net"
	// . "strconv"
	"time"
)

func TCP_master_connect(slave_order chan Dict, queues chan Queues) {

	Fo.WriteString("Entered TCP_master_connect\n")

	ln, _ := Listen("tcp", TCP_PORT)
	for {
		conn, _ := ln.Accept()
		go TCP_master_com(conn, queues)
		go func() {
			for {
				b := make([]byte, BUF_LEN)
				conn.SetReadDeadline(time.Now().Add(250 * time.Millisecond))
				length, err := conn.Read(b)
				if err != nil {
					if err.Error() == "EOF" {
						Println("closed connection")
						return
					}
				} else {
					var c Dict
					json.Unmarshal(b[0:length], &c)
					Println("From Slave:", c)
					slave_order <- c
				}
			}
		}()
	}
}

func TCP_master_com(conn Conn, queues chan Queues) {

	Fo.WriteString("Entered TCP_master_com\n")

	for {
		select {
		case msg := <-queues:
			Println("To slave:", msg)
			b, _ := json.Marshal(msg)
			conn.Write(b)
		/*case msg := <-order:
			Println("jeg passet")
			master_order <- msg*/
		}
	}
}

func TCP_slave_com(master_ip string, order chan Dict, queues chan Queues) bool {

	conn, err := Dial("tcp", IP_BASE+master_ip+TCP_PORT)
	if err != nil {
		return true
	}

	// Vi må få fjernet denne goroutinen
	go func() {
		for {
			msg := <-order
			b, _ := json.Marshal(msg)
			conn.Write(b)
		}
	}()

	for {
		b := make([]byte, BUF_LEN)
		conn.SetReadDeadline(time.Now().Add(250 * time.Millisecond))
		length, err2 := conn.Read(b)
		// Lukker connection dersom det brytes på andre siden
		if err2 != nil {
			if err2.Error() == "EOF" {
				Println("closed connection")
				return true
			}
		} else {
			var c Queues
			json.Unmarshal(b[0:length], &c)
			queues <- c
		}
	}
}
