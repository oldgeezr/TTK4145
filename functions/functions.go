package functions

import (
	. ".././network"
	. "net"
	"os"
	"time"
)

type Dict struct {
	Ip_order string
	Floor    int
	Dir      string
}

type Jobs struct {
	Ip   string
	Dest []Dict
}

type Queues struct {
	Job_queue  []Jobs
	Ext_queue  []Dict
	Last_queue []Dict
}

var Fo *os.File

func Got_net_connection(lost_conn chan bool, alive bool) {

	for {
		saddr, _ := ResolveUDPAddr("udp", "www.google.com:http")
		conn, err := DialUDP("udp", nil, saddr)
		time.Sleep(50 * time.Millisecond)

		switch {
		case err == nil && alive:
			time.Sleep(50 * time.Millisecond)
			conn.Close()

		case err != nil && alive:
			lost_conn <- true
			alive = false

		case err != nil && !alive:
			time.Sleep(50 * time.Millisecond)

		case err == nil && !alive:
			lost_conn <- false
			alive = true
			return
		}
	}
}

func Flush_IP_array(flush chan bool) {

	Fo.WriteString("Entered Timer\n")

	for {
		for timer := range time.Tick(1 * time.Second) {
			_ = timer
			flush <- true
		}
		flush <- false
	}
}

func Create_job_queue_if_missing(queues []Jobs, ip string) []Jobs {

	for _, yours := range queues {
		if yours.Ip == ip {
			return queues
		}
	}
	return append(queues, Jobs{ip, []Dict{}})
}

func Update_last_queue(slice []Dict, last Dict, dir bool) []Dict {

	for i, yours := range slice {
		if yours.Ip_order == last.Ip_order {
			if yours.Floor != last.Floor {
				if dir {
					slice[i].Dir = last.Dir
				} else {
					slice[i].Ip_order = last.Ip_order
					slice[i].Floor = last.Floor
				}
				return slice
			}
			return slice
		}
	}
	return append(slice, last)
}

func Append_if_missing_ext_queue(slice []Dict, floor int, dir string) []Dict {

	for _, yours := range slice {
		if yours.Floor == floor && yours.Dir == dir {
			return slice
		}
	}
	return append(slice, Dict{"ext", floor, dir})
}

func Append_if_missing_ip(slice []int, i int) []int {

	for _, yours := range slice {
		if yours == i {
			return slice
		}
	}
	return append(slice, i)
}

func Append_to_correct_queue(queue []Jobs, msg Dict) []Jobs {

	for i, job := range queue {
		if job.Ip == msg.Ip_order {
			queue[i].Dest = Append_if_missing_order(queue[i].Dest, msg.Floor)
		}
	}
	return queue
}

func Append_if_missing_order(slice []Dict, floor int) []Dict {

	if len(slice) != 0 {
		for _, queue := range slice {
			if queue.Floor == floor {
				return slice
			}
		}
	}
	return append(slice, Dict{"ip_order", floor, "int"})
}

func Remove_from_ext_queue(this []Dict, floor int, dir string) []Dict {

	var length int = len(this)

	if length != 0 {
		for i, orders := range this {
			if orders.Floor == floor && (dir == orders.Dir || dir == "standby") {
				if length > 1 {
					this = this[:i+copy(this[i:], this[i+1:])]
					break
				} else if length == 1 {
					this = []Dict{}
					break
				}
			}
		}
	}

	return this
}

func Remove_job_queue(this Jobs, floor int) Jobs {

	var length int = len(this.Dest)

	if length != 0 {
		for i, orders := range this.Dest {
			if orders.Floor == floor {
				if length > 1 {
					this.Dest = this.Dest[:i+copy(this.Dest[i:], this.Dest[i+1:])] //Kan være et problem?
				} else if length == 1 {
					this.Dest = []Dict{}
				}
			}
		}
	}

	return this
}

func Insert_at_pos(ip string, this []Dict, value, pos int) []Dict {

	// DO THIS ORDER APPEAR IN THE JOB_QUEUE. APPEND IF NOT
	if !Someone_getting_off(this, value) {
		this = append(this[:pos], append([]Dict{Dict{ip, value, "int"}}, this[pos:]...)...)
	} else {
		if len(this) == 0 {
			this = []Dict{Dict{ip, value, "int"}}
		}
	}
	return this
}

func Someone_getting_off(job_queue []Dict, floor int) bool {

	if len(job_queue) != 0 {
		for _, orders := range job_queue {
			if orders.Floor == floor {
				return true
			}
		}
	}
	return false
}

func Someone_getting_on(job_queue []Dict, floor int, dir string) bool {

	if len(job_queue) != 0 {
		for _, orders := range job_queue {
			if orders.Floor == floor && (dir == orders.Dir || dir == "standby") && (orders.Ip_order == "ext" || orders.Ip_order == GetMyIP()) {
				return true
			}
		}

	}

	return false
}
