package functions

import (
	. "fmt"
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

func Timer(flush chan bool) {

	Fo.WriteString("Entered Timer\n")

	for {
		for timer := range time.Tick(1 * time.Second) {
			_ = timer
			flush <- true
		}
		flush <- false
	}
}

func Insert_at_pos(ip string, this []Dict, value, pos int) []Dict {

	if len(this) != 0 {
		this = append(this[:pos], append([]Dict{Dict{ip, value, "int"}}, this[pos:]...)...)
	} else {
		this = []Dict{Dict{ip, value, "int"}}
		Println("ALGO: WTF?", this, pos)
	}
	return this
}

/*
func Pop_first(this []Dict) []Dict {

	return this[1:len(this)]
}

func Read_first(this []Dict) int {

	return this[len(this)-1].Floor
}


func Remove_from_pos(this []Dict, floor int) []Dict {

	for i, order := range this {
		if order.Floor == floor {
			this = this[:i+copy(this[i:], this[i+1:])]
		}
	}
	return this
}
*/

func Append_if_missing_queue(queues []Jobs, ip string) ([]Jobs, bool) {

	for _, yours := range queues {
		if yours.Ip == ip {
			return queues, false
		}
	}
	return append(queues, Jobs{ip, []Dict{}}), true
}

func Append_if_missing_floor(slice []Dict, floor int) ([]Dict, bool) {

	if len(slice) != 0 {
		for _, queue := range slice {
			if queue.Floor == floor {
				return slice, false
			}
		}
	}
	return append(slice, Dict{"ip_order", floor, "int"}), true
}

func Append_if_missing_dict(slice []Dict, last Dict) ([]Dict, bool) {

	for i, yours := range slice {
		if yours.Ip_order == last.Ip_order {
			if yours.Floor != last.Floor {
				slice[i].Ip_order = last.Ip_order
				slice[i].Floor = last.Floor
				// slice[i].Dir = last.Dir
				return slice, true
			}
			return slice, false
		}
	}
	return append(slice, last), false
}

func Update_Direction(slice []Dict, last Dict) ([]Dict, bool) {

	for i, yours := range slice {
		if yours.Ip_order == last.Ip_order {
			if yours.Floor != last.Floor {
				slice[i].Dir = last.Dir
				return slice, true
			}
			return slice, false
		}
	}
	return append(slice, last), true
}

func Append_if_missing_ext_queue(slice []Dict, floor int, dir string) ([]Dict, bool) {
	Println("AIM_ext: ", slice, floor, dir)
	for _, yours := range slice {
		if yours.Floor == floor && yours.Dir == dir {
			return slice, false
		}
	}
	return append(slice, Dict{"ext", floor, dir}), true
}

func Append_if_missing_ip(slice []int, i int) []int {

	for _, yours := range slice {
		if yours == i {
			return slice
		}
	}
	return append(slice, i)
}

func Append_if_missing_right_queue(queue []Jobs, msg Dict) []Jobs {
	for i, job := range queue {
		if job.Ip == msg.Ip_order {
			queue[i].Dest, _ = Append_if_missing_floor(queue[i].Dest, msg.Floor)
		}
	}
	return queue
}

func Remove_dict_ext_queue(this []Dict, floor int, dir string) []Dict {

	var length int = len(this)

	if length != 0 {
		for i, orders := range this {
			if orders.Floor == floor && length > 1 {
				this = this[:i+copy(this[i:], this[i+1:])]
				length = len(this)
			} else if length == 1 {
				this = []Dict{}
			}
		}
	}
	return this
}

func Remove_job_queue(this Jobs, floor int) Jobs {

	if len(this.Dest) != 0 {
		for i, orders := range this.Dest {
			if orders.Floor == floor {
				Fprintln(Fo, "Deleted from queue: ", orders)
				if len(this.Dest) != 0 {
					this.Dest = this.Dest[:i+copy(this.Dest[i:], this.Dest[i+1:])] //Kan være et problem?
				}
			}
		}
	}
	return this

}

func Someone_getting_off(job_queue Jobs, floor int) bool {

	if len(job_queue.Dest) != 0 {
		for _, orders := range job_queue.Dest {
			if orders.Floor == floor {
				return true
			}
		}
	}
	return false
}

func Someone_getting_on(job_queue []Dict, floor int, dir string) bool {
	//Print("Someone_on: ", floor, dir)
	if len(job_queue) != 0 {
		for _, orders := range job_queue {
			if orders.Floor == floor && (dir == orders.Dir || dir == "standby") {
				Println("true")
				return true
			}
		}

	}
	Println("false")
	return false
}

/*
func Determine_best_elevator(Ext_queue []Dict, Last_queue []Dict, myIP string) bool {

	var best int = 100
	var best_IP string

	for _, last := range Last_queue {
		temp := Ext_queue[0].Floor - last.Floor
		if temp < 0 {
			temp = temp * (-1)
		}
		if temp < best {
			best = temp
			best_IP = last.Ip_order
		}
	}
	if best_IP == myIP {
		return true
	} else {
		return false
	}

}
*/

/*
func Determine_dir(job_queue []Jobs, last Dict) string {
	for _, job := range job_queue {
		if last.Ip_order == job.Ip {
			if len(job.Dest) != 0 {
				if job.Dest[0].Floor-last.Floor > 0 {
					return "up"
				} else if job.Dest[0].Floor-last.Floor < 0 {
					return "down"
				} else {
					return "standby"
				}
			} else {
				return "standby"
			}
		}
	}
	return "standby"
}
*/
