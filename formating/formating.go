package formating

import (
	. ".././functions"
	. "fmt"
)

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
	Fprintf(Fo, "Last queue:\n")
	if len(last_queue) != 0 {
		for j := 0; j < len(last_queue); j++ {
			Fprint(Fo, last_queue[j].Ip_order)
			Fprintf(Fo, ":")
			Fprint(Fo, last_queue[j].Floor)
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
	Fprintf(Fo, "	   V\n")
}
