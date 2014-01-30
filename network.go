package main

import(
        ."fmt" 
        "time"
        ."net"
)

func send_message(conn Conn) {

        str := "Hei!"
        _, err := conn.Write([]byte(str))
        _ = err
}

func udp_listen(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp", ":20009")        
        ln, _ := ListenUDP("udp", saddr)
        
        for {
                b2 := make([]byte,1024)
                _, _, err := ln.ReadFromUDP(b2)
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
                }
        }        
}

func main() {
    
    go udp_listen()
    go udp_send()
}
