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
					master_order <- c
				}
			}
		}()
	}
}

func TCP_master_com(conn Conn, order, master_order chan Dict, queues chan Queues) {

	for {
		select {
		case msg := <-queues:
			Println("queue:", msg)
			b, _ := json.Marshal(msg)
			conn.Write(b)
			/*case msg := <-order:
			Println("order:", msg)
			master_order <- msg*/
		}
	}
}

func TCP_slave_com(master_ip string, order chan Dict, queues chan Queues) bool {

	conn, err := Dial("tcp", IP_BASE+master_ip+TCP_PORT)
	if err != nil {
		return true
	}
	for {
		select {
		case msg := <-order:
			Println(msg)
			b, _ := json.Marshal(msg)
			conn.Write(b)
		default:
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
			}
		}
	}
}
