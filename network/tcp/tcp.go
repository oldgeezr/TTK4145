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
	println("Starting the server")

	listener, err := Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		Println("error listening:", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			Println("Error accept:", err.Error())
			return
		}
		go EchoFunc(conn)
	}
}

func EchoFunc(conn Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	n, err := conn.Read(buf)
	if err != nil {
		Println("Error reading:", err.Error())
		return
	}
	Println("received ", n, " bytes of data =", string(buf))

	//send reply
	for {
		_, err = conn.Write(buf)
		if err != nil {
			Println("Error send reply:", err.Error())
		} else {
			Println("Reply sent")
		}
	}
}

func TCP_connect(address, port string) {

	conn, _ := Dial("tcp", address+port)

	for {

		b := make([]byte, 1024)
		_, err := conn.Read(b)

		_, err = conn.Write([]byte("Er du på TCP, MASTER?\n"))
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
