package main 

import(
    ."fmt"
    "time"
    ."net"
    ."strconv"
    "os/exec"
)

var state string
var i int

func udp_send() {

    saddr, _ := ResolveUDPAddr("udp","localhost:40000")
    conn, _ := DialUDP("udp", nil, saddr)
	msg := Itoa(i)
    
    for {
        msg = Itoa(i)
    	_, err := conn.Write([]byte(msg))
    	_ = err
    	time.Sleep(100*time.Millisecond)
    }        
}

func udp_listen() {

    saddr, _ := ResolveUDPAddr("udp", "localhost:40000")        
    ln, err := ListenUDP("udp", saddr)
    msg := 0

    for {

		ln.SetReadDeadline(time.Now().Add(300*time.Millisecond))
		if err != nil {
        	state = "master"
        	ln.Close()
        	i = 2
        	Println("start from this value: ")
        	Println(msg)
        	return
        }
        b2 := make([]byte,1024)
        _, _, err = ln.ReadFromUDP(b2)
        msg, _ = Atoi(string(b2[0])) 
    }
}

func print() {
	
	for {
		i++
		Println(i)
		time.Sleep(time.Second)
		i = i%3
	}
}

func main() {
	state = "slave"
	i = 0
	
	// Initiate program    
    saddr, _ := ResolveUDPAddr("udp", "localhost:40000")        
    ln, _ := ListenUDP("udp", saddr)
    ln.SetReadDeadline(time.Now().Add(300*time.Millisecond))
	b := make([]byte, 1024)
	_, _, err := ln.ReadFromUDP(b)
	ln.Close()

	if err != nil {
		state = "master"
	}
	// Initiate program -- END
	
	for {
		switch state {
		case "master":
			go udp_send()
			go print()
			cmd := exec.Command("mate-terminal", "-x", "go", "run", "pheonix.go")
			cmd.Start()
			Println("master")
			state = "default"
		case "slave":
			go udp_listen()
			Println("slave")
			state = "default"
		default:
            time.Sleep(100*time.Millisecond)
		}
	}
}
