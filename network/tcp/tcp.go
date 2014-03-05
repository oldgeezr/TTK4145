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

func TCP_master_connect(order, master_order chan Dict, queues chan Queues) {

	ln, _ := Listen("tcp", TCP_PORT)
	for {
		conn, _ := ln.Accept()
		go TCP_master_com(conn, order, master_order, queues)
	}
}

func TCP_master_com(conn Conn, order, master_order chan Dict, queues chan Queues) {

	for {
		select {
		case msg := <-queues:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		case msg := <-order:
			master_order <- msg
		default:
			b := make([]byte, BUF_LEN)
			conn.SetReadDeadline(time.Now().Add(50 * time.Millisecond))
			length, err := conn.Read(b)
			if err != nil {
				if err.Error() == "EOF" {
					Println("closed connection")
					return
				}
			} else {
				var c Dict
				json.Unmarshal(b[0:length], &c)
				Println(c)
				master_order <- c
			}
		}
	}
}

func TCP_slave_com(master_ip string, order chan Dict, queues chan Queues) {

	conn, _ := Dial("tcp", IP_BASE+master_ip+TCP_PORT)
	for {
		select {
		case msg := <-order:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		default:
			b := make([]byte, BUF_LEN)
			conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
			_, err := conn.Read(b)
			// Lukker connection dersom det brytes på andre siden
			if err != nil {
				if err.Error() == "EOF" {
					Println("closed connection")
					return
				}
			} else {
				// Do something
			}
		}
	}
}
