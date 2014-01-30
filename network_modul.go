package main

import(
        ."fmt" 
        "time"
        ."net"
)

func read_from_input() (string) {
        
        var str string
        
        Println("Type some shit:")
        Scanln(&str)
        str = str+"\x00"
        
        return str
}

func send_message(conn Conn) {

        str := read_from_input()
        _, err := conn.Write([]byte(str))
        _ = err
}

func udp_listen(ch chan bool) {

        saddr, _ := ResolveUDPAddr("udp", ":20006")        
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

        saddr, _ := ResolveUDPAddr("udp","129.241.187.255:20006")
        conn, _ := DialUDP("udp", nil, saddr)
        
        for {
                if <-ch == true {
                        send_message(conn)
                }
        }        
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

func main() {

        
        ch := make(chan bool)

        go udp_send(ch)
        go udp_listen(ch)
        
        ch<- true
        
        
        /*
        go tcp_connect("129.241.187.161", "34933")
        */
        
        neverQuit := make(chan string)
        <-neverQuit
}
