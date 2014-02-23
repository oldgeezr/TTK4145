package network

import (
	. "net"
	"strings"
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
		if IPString[i][0:3] == "192" {
			myIP = IPString[i]
		}
	}

	return myIP[10:] // HUSK OG SETTE DENNE TIL [12:] når du er på LAB !
}
