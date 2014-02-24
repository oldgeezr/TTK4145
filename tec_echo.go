package main

import (
	. "fmt"
	. "net"
	"time"
)

func main() {

	go TCP_listen()

	neverQuit := make(chan string)
	<-neverQuit

}

func TCP_listen() {

	ln, _ := Listen("tcp", ":27731")

	for {

		conn, _ := ln.Accept()
		go TCP_echo(conn)
	}

}

func TCP_echo(conn Conn) {

	b := make([]byte, 1024)
	conn.Read(b)
	Println(string(b))
}

func TCP_send() {

	conn, _ := Dial("tcp", "129.241.187.255:27731")
	for {
		time.Sleep(time.Second)
		b := make([]byte, 1024)
		b = []byte("yei!")
		conn.Write(b)
	}
}
