package main

import(
    ."./network"
    ."net"
    "time"
    ."fmt"
)

func main() {

    saddr, _ := ResolveUDPAddr("udp", "129.241.187.255:39773")
    ln, _ := ListenUDP("udp", saddr)
    ln.SetReadDeadline(time.Now().Add(6000*time.Millisecond))
    
    b := make([]byte, 16)
    
    _, _, err := ln.ReadFromUDP(b)
    _ = err
    ln.Close()
    
    if err != nil {
            Println("master")
            go IMA()
            go UDP_listen()
    } else {
            Println("slave")
            go IMA()
            go UDP_listen()
    }

    neverQuit := make(chan string)
    <-neverQuit    
}
