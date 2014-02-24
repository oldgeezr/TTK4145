package tcp

import (
	. "../.././network"
	// . "fmt" // temp
	. "net"
	"os"
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

func TCP_echo() {

	println("Starting the server")

	listener, err := Listen("tcp", TCP_PORT)
	if err != nil {
		println("error listening:", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			println("Error accept:", err.Error())
			return
		}
		go EchoFunc(conn)
	}

}

func EchoFunc(conn Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	n, err := conn.Read(buf)
	if err != nil {
		println("Error reading:", err.Error())
		return
	}
	println("received ", n, " bytes of data =", string(buf))

	//send reply
	_, err = conn.Write(buf)
	if err != nil {
		println("Error send reply:", err.Error())
	} else {
		println("Reply sent")
	}
}

func MASTER_TCP_read() {

	ln, _ := Listen("tcp", TCP_PORT)

	for {

		time.Sleep(500 * time.Millisecond)
		conn, _ := ln.Accept()
		_ = conn

		go TCP_echo()
	}
}

func TCP_connect(address, port string) {

	conn, _ := Dial("tcp", address+port)

	for {

		b := make([]byte, 1024)
		_, err := conn.Read(b)

		_, err = conn.Write([]byte("Er du på TCP, MASTER?\x00"))
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
					go TCP_connect(IP_BASE+Itoa(master_ip), TCP_PORT)
					break
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
