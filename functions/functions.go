package functions

import (
	. "fmt"
	"os"
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

type Order struct {
	Ip    string
	Pos   int
	Floor int
}

type Queues struct {
	Int_queue  []Jobs
	Ext_queue  []Dict
	Last_queue []Dict
}

var Fo *os.File

// Insert int if unique : FUNKER!
func Insert_at_pos(ip string, this []Dict, value, i int) []Dict {

	if len(this) != 0 {
		this = append(this[:i], append([]Dict{Dict{ip, value, "int"}}, this[i:]...)...)
	} else {
		this = append(this, Dict{ip, value, "int"})
	}
	return this
}

// Pop first int : FUNKER!
func Pop_first(this []Dict) []Dict {

	return this[1:len(this)]
}

// Read first ; FUNKER!
func Read_first(this []Dict) int {

	return this[len(this)-1].Floor
}

// Remove int : FUNKER!
func Remove_from_pos(this []Dict, value int) []Dict {

	for i, floor := range this {
		if floor.Floor == value {
			this = this[:i+copy(this[i:], this[i+1:])]
		}
	}
	return this
}

// Insert at pos (ext)

func AIM_Jobs(steve []Jobs, ip string) ([]Jobs, bool) {

	for _, ele := range steve {
		if ele.Ip == ip {
			return steve, false
		}
	}
	return append(steve, Jobs{ip, []Dict{}}), true
}

func AIM_Int(slice []Dict, i int) ([]Dict, bool) {

	if len(slice) != 0 {
		for _, ele := range slice {
			if ele.Floor == i {
				return slice, false
			}
		}
	}
	return append(slice, Dict{"ip_order", i, "int"}), true
}

func AIM_Dict(slice []Dict, last Dict) ([]Dict, bool) {

	for i, ele := range slice {
		if ele.Ip_order == last.Ip_order {
			if ele.Floor != last.Floor {
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

	for i, ele := range slice {
		if ele.Ip_order == last.Ip_order {
			if ele.Floor != last.Floor {
				slice[i].Dir = last.Dir
				return slice, true
			}
			return slice, false
		}
	}
	return append(slice, last), true
}

func AIM_Ext(slice []Dict, i int, G string) ([]Dict, bool) {

	for _, ele := range slice {
		if ele.Floor == i && ele.Dir == G {
			return slice, false
		}
	}
	return append(slice, Dict{"ext", i, G}), true
}

func AIM_ip(slice []int, i int) []int {

	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func Missing_int_job(job_queue Jobs, floor int) bool {

	if len(job_queue.Dest) != 0 {
		for _, orders := range job_queue.Dest {
			if orders.Floor == floor && orders.Dir == "int" {
				return false
			}
		}
	}
	return true
}

func Missing_ext_job(job_queue []Dict, floor int, dir string) bool {

	if len(job_queue) != 0 {
		for _, orders := range job_queue {
			if (orders.Dir == dir || dir == "standby") && orders.Floor == floor {
				return false
			}
		}

	}
	return true
}

func Remove_order_ext_queue(this []Dict, floor int, dir string) []Dict {

	if len(this) != 0 {
		for i, orders := range this {
			if orders.Floor == floor {
				Fprintln(Fo, "Deleted from queue: ", orders)
				this = this[:i+copy(this[i:], this[i+1:])]
			}
		}
	}
	return this
}

func Remove_order_int_queue(this Jobs, floor int) Jobs {

	if len(this.Dest) != 0 {
		for i, orders := range this.Dest {
			if orders.Floor == floor {
				Fprintln(Fo, "Deleted from queue: ", orders)
				if len(this.Dest) != 0 {
					this.Dest = this.Dest[:i+copy(this.Dest[i:], this.Dest[i+1:])]
				}
			}
		}
	}
	return this

}

func Determine_best_elevator(Ext_queue []Dict, Last_queue []Dict, myIP string) bool {

	var best int = 100
	var best_IP string
	for _, last := range Last_queue {
		temp := Ext_queue[0].Floor - last.Floor
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
