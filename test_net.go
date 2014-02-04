package main 

import(
	."fmt"
	."net"
	"time"
	"strings"
)

func send_message(conn Conn) {
        
        str := GetMyIP()
        _, err := conn.Write([]byte(str))
        _ = err
}

func tcp_connect(address, port string, ch chan bool) {

	Println("but hei222")

        conn, err := Dial("tcp", address+":"+port)
	if err != nil {
		Println("but hei")
		ch<- false
	}        
		
        go tcp_read(conn)
        go send_message(conn)

	ch<- true
}

func tcp_read(conn Conn) {
        
        b := make([]byte, 1024)
        for {
                conn.Read(b)
                Println(string(b))
        }
}

func udp_listen(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp", ":10020")        
        ln, _ := ListenUDP("udp", saddr)
	IPaddrses := make([]string,2)

        for {
                b := make([]byte,16)
                _, _, err := ln.ReadFromUDP(b)
		remoteIP := string(b[0:15]) 
                if err == nil {
                        time.Sleep(50*time.Millisecond)
                }
                
                if remoteIP != GetMyIP() {
			if IPaddrses[0] == "" {
				IPaddrses[0] = remoteIP
				go tcp_connect(remoteIP, "10021", ch)
			} else {
				for i := 1; i < len(IPaddrses); i++ {
					if remoteIP != IPaddrses[i-1] && IPaddrses[i] == "" {
						IPaddrses[i] = remoteIP
						// go tcp_connect(remoteIP, "10021", ch)
						break
					}
				}
			}
			// Brukes til å sjekke om man mottar IP
                	// Println(remoteIP)
                }
        }
}

func udp_send() {

        saddr, _ := ResolveUDPAddr("udp","129.241.187.255:10020")
        conn, _ := DialUDP("udp", nil, saddr)
        
        for {
		send_message(conn)
            	time.Sleep(100*time.Millisecond)		
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
        for i:=range IPString {
                if IPString[i][0:3] == "129" {
                        myIP = IPString[i]
                }
        }
        return myIP
}

func network_modul() {

	ch := make(chan bool)	

	go udp_listen(ch)
	go udp_send()

}

func main() {

	go network_modul()
	
	neverQuit := make(chan string)
	<-neverQuit
}
