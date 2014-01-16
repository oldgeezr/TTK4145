package main 

import "fmt"
import "time"

func A(ch chan string) {
	for {
	if <-ch == "A" {
			ch<- "a"
		}
	}
}

func B(ch chan string) {
	for {
		if <-ch == "B" {
			ch<- "b"
		}
	}
}

func C(ch chan string) {
	for {
		if <-ch == "C" {
			ch<- "c"
		}
	}
}

func main() {

	ch := make(chan string)
	
	go A(ch)
	go B(ch)
	go C(ch)

	ch<- "A"

	for {
		time.Sleep(1000*time.Millisecond)

		if <-ch == "a" {
			fmt.Println("Du er i A \n")
			ch<- "B"
		} else if <-ch == "b" {
			fmt.Println("Du er i B \n")
			ch<- "C"
		} else if <-ch == "c" {
			fmt.Println("Du er i C \n")
			ch<- "A"
		}
	}
}
