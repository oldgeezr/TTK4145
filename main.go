package main

import(
    ."./network"
    ."./messages"
    ."net"
    "time"
    ."fmt"
)

func main() {

    saddr, _ := ResolveUDPAddr("udp", "129.241.187.255:39773")
    ln, _ := ListenUDP("udp", saddr)
    ln.SetReadDeadline(time.Now().Add(155*time.Millisecond))
    
    b := make([]byte, 16)
    
    _, _, err := ln.ReadFromUDP(b)
    _ = err
    ln.Close()

    array_update := make(chan int)
    get_array := make(chan []int)
    flush := make(chan bool)
    master := make(chan bool)
                
    go IP_array(array_update, get_array, flush)
    go Timer(flush)

    if err != nil {
            Println("master")
            go IMA(master)
            master <- true
            go UDP_listen(array_update)
    } else {
            Println("slave")
            go IMA(master)
            master <- false
            go UDP_listen(array_update)
            go IMA_master(get_array, master)
    }

    for {
        select {

        case msg := <-get_array:
            Println(msg)
            time.Sleep(123*time.Millisecond)
        }
    }

    neverQuit := make(chan string)
    <-neverQuit    
}
