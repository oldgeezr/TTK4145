package formating

import (
	. ".././functions"
	. "fmt"
)

var print_counter int = 1

func Format_job_queue_term(job_queue []Jobs) {
	Print("Job queues:\n")
	if len(job_queue) != 0 {
		for _, job := range job_queue {
			Print(job.Ip)
			Print(":")
			for j := 0; j < len(job.Dest); j++ {
				Print(job.Dest[j].Floor)
				Print(" ")
			}
			Print("\n")
		}
	} else {
		Print("<empty> \n")
	}
}

func Format_last_queue_graphic(last_queue []Dict) {
	Print("Elevators:\n")
	if len(last_queue) != 0 {
		for j := 0; j < len(last_queue); j++ {
			Print(last_queue[j].Ip_order)
			Print(":")
			if last_queue[j].Dir == "up" {
				switch {
				case last_queue[j].Floor == 0:
					Println("|H->| 1 | 2 | 3 |")
				case last_queue[j].Floor == 1:
					Println("| 0 |H->| 2 | 3 |")
				case last_queue[j].Floor == 2:
					Println("| 0 | 1 |H->| 3 |")
				case last_queue[j].Floor == 3:
					Println("| 0 | 1 | 2 |H->|")
				}
			} else if last_queue[j].Dir == "down" {
				switch {
				case last_queue[j].Floor == 0:
					Println("|<-H| 1 | 2 | 3 |")
				case last_queue[j].Floor == 1:
					Println("| 0 |<-H| 2 | 3 |")
				case last_queue[j].Floor == 2:
					Println("| 0 | 1 |<-H| 3 |")
				case last_queue[j].Floor == 3:
					Println("| 0 | 1 | 2 |<-H|")
				}
			} else {
				switch {
				case last_queue[j].Floor == 0:
					Println("| H | 1 | 2 | 3 |")
				case last_queue[j].Floor == 1:
					Println("| 0 | H | 2 | 3 |")
				case last_queue[j].Floor == 2:
					Println("| 0 | 1 | H | 3 |")
				case last_queue[j].Floor == 3:
					Println("| 0 | 1 | 2 | H |")
				}
			}
		}
	} else {
		Print("<empty> \n")
	}
}

func Format_ext_queue_term(ext_queue []Dict) {
	Print("Ext queue:\n")
	if len(ext_queue) != 0 {
		for j := 0; j < len(ext_queue); j++ {
			if ext_queue[j].Ip_order != "ext" {
				Print("*")
			}
			Print(ext_queue[j].Floor)
			Print("->")
			Print(ext_queue[j].Dir)
			Print("\n")
		}
	} else {
		Print("<empty> \n")
	}
}

func Format_queues_term(queues Queues, state string) {
	Print("\n")
	Print("#", print_counter)
	Print("---Queues--")
	Print("(", state)
	Print(")\n")
	Format_job_queue_term(queues.Job_queue)
	Print("\n")
	Format_ext_queue_term(queues.Ext_queue)
	Print("\n")
	Format_last_queue_graphic(queues.Last_queue)
	//Format_last_queue_term(queues.Last_queue)
	Print("---------------------\n")
	print_counter++
}

func Elevator_art() {
	Println("   ______      ________                __            ")
	Println("  / ____/___  / ____/ /__ _   ______ _/ /_____  _____")
	Println(" / / __/ __ \\/ __/ / / _ \\ | / / __ `/ __/ __ \\/ ___/")
	Println("/ /_/ / /_/ / /___/ /  __/ |/ / /_/ / /_/ /_/ / /    ")
	Println("\\____/\\____/_____/_/\\___/|___/\\__,_/\\__/\\____/_/     ")
	Println("                                                     ")
	Println("By: Christoffer Ramstad-Evensen and Erlend Hestnes")
}
