package udp

import (
	. "../.././functions"
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
				Println("Became MASTER..!")
				temp, _ := Atoi(GetMyIP())
				temp = temp + 255
				myIP = Itoa(temp) // master IP
			} else {
				Println("Became SLAVE..!")
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

	saddr, _ := ResolveUDPAddr("udp", UDP_PORT)
	ln, _ := ListenUDP("udp", saddr)

	for {
		b := make([]byte, 16)
		ln.ReadFromUDP(b)
		remoteIP, _ := Atoi(string(b[0:3]))
		ip_array_update <- remoteIP
	}
}

func IMA_master(get_ip_array chan []int, master, new_master, kill_IMA_master chan bool) {

	Fo.WriteString("Entered IMA_master\n")

	var count int = 0
	var count1 int = 0
	for {
		select {
		case <-kill_IMA_master:
			Fprintln(Fo, "CLOSED: Killed IMA_master")
			return
		default:
			time.Sleep(500 * time.Millisecond)
			array := <-get_ip_array
			if len(array) != 0 {
				if array[len(array)-1] < 255 {
					temp, _ := Atoi(GetMyIP())
					if temp == array[0] {
						count++
						if count == 2 { // SIKKERTHETSGRAD!
							Println("MASTER forsvant..!")
							master <- true
							time.Sleep(50 * time.Microsecond)
							return
						}
						if count1 == 2 {
							new_master <- true
						}
					} else {
						count1++
					}
				} else {
					count = 0
					count1 = 0
				}
			}
		}
	}
}
