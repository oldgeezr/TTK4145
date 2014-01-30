package main

import(
        ."fmt" 
        "time"
        ."net"
        "strings"
)

func send_message(conn Conn) {
        
        str := GetMyIP() + "Hei!"
        _, err := conn.Write([]byte(str))
        _ = err
}

func udp_listen(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp", ":10020")        
        ln, _ := ListenUDP("udp", saddr)
        
        for {
                b2 := make([]byte,1024)
                _, _, err := ln.ReadFromUDP(b2) 
                _ = err
                if err == nil {
                        time.Sleep(time.Second)
                        ch<- true
                }
                
                if string(b2[0:15]) != GetMyIP() {
                    Println("UDP from " + string(b2[0:15]) + " saying " + string(b2[15:200]))
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
        allIPs, err := InterfaceAddrs()
        if err != nil {
                Println("network.GetMyIP()--> Error receiving IPs")
                return ""
        }
        
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

func main() {

    ch := make(chan bool)
    
    go udp_listen(ch)
    go udp_send(ch)
    
    ch<- true
    
    neverQuit := make(chan string)
    <-neverQuit
}
