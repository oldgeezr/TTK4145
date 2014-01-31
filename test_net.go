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

func udp_listen() {

        saddr, _ := ResolveUDPAddr("udp", ":10020")        
        ln, _ := ListenUDP("udp", saddr)

        for {
                b := make([]byte,16)
                _, _, err := ln.ReadFromUDP(b)
		remoteIP := string(b[0:15]) 
                if err == nil {
                        time.Sleep(50*time.Millisecond)
                }
               
                if remoteIP != GetMyIP() {
                	Println(remoteIP)
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

	go udp_listen()
	go udp_send()

	ch<- true
}

func main() {

	go network_modul()
	
	neverQuit := make(chan string)
	<-neverQuit
}
