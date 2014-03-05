package messages

import (
	. ".././functions"
	. ".././network"
	. "fmt"
	"sort"
	. "strconv"
	"time"
)

func IP_array(ip_array_update chan int, get_ip_array chan []int, flush chan bool) {

	IPaddresses := []int{}
	// Println("IP_array startet..!")
	for {
		select {
		case ip := <-ip_array_update:
			// Println("Oppdaterte arrayet..!")
			IPaddresses = AIM_ip(IPaddresses, ip)
			sort.Ints(IPaddresses)
		case get_ip_array <- IPaddresses:
			// Println("Noen leste arrayet..!")
		case msg := <-flush:
			// Println("Tømte arrayet..!")
			_ = msg
			IPaddresses = IPaddresses[:0]
		}
	}
}

func Timer(flush chan bool) {

	// Println("Timer startet..!")
	for {
		for timer := range time.Tick(1 * time.Second) {
			_ = timer
			flush <- true
		}
		flush <- false
	}
}

func IMA_master(get_ip_array chan []int, master, new_master chan bool) {

	// Println("IMA_master startet..!")
	count := 0
	count1 := 0
	for {
		time.Sleep(500 * time.Millisecond)
		array := <-get_ip_array
		// Println("Got array: ", array)
		if len(array) != 0 {
			if array[len(array)-1] < 255 {
				temp, _ := Atoi(GetMyIP())
				if temp == array[0] {
					count++
					if count == 2 { // SIKKERTHETSGRAD!
						// Println("Sender master request...")
						Println("MASTER forsvant..!")
						master <- true
						time.Sleep(50 * time.Microsecond)
						return
					}
					if count1 == 2 {
						new_master <- true
					}
				} else {
					count1++
				}
			} else {
				count = 0
				count1 = 0
			}
		}
	}
}
