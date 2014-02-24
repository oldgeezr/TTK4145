package tcp

import (
	. "../.././network"
	. "fmt" // temp
	. "net"
	. "strconv"
	"time"
)

func TCP_Send(conn Conn, msg string) {

	time.Sleep(1100 * time.Millisecond)
	_, err := conn.Write([]byte(msg))
	_ = err
}

func TCP_read(conn Conn) {

	b := make([]byte, 1024)
	for {
		conn.Read(b)
	}
}

func TCP_echo(conn Conn) {

	b := make([]byte, 1024)
	_, err := conn.Read(b)
	Println(string(b[0:20]))
	_, err = conn.Write([]byte("Seff..!"))
	_ = err

}

func MASTER_TCP_read() {

	ln, _ := Listen("tcp", TCP_PORT)

	for {

		time.Sleep(500 * time.Millisecond)
		conn, _ := ln.Accept()

		go TCP_echo(conn)
	}
}

func TCP_connect(address, port string) {

	conn, _ := Dial("tcp", address+":"+port)

	for {

		b := make([]byte, 1024)
		_, err := conn.Read(b)
		Println(string(b[0:20]))

		_, err = conn.Write([]byte("Er du på TCP, MASTER?"))
		_ = err
	}
}

func Connect_to_MASTER(get_array chan []int, port string) {

	for {
		select {
		case ip := <-get_array:
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					TCP_connect(IP_BASE+Itoa(master_ip), port)
					break
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
