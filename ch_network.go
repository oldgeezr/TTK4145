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

func udp_listen(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp", ":10020")        
        ln, _ := ListenUDP("udp", saddr)
		tcp := false 
        
        for {
                b2 := make([]byte,1024)
                _, raddr, err := ln.ReadFromUDP(b2) 
                if err == nil {
                        time.Sleep(200*time.Millisecond)
                        ch<- true
                }
                
                if string(b2[0:15]) != GetMyIP() {
                    Println(string(b2[0:15]))
					if string(b2[0:15]) == string(raddr.IP) && tcp == false {
						go tcp_connect(string(b2[0:15]), "10020")
						tcp = true
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

func tcp_connect(address, port string) {

        conn, _ := Dial("tcp", address+":"+port)
        
        go tcp_read(conn)
        go send_message(conn)
}

func network_modul() {	
	
	ch := make(chan bool)

	go udp_listen(ch)
	go udp_send(ch)

	ch<- true
}

func main() {

    go network_modul()
    
    neverQuit := make(chan string)
    <-neverQuit
}
