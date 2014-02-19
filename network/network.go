package network

import(
    ."fmt" // temp
    "sort"
    ."net"
    "time"
    "strings"
    ."strconv"
)

func Send(conn Conn, msg string) {
        
        _, err := conn.Write([]byte(msg))
        _ = err
}

func TCP_read(conn Conn) {
        
        b := make([]byte, 1024)
        for {
                conn.Read(b)
        }
}

func TCP_connect(address, port string) {

        conn, err := Dial("tcp", address+":"+port)
        _ = err

        go TCP_read(conn)
}

func IMA() {

        saddr, _ := ResolveUDPAddr("udp","129.241.187.255:39773")
        conn, _ := DialUDP("udp", nil, saddr)
        
        myIP := GetMyIP()
        
        for {
            	time.Sleep(5000*time.Millisecond)	
        		Send(conn, myIP)
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
        return myIP[12:15]
}

func UDP_listen() {

        saddr, _ := ResolveUDPAddr("udp", ":39773")        
        ln, _ := ListenUDP("udp", saddr)
        IPaddresses := []int{}
        
        for {
                b := make([]byte,16)
                _, _, err := ln.ReadFromUDP(b)
                _ = err
	            remoteIP, _ := Atoi(string(b[0:3]))
	            
	            IPaddresses = AppendIfMissing(IPaddresses, remoteIP)
	            sort.Ints(IPaddresses)
	            
	            Println(IPaddresses)
	            Println("")
        }
}

func AppendIfMissing(slice []int, i int) []int {

	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
