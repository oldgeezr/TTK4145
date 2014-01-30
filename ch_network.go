package main

import(
        ."fmt" 
        "time"
        ."net"
        "strings"
)

func send_message(conn Conn) {
        
        str := GetMyIP()
        _, err := conn.Write([]byte(str))
        _ = err
}

func udp_listen(ch, ch2 chan bool) {

        saddr, _ := ResolveUDPAddr("udp", ":10020")        
        ln, _ := ListenUDP("udp", saddr)
		tcp := false 
        
        for {
                b2 := make([]byte,1024)
                _, _, err := ln.ReadFromUDP(b2)
				remoteIP := string(b2[0:15]) 
                if err == nil {
                        time.Sleep(200*time.Millisecond)
                        ch<- true
                }
               
                if remoteIP != GetMyIP() {
                    Println(remoteIP)
					if tcp != true {
						go tcp_connect(remoteIP, "10020", ch2)
						time.Sleep(500*time.Millisecond)
						Println("I was here")
						tcp = <-ch2
					}
                }
        }
}

func udp_send(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp","129.241.187.255:10020")
        conn, _ := DialUDP("udp", nil, saddr)
        
        for {
                if <-ch == true {
                    send_message(conn)
                } else {
                    time.Sleep(time.Second)
                    send_message(conn)
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
        for i:=range IPString {
                if IPString[i][0:3] == "129" {
                        myIP = IPString[i]
                }
        }
        return myIP
}

func tcp_read(conn Conn) {
        
        b := make([]byte, 1024)
        for {
                conn.Read(b)
                Println(string(b))
        }
}

func tcp_connect(address, port string, ch2 chan bool) {

        conn, err := Dial("tcp", address+":"+port)
		if err != nil {
			ch2<- false
		}        
		
        go tcp_read(conn)
        go send_message(conn)

		ch2<- true
}

func network_modul() {	
	
	ch := make(chan bool)
	ch2 := make(chan bool)

	go udp_listen(ch, ch2)
	go udp_send(ch)

	ch<- true
	ch2<- false
}

func main() {

    go network_modul()
    
    neverQuit := make(chan string)
    <-neverQuit
}
