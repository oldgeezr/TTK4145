package udp

import (
	. "../.././network"
	. "../.././functions"
	. "fmt"
	. "net"
	. "strconv"
	"time"
)

func UDP_send(conn Conn, msg string) {

	Fo.WriteString("Entered UDP_Send\n")

	length, _ := Atoi(msg)
	if length < 100 {
		msg = "0" + msg
	}
	conn.Write([]byte(msg))
}

func IMA(udp chan bool) {

	Fo.WriteString("Entered IMA\n")

	saddr, _ := ResolveUDPAddr("udp", BROADCAST+UDP_PORT)
	conn, _ := DialUDP("udp", nil, saddr)
	var myIP string

	for {
		select {
		case state := <-udp:
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

	Fo.WriteString("Entered UDP_listen\n")

	// Println("UDP_listen startet..!")
	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 16)
		ln.ReadFromUDP(b)
		remoteIP, _ := Atoi(string(b[0:3]))
		ip_array_update <- remoteIP
	}
}
