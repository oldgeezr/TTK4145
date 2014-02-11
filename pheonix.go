package main 

import(
    ."fmt"
    "time"
    ."net"
)

func udp_send() {

    saddr, _ := ResolveUDPAddr("udp","localhost:40000")
    conn, _ := DialUDP("udp", nil, saddr)
    
    for {
        
        _, err := conn.Write([]byte("lol"))
        _ = err
        time.Sleep(100*time.Millisecond)
    }        
}

func udp_listen(ch chan string) {

    saddr, _ := ResolveUDPAddr("udp", "localhost:40000")        
    ln, _ := ListenUDP("udp", saddr)

    for {

        b := make([]byte,1024)
        _, _, err := ln.ReadFromUDP(b)
        _ = err
    }
}

func main() {
    
    /* i := 0
    for {
        Println(i)
        i++
        i = i%4
        time.Sleep(500*time.Millisecond)
    }*/
    
    ch := make(chan string)
    
    go udp_listen(ch)
    
    if timer := time.Tick(250*time.Millisecond) {
        Println("ffsadf")
    }

    neverQuit := make(chan string)
    <-neverQuit
}
