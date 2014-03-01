package functions

import (
// 	. "fmt"
)

type Dict struct {
	Ip    string
	Floor int
}

type Slice struct {
	Floor int
}

type Jobs struct {
	Ip   string
	Dest []Slice
}

type Order struct {
	Ip    string
	Pos   int
	Floor int
}

// Insert int if unique : FUNKER!
func Insert_at_pos(this []Slice, value, i int) []Slice {

	_, missing := AppendIfMissing(this, value)
	if missing {
		this = append(this[:i], append([]Slice{Slice{value}}, this[i:]...)...)
	}

	return this
}

// Pop first int : FUNKER!
func Pop_first(this []Slice) []Slice {

	return this[1:len(this)]
}

// Read first ; FUNKER!
func Read_first(this []Slice) int {

	return this[len(this)-1].Floor
}

// Remove int : FUNKER!
func Remove_from_pos(this []Slice, value int) []Slice {

	for i, floor := range this {
		if floor.Floor == value {
			this = this[:i+copy(this[i:], this[i+1:])]
		}
	}
	return this
}

// Insert at pos (ext)

func AppendIfMissing(slice []Slice, i int) ([]Slice, bool) {

	for _, ele := range slice {
		if ele.Floor == i {
			return slice, false
		}
	}
	return append(slice, Slice{i}), true
}
