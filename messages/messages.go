package messages

import (
	. ".././network"
	. "fmt"
	"sort"
	. "strconv"
	"time"
)

func IP_array(array_update chan int, get_array chan []int, flush chan bool) {

	IPaddresses := []int{}
	// Println("IP_array startet..!")

	for {

		select {
		case ip := <-array_update:

			// Println("Oppdaterte arrayet..!")
			IPaddresses = AppendIfMissing(IPaddresses, ip)
			sort.Ints(IPaddresses)

		case get_array <- IPaddresses:
			// Println("Noen leste arrayet..!")

		case msg := <-flush:
			// Println("Tømte arrayet..!")
			_ = msg
			IPaddresses = IPaddresses[:0]
		}

	}
}

func AppendIfMissing(slice []int, i int) []int {

	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func Timer(flush chan bool) {

	// Println("Timer startet..!")

	for {
		for timer := range time.Tick(2 * time.Second) {
			_ = timer
			flush <- true
		}
		flush <- false
	}
}

func IMA_master(get_array chan []int, master chan bool) {

	// Println("IMA_master startet..!")
	count := 0

	for {

		time.Sleep(500 * time.Millisecond)
		array := <-get_array
		// Println("Got array: ", array)
		if len(array) != 0 {
			if array[len(array)-1] < 255 {
				temp, _ := Atoi(GetMyIP())
				if temp == array[0] {
					count++
					if count == 3 { // SIKKERTHETSGRAD!
						// Println("Sender master request...")
						Println("MASTER forsvant..!")
						master <- true
					}
				}
			} else {
				count = 0
			}
		}
	}
}
