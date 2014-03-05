package tcp

import (
	. "../.././functions"
	. "../.././network"
	"encoding/json"
	. "fmt"
	. "net"
	. "strconv"
	"time"
)

func TCP_master_connect(order chan Dict, queues chan Queues) {

	ln, _ := Listen("tcp", TCP_PORT)
	for {
		conn, _ := ln.Accept()
		go TCP_master_com(conn, order, queues)
	}
}

func TCP_master_com(conn Conn, order chan Dict, queues chan Queues) {

	for {
		b := make([]byte, BUF_LEN)
		select {
		case length, _ := conn.Read(b):
			var c Dict
			json.Unmarshal(b[0:length], &c)
			order <- c
		case msg := <-queues:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Connect_to_MASTER(get_ip_array chan []int, new_master chan bool, order chan Dict, queues chan Queues) {

	for {
		select {
		case <-new_master:
			time.Sleep(time.Second) // temp
			ip := <-get_array
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					go TCP_slave_send(Itoa(master_ip), order, queues)
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TCP_slave_send(master_ip string, order chan Dict, queues chan Queues) {

	conn, _ := Dial("tcp", IP_BASE+master_ip+TCP_PORT)

	go TCP_slave_recieve(conn, queues)

	/*b2 := make([]byte, BUF_LEN)

	go func() {
		_, err := conn.Read(b2)
		if err != nil {
			conn.Close()
		}
	}()*/

	for {
		select {
		case msg := <-order:
			Println(msg)
			b, _ := json.Marshal(msg)
			conn.Write(b)
		}
	}
}

func TCP_slave_recieve(conn Conn, queues chan Queues) {

	for {
		b := make([]byte, BUF_LEN)
		length, _ := conn.Read(b)
		var c Queues
		json.Unmarshal(b[0:length], &c)
		queues <- c
	}
}
