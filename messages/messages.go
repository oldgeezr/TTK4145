package messages

import(
    // . "fmt"
    "sort"
    "time"
    ."strconv"
    .".././network"
)

func IP_array(array_update chan int, get_array chan []int, flush chan bool) {
    
    IPaddresses := []int{}

    for {
        
        select {
        case ip := <-array_update:
            
	            IPaddresses = AppendIfMissing(IPaddresses, ip)
	            sort.Ints(IPaddresses)

        case get_array<- IPaddresses:

        case msg := <-flush:
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

    for {
        for timer := range time.Tick(time.Second) {
            _ = timer
            flush <- true
        }
        flush <- false
    }
}

func IMA_master(get_array chan []int, master chan bool) {

    for {

        time.Sleep(1333*time.Millisecond)
        array := <-get_array
        if len(array) != 0 {
            if array[len(array)-1] != 300 {
                temp, _:= Atoi(GetMyIP())
                if array[0] == temp {
                    master <- true
                }
            }
        }
    }    
}

