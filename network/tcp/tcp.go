package tcp

import (
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
	}
}

func TCP_connect(master_ip string, int_order chan string) {

	conn, _ := Dial("tcp", IP_BASE+master_ip+TCP_PORT)
	for {
		time.Sleep(time.Second)
		b := make([]byte, BUF_LEN)
		msg := <-int_order
		b = []byte(msg)
		conn.Write(b)
	}
}

func Connect_to_MASTER(get_array chan []int, port string, new_master chan bool) {

	for {
		select {
		case <-new_master:
			time.Sleep(time.Second)
			ip := <-get_array
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					go TCP_connect(Itoa(master_ip))
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
