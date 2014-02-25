package tcp

import (
	. "../.././network"
	. "fmt" // temp
	. "net"
	"os"
	. "strconv"
	"time"
)

func TCP_listen() {

	ln, err := Listen("tcp", TCP_PORT)
	for {

		conn, _ := ln.Accept()
		go TCP_echo(conn)
	}

}

func TCP_echo(conn Conn) {

	for {
		b := make([]byte, 1024)
		conn.Read(b)
		Println(string(b))
	}
}

func TCP_connect(master_ip string) {

	conn, err := Dial("tcp", IP_BASE+master_ip+TCP_PORT)
	for {
		time.Sleep(time.Second)
		b := make([]byte, 1024)
		b = []byte("yei!")
		conn.Write(b)
	}
}

func Connect_to_MASTER(get_array chan []int, port string) {

	for {
		select {
		case ip := <-get_array:
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					go TCP_connect(Itoa(master_ip))
					break
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
