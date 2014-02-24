package driver

/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
#include "channels.h"
#include "elev.h"
*/
import "C"

// Wrapper for libComedi Elevator control.
// These functions provides an interface to the elevators in the real time lab
//
// 2007, Martin Korsgaard

/**
  Button types for function elev_set_button_lamp() and elev_get_button().
*/
type elev_button_type_t int

const (
	BUTTON_CALL_UP elev_button_type_t = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
)

func Init() int {
	return int(C.elev_init())
}

func Speed(speed int) {
	C.elev_set_speed(C.int(speed))
}

func Get_floor_sensor() int {
	return int(C.elev_get_floor_sensor_signal())
}

func Get_button_signal(button elev_button_type_t, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func Get_stop_signal() int {
	return int(C.elev_get_stop_signal())
}

func Get_obstruction() int {
	return int(C.elev_get_obstruction_signal())
}

func Set_floor_indicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func Set_button_lamp(button elev_button_type_t, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))
}

func Set_stop_lamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func Elev_set_door_open_lamo(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}
