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

func TCP_read(conn Conn) {

	b := make([]byte, 1024)
	for {
		conn.Read(b)
	}
}

func TCP_connect(address, port string) {

	conn, err := Dial("tcp", address+":"+port)
	_ = err

	go TCP_read(conn)
}

func IMA(address, port string, master chan bool) {

	saddr, _ := ResolveUDPAddr("udp", address+":"+port)
	conn, _ := DialUDP("udp", nil, saddr)
	var myIP string

	for {
		select {
		case state := <-master:
			if state {
				Println("Satte masterIP..!")
				myIP = "300" // master IP
			} else {
				Println("Starter GetMyIP...")
				myIP = GetMyIP()
			}
		default:
			time.Sleep(333 * time.Millisecond)
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
	Println("Sendte IP..!")
	return myIP[12:]
}

func UDP_listen(array_update chan int) {

	Println("UDP_listen startet..!")

	saddr, _ := ResolveUDPAddr("udp", ":39773")
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 16)
		_, _, err := ln.ReadFromUDP(b)
		_ = err
		remoteIP, _ := Atoi(string(b[0:3]))

		Println(remoteIP)
		array_update <- remoteIP
	}
}
