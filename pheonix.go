package main 

import(
    ."fmt"
    "time"
    ."net"
    ."strconv"
    "os/exec"
)

var state string = "default"

func udp_send(ch <-chan int) {

    saddr, _ := ResolveUDPAddr("udp","localhost:40000")
    conn, _ := DialUDP("udp", nil, saddr)
    temp := <-ch
	msg := Itoa(temp)
	first := true
    
    for {
        select {
        case temp = <-ch:
        	msg = Itoa(temp)
	    default:
	    	if first {
	        	msg = msg + "!"
	        	first = false
	        }
    	}
    	_, err := conn.Write([]byte(msg))
    	_ = err
    	Println(msg)
    	time.Sleep(100*time.Millisecond)
    }        
}

func udp_listen() {

    saddr, _ := ResolveUDPAddr("udp", "localhost:40000")        
    ln, _ := ListenUDP("udp", saddr)

    for {

		ln.SetReadDeadline(time.Now().Add(500*time.Millisecond))
        b := make([]byte,1024)
        _, _, err := ln.ReadFromUDP(b)
        if err != nil {
        	state = "master"
        	break
        }
    }
}

func print(ch chan int) {
	
	i := 0
	for {
		i++
		Println(i)
		ch <- i
		time.Sleep(time.Second)
		i = i%3
	}
}

func main() {

	ch := make(chan int)
	
	// Initiate program    
    saddr, _ := ResolveUDPAddr("udp", "localhost:40000")        
    ln, _ := ListenUDP("udp", saddr)
    ln.SetReadDeadline(time.Now().Add(300*time.Millisecond))
    
	b := make([]byte, 1024)
	_, _, err := ln.ReadFromUDP(b)
	ln.Close()

	if err != nil {
		state = "master"
	} else {
		state = "slave"
	}
	// Initiate program -- END
	
	for {
		switch state {
		case "master":
			go udp_send(ch)
			go print(ch)
			cmd := exec.Command("mate-terminal", "-x", "go", "run", "pheonix.go")
			cmd.Start()
			Println("master")
			state = "default"
		case "slave":
			go udp_listen()
			Println("slave")
			state = "default"
		case "default":

		}
	}
}
