package network

import (
	. "fmt" // temp
	. "net"
	. "strconv"
	"strings"
	"time"
)

func Send(conn Conn, msg string) {

	length, _ := Atoi(msg)

	if length < 100 {
		msg = "0" + msg
	}
	_, err := conn.Write([]byte(msg))
	_ = err
}

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
	_, err = conn.Write([]byte("Seff..!"))
	_ = err

}

func MASTER_TCP_read() {

	ln, _ := Listen("tcp", ":27731")

	for {

		time.Sleep(500 * time.Millisecond)
		conn, _ := ln.Accept()

		go TCP_echo(conn)
	}
}

func TCP_connect(address, port string) {

	conn, err := Dial("tcp", address+":"+port)
	Println(err)

	for {

		b := make([]byte, 1024)
		_, err := conn.Read(b)
		_ = err
	}
}

func IMA(address, port string, master chan bool, get_array chan []int) {

	saddr, _ := ResolveUDPAddr("udp", address+":"+port)
	conn, _ := DialUDP("udp", nil, saddr)
	var myIP string

	for {
		select {
		case state := <-master:
			if state {
				go MASTER_TCP_read()
				// Println("Satte masterIP..!")
				Println("Ble MASTER..!")
				temp, _ := Atoi(GetMyIP())
				temp = temp + 255
				myIP = Itoa(temp) // master IP
			} else {
				// Println("Starter GetMyIP...")
				Println("Ble SLAVE..!")
				myIP = GetMyIP()

				// Her kan vi godt gjøre oppkoblingen av TCP: Fra slave til master
				go Connect_to_MASTER(get_array, "27731")
			}
		default:
			time.Sleep(100 * time.Millisecond)
			Send(conn, myIP)
		}
	}

}

func GetMyIP() string {

	allIPs, _ := InterfaceAddrs()

	IPString := make([]string, len(allIPs))
	for i := range allIPs {
		temp := allIPs[i].String()
		ip := strings.Split(temp, "/")
		IPString[i] = ip[0]
	}
	var myIP string
	for i := range IPString {
		if IPString[i][0:3] == "192" {
			myIP = IPString[i]
		}
	}

	return myIP[10:] // HUSK OG SETTE DENNE TIL [12:] når du er på LAB !
}

func UDP_listen(array_update chan int) {

	// Println("UDP_listen startet..!")

	saddr, _ := ResolveUDPAddr("udp", ":39773")
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 16)
		_, _, err := ln.ReadFromUDP(b)
		_ = err
		remoteIP, _ := Atoi(string(b[0:3]))
		array_update <- remoteIP
	}
}

func Connect_to_MASTER(get_array chan []int, port string) {

	for {
		select {
		case ip := <-get_array:
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					TCP_connect("192.168.1."+Itoa(master_ip), port)
					break
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
