package tcp

import (
	. "../.././lift"
	. "../.././network"
	. "fmt"
	. "net"
	. "strconv"
	"time"
)

func TCP_listen() {

	ln, _ := Listen("tcp", TCP_PORT)
	for {

		conn, _ := ln.Accept()
		go TCP_echo(conn)
	}
}

func TCP_echo(conn Conn) {

	for {
		b := make([]byte, BUF_LEN)
		conn.Read(b)
		Println(string(b))
		conn.Write(b)
	}
}

func TCP_slave(conn Conn) {

	for {
		b := make([]byte, BUF_LEN)
		conn.Read(b)
		Println(string(b))
		msg, _ := Atoi(string(b[0]))
		Println(msg)
		Send_to_floor(msg)
	}
}

func TCP_connect(master_ip string, int_order, ext_order chan string) {

	conn, _ := Dial("tcp", IP_BASE+master_ip+TCP_PORT)
	go TCP_slave(conn)
	time.Sleep(time.Second)
	for {
		b := make([]byte, BUF_LEN)
		select {
		case msg := <-int_order:
			b = []byte(msg)
		case msg := <-ext_order:
			b = []byte(msg)
		}
		conn.Write(b)
	}
}

func Connect_to_MASTER(get_array chan []int, port string, new_master chan bool, int_order, ext_order chan string) {

	for {
		select {
		case <-new_master:
			time.Sleep(time.Second)
			ip := <-get_array
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					go TCP_connect(Itoa(master_ip), int_order, ext_order)
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
