package udp

import (
	. "../.././network"
	. "fmt"
	. "net"
	. "strconv"
	"time"
)

func UDP_send(conn Conn, msg string) {

	length, _ := Atoi(msg)
	if length < 100 {
		msg = "0" + msg
	}
	_, err := conn.Write([]byte(msg))
	_ = err
}

func IMA(master chan bool) {

	saddr, _ := ResolveUDPAddr("udp", BROADCAST+UDP_PORT)
	conn, _ := DialUDP("udp", nil, saddr)
	var myIP string

	for {
		select {
		case state := <-master:
			if state {
				Println("Ble MASTER..!")
				temp, _ := Atoi(GetMyIP())
				temp = temp + 255
				myIP = Itoa(temp) // master IP
			} else {
				Println("Ble SLAVE..!")
				myIP = GetMyIP()
			}
		default:
			time.Sleep(100 * time.Millisecond)
			UDP_send(conn, myIP)
		}
	}
}

func UDP_listen(ip_array_update chan int) {

	// Println("UDP_listen startet..!")
	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 16)
		_, _, err := ln.ReadFromUDP(b)
		_ = err
		remoteIP, _ := Atoi(string(b[0:3]))
		ip_array_update <- remoteIP
	}
}
