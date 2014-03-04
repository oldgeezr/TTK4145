package functions

import (
// 	. "fmt"
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

// Insert int if unique : FUNKER!
func Insert_at_pos(this []Dict, value, i int) []Dict {

	_, missing := AppendIfMissing(this, value)
	if missing {
		this = append(this[:i], append([]Dict{Dict{"IP_order", value, "dir"}}, this[i:]...)...)
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

func AppendIfMissing(slice []Dict, i int) ([]Dict, bool) {

	for _, ele := range slice {
		if ele.Floor == i {
			return slice, false
		}
	}
	return append(slice, Dict{"ip_order", i, "dir"}), true
}

func Missing_int_job(job_queue Jobs, floor int) bool {

	for _, orders := range job_queue.Dest {
		if orders.Floor == floor && orders.Ip_order == "int" {
			return false
		}
	}
	return true
}

func Missing_ext_job(job_queue []Dict, floor int, dir string) bool {

	for _, orders := range job_queue {
		if orders.Dir == dir && orders.Floor == floor {
			return false
		}
	}
	return true
}

func Remove_order_ext_queue(this []Dict, floor int, dir string) []Dict {

	for i, orders := range this {
		if orders.Dir == dir && orders.Floor == floor {
			this = this[:i+copy(this[i:], this[i+1:])]
		}
	}
	return this
}

func Remove_order_int_queue(this Jobs, floor int) Jobs {

	for i, orders := range this.Dest {
		if orders.Floor == floor {
			this.Dest = this.Dest[:i+copy(this.Dest[i:], this.Dest[i+1:])]
		}
	}
	return this

}