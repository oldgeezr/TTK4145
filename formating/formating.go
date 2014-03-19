package formating

import (
	. ".././functions"
	. "fmt"
)

var print_counter int = 1

func Format_int_queue(int_queue []Jobs) {
	Fprintf(Fo, "Int queues:\n")
	if len(int_queue) != 0 {
		for _, job := range int_queue {
			Fprint(Fo, job.Ip)
			Fprintf(Fo, ":")
			for j := 0; j < len(job.Dest); j++ {
				Fprint(Fo, job.Dest[j].Floor)
				Fprintf(Fo, " ")
			}
			Fprintf(Fo, "\n")
		}
	} else {
		Fprintf(Fo, "<empty> \n")
	}
}

func Format_ext_queue(ext_queue []Dict) {
	Fprintf(Fo, "Ext queue:\n")
	if len(ext_queue) != 0 {
		for j := 0; j < len(ext_queue); j++ {
			Fprint(Fo, ext_queue[j].Floor)
			Fprintf(Fo, "->")
			Fprint(Fo, ext_queue[j].Dir)
			Fprintf(Fo, "\n")
		}
	} else {
		Fprintf(Fo, "<empty> \n")
	}
}

func Format_last_queue(last_queue []Dict) {
	Fprintf(Fo, "Last queues:\n")
	if len(last_queue) != 0 {
		for j := 0; j < len(last_queue); j++ {
			Fprint(Fo, last_queue[j].Ip_order)
			Fprintf(Fo, ":")
			Fprint(Fo, last_queue[j].Floor)
			Print("->")
			Print(last_queue[j].Dir)
			Fprintf(Fo, "\n")
		}
	} else {
		Fprintf(Fo, "<empty> \n")
	}
}

func Format_queues(queues Queues) {
	Fprintf(Fo, "---------Queues--------\n")
	Format_int_queue(queues.Int_queue)
	Fprintf(Fo, "\n")
	Format_ext_queue(queues.Ext_queue)
	Fprintf(Fo, "\n")
	Format_last_queue(queues.Last_queue)
	Fprintf(Fo, "-----------------------\n")
}

////////////////////////////////////////////////

func Format_int_queue_term(int_queue []Jobs) {
	Print("Int queues:\n")
	if len(int_queue) != 0 {
		for _, job := range int_queue {
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

func Format_ext_queue_term(ext_queue []Dict) {
	Print("Ext queue:\n")
	if len(ext_queue) != 0 {
		for j := 0; j < len(ext_queue); j++ {
			Print(ext_queue[j].Floor)
			Print("->")
			Print(ext_queue[j].Dir)
			Print("\n")
		}
	} else {
		Print("<empty> \n")
	}
}

func Format_last_queue_term(last_queue []Dict) {
	Print("Last queues:\n")
	if len(last_queue) != 0 {
		for j := 0; j < len(last_queue); j++ {
			Print(last_queue[j].Ip_order)
			Print(":")
			Print(last_queue[j].Floor)
			Print("->")
			Print(last_queue[j].Dir)
			Print("\n")
		}
	} else {
		Print("<empty> \n")
	}
}

func Format_queues_term(queues Queues) {
	Print("\n")
	Print("#",print_counter)
	Print("-------Queues--------\n")
	Format_int_queue_term(queues.Int_queue)
	Print("\n")
	Format_ext_queue_term(queues.Ext_queue)
	Print("\n")
	Format_last_queue_term(queues.Last_queue)
	Print("-----------------------\n")
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
