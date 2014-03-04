package udp

import (
	// . "../.././lift/log"
	. "../.././functions"
	. "../.././network"
	. ".././tcp"
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

func IMA(master chan bool, get_array chan []int, job_queue chan []Jobs, last_queue chan []Dict, last_floor chan Dict) {

	saddr, _ := ResolveUDPAddr("udp", BROADCAST+UDP_PORT)
	conn, _ := DialUDP("udp", nil, saddr)
	var myIP string

	for {
		select {
		case state := <-master:
			if state {
				go TCP_master_recieve(job_queue, last_queue, last_floor)
				// Println("Satte masterIP..!")
				Println("Ble MASTER..!")
				temp, _ := Atoi(GetMyIP())
				temp = temp + 255
				myIP = Itoa(temp) // master IP
			} else {
				// go TCP_listen(false)
				// Println("Starter GetMyIP...")
				Println("Ble SLAVE..!")
				myIP = GetMyIP()
			}
		default:
			time.Sleep(100 * time.Millisecond)
			UDP_send(conn, myIP)
		}
	}

}

func UDP_listen(array_update chan int) {

	// Println("UDP_listen startet..!")
	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 16)
		_, _, err := ln.ReadFromUDP(b)
		_ = err
		remoteIP, _ := Atoi(string(b[0:3]))
		array_update <- remoteIP
	}
}
