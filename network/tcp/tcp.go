package tcp

import (
	. "../.././functions"
	. "../.././network"
	"encoding/json"
	. "fmt"
	. "net"
	"time"
)

func TCP_master_connect(log_order chan Dict, queues_to_tcp chan Queues, get_ip_array chan []int) {

	Fo.WriteString("Entered TCP_master_connect\n")

	ln, _ := Listen("tcp", TCP_PORT)

	for {
		conn, _ := ln.Accept()

		go TCP_master_com(conn, queues_to_tcp)

		go func() {
			for {
				b := make([]byte, BUF_LEN)
				conn.SetReadDeadline(time.Now().Add(250 * time.Millisecond))
				length, err := conn.Read(b)

				if err != nil {
					if err.Error() == "EOF" || Ping_PC(get_ip_array, conn.RemoteAddr()) {
						Println("CLOSED CONNECTION")
						return
					}
				} else {
					var c Dict
					json.Unmarshal(b[0:length], &c)
					log_order <- c
				}
			}
		}()
	}
}

func TCP_master_com(conn Conn, queues_to_tcp chan Queues) {

	Fo.WriteString("Entered TCP_master_com\n")

	for {
		msg := <-queues_to_tcp
		b, _ := json.Marshal(msg)
		conn.Write(b)
		time.Sleep(100 * time.Millisecond)
	}
}

func TCP_slave_com(master_ip string, order chan Dict, slave_queues chan Queues) bool {

	conn, err := Dial("tcp", IP_BASE+master_ip+TCP_PORT)

	if err != nil {
		return true
	}

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

		Println("TCP_SLAVE IS STILL RUNNGING")

		if err2 != nil {
			if err2.Error() == "EOF" || Ping_PC(get_ip_array, conn.RemoteAddr()) {
				Println("CLOSED CONNECTION")
				return true
			}
		} else {
			var c Queues
			json.Unmarshal(b[0:length], &c)
			slave_queues <- c
		}
	}
}
