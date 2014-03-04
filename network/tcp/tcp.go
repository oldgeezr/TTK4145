package tcp

import (
	// . "../.././lift/log"
	. "../.././functions"
	. "../.././network"
	"encoding/json"
	// . "fmt"
	. "net"
	. "strconv"
	"time"
)

func TCP_master_recieve(job_queue chan []Jobs, last_queue chan []Dict, last_floor chan Dict) {

	ln, _ := Listen("tcp", TCP_PORT)
	for {

		conn, _ := ln.Accept()

		// go TCP_master_send(conn, job_queue, last_queue)
		go TCP_master_echo(conn, last_floor)
	}
}

func TCP_master_echo(conn Conn, last_floor chan Dict) {

	for {
		b := make([]byte, BUF_LEN)
		length, _ := conn.Read(b)
		var c Dict
		json.Unmarshal(b[0:length], &c)
		if len(c.Ip_order) != 3 {
			if c.Ip_order[0] == 'X' {
				// Fikk en last order og må oppdatere last queue
				last_floor <- c
			} else {
				// Fikk en ext order og må sende til algoritme
				// Println("ext:", c)
			}
		} else {
			// Fikk int order. Må sende til algoritme
			// Println("int:", c)
		}
	}
}

func TCP_master_send(conn Conn, job_queue chan []Jobs, last_queue chan []Dict) {

	for {
		select {
		case msg := <-job_queue:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		case msg := <-last_queue:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TCP_slave_recieve(conn Conn, job_queue chan []Jobs, last_queue chan []Dict) {

	for {
		b := make([]byte, BUF_LEN)
		conn.Read(b)
		var c []Jobs
		var d []Dict
		err := json.Unmarshal(b, &c)
		if err != nil {
			_ = json.Unmarshal(b, &d)
			last_queue <- d
		} else {
			job_queue <- c
		}

	}
}

func TCP_slave_send(master_ip string, int_order, ext_order, last_floor chan Dict, job_queue chan []Jobs, last_queue chan []Dict) {

	conn, _ := Dial("tcp", IP_BASE+master_ip+TCP_PORT)
	time.Sleep(time.Second)

	go TCP_slave_recieve(conn, job_queue, last_queue)

	/*b2 := make([]byte, BUF_LEN)

	go func() {
		_, err := conn.Read(b2)
		if err != nil {
			conn.Close()
		}
	}()*/

	for {
		select {
		case msg := <-int_order:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		case msg := <-ext_order:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		case msg := <-last_floor:
			b, _ := json.Marshal(msg)
			conn.Write(b)
		default:
			time.Sleep(23 * time.Millisecond)
			// Println("default slave send")
		}
	}
}

func Connect_to_MASTER(get_array chan []int, new_master chan bool, int_order, ext_order, last_floor chan Dict, job_queue chan []Jobs, last_queue chan []Dict) {

	for {
		select {
		case <-new_master:
			time.Sleep(time.Second)
			ip := <-get_array
			if len(ip) != 0 {
				if ip[len(ip)-1] > 255 {
					master_ip := ip[len(ip)-1] - 255
					go TCP_slave_send(Itoa(master_ip), int_order, ext_order, last_floor, job_queue, last_queue)
				}
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}
