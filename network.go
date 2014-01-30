package main

import(
        ."fmt" 
        "time"
        ."net"
)

func send_message(conn Conn) {

        str := "Hallo!"
        _, err := conn.Write([]byte(str))
        _ = err
}

func udp_listen(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp", ":20010")        
        ln, _ := ListenUDP("udp", saddr)
        
        for {
                b2 := make([]byte,1024)
                _, _, err := ln.ReadFromUDP(b2) 
                _ = err   
                Println(string(b2))
                if err == nil {
                        time.Sleep(time.Second)
                        ch<- true
                }
        }
}

func udp_send(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp","129.241.187.255:20009")
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

func main() {

    ch := make(chan bool)
    
    go udp_listen(ch)
    go udp_send(ch)
    
    ch<- true
    
    neverQuit := make(chan string)
    <-neverQuit
}
