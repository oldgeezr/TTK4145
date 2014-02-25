package main

import (
	. "fmt"
	. "net"
	"time"
)

func main() {

	go TCP_listen()
	go TCP_send()

	neverQuit := make(chan string)
	<-neverQuit

}

func TCP_listen() {

	ln, err := Listen("tcp", ":27731")

	for {

		conn, _ := ln.Accept()
		go TCP_echo(conn)
	}

}

func TCP_echo(conn Conn) {

	for {
		b := make([]byte, 1024)
		conn.Read(b)
		Println(string(b))
	}
}

func TCP_send() {

	conn, err := Dial("tcp", "129.241.187.147:27731")
	for {
		time.Sleep(time.Second)
		b := make([]byte, 1024)
		b = []byte("yei!")
		conn.Write(b)
	}
}
