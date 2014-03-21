package network

import (
	. "net"
	"strings"
)

const (
	// UDP_PORT string = ":39777"
	// TCP_PORT string = ":27731"
	UDP_PORT string = ":39717"
	TCP_PORT string = ":27721"
	// BROADCAST string = "129.241.187.255"
	// IP_BASE   string = "129.241.187."
	BROADCAST string = "78.91.9.255"
	IP_BASE   string = "78.91.9."
	BUF_LEN   int    = 1024
	IP_LEN    int    = 8
)

func GetMyIP() string {

	allIPs, _ := InterfaceAddrs()

	IPString := make([]string, len(allIPs))
	for i := range allIPs {
		temp := allIPs[i].String()
		ip := strings.Split(temp, "/")
		IPString[i] = ip[0]
	}
	var myIP string
	for i := range IPString {
		if IPString[i][0:3] == BROADCAST[0:3] {
			myIP = IPString[i]
		}
	}

	return myIP[IP_LEN:]
}
