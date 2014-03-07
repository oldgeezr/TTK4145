package log

import (
	. "../.././functions"
	. "fmt"
)

func Job_queues(master_order, slave_order, get_at_floor chan Dict, queues, get_queues, slave_queues, do_first chan Queues) {

	Fo.WriteString("Entered Job_queues\n")

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}
	the_queue := Queues{job_queue, ext_queue, last_queue}

	for {
		select {
		case msg := <-master_order:
			Fprintln(Fo, "Got messag on master_order: ", msg)
			if msg.Dir == "int" {
				for _, lift := range last_queue {
					if lift.Ip_order == msg.Ip_order {
						if lift.Floor != msg.Floor {
							Fprintln(Fo, "TRASE ORDER: Mottok int ordre på master_order")
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							Fprintln(Fo, "Opprettet jobbkø: ", job_queue)
							job_queue = ARQ(job_queue, msg)
							the_queue = Queues{job_queue, ext_queue, last_queue}
							Fprintln(Fo, "Oppdaterte The_Queue: ", the_queue)
							do_first <- the_queue
							Fprintln(Fo, "TRASE ORDER: Sendte hele the_queue til do_first")
						}
					}
				}
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
				the_queue = Queues{job_queue, ext_queue, last_queue}
				Fprintln(Fo, "Oppdaterte The_Queue: ", the_queue)
				do_first <- the_queue
			} else if msg.Dir == "last" {
				var update bool 
				last_queue, update = AIM_Dict(last_queue, msg)
				if update {
					get_at_floor <- msg
					Println(last_queue)
				}
			}
		case msg := <- slave_order:

			Pritnln("TRACE ORDER: Got messag on slave_order: ")
			Fprintln(Fo, "TRACE ORDER: Got messag on slave_order: ", msg)
			if msg.Dir == "int" {
				for _, lift := range last_queue {
					if lift.Ip_order == msg.Ip_order {
						if lift.Floor != msg.Floor {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							Fprintln(Fo, "Opprettet jobbkø: ", job_queue)
							for i, job := range job_queue {
								if job.Ip == msg.Ip_order {
									job_queue[i].Dest, _ = AIM_Int(job_queue[i].Dest, msg.Floor)
									Fprintln(Fo, "La til i jobbkøen:", job_queue[i].Dest)
								}
							}
							the_queue = Queues{job_queue, ext_queue, last_queue}
							Fprintln(Fo, "Oppdaterte The_Queue: ", the_queue)
							slave_queues <- the_queue
						}
					}
				}
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
				the_queue = Queues{job_queue, ext_queue, last_queue}
				Fprintln(Fo, "Oppdaterte The_Queue: ", the_queue)
				slave_queues <- the_queue
			} else if msg.Dir == "last" {
				var update bool 
				last_queue, update = AIM_Dict(last_queue, msg)
				if update {
					get_at_floor <- msg
					Println(last_queue)
				}
			}
		case msg := <-queues:
			Fprintln(Fo, "Fikk noe på kanal \"queues\":", msg)
			the_queue = msg
			job_queue = msg.Int_queue
			ext_queue = msg.Ext_queue
			do_first <- the_queue
		case msg := <-get_queues:
			Fprintln(Fo, "Fikk noe på get_queues")
			the_queue = msg
			slave_queues <- the_queue
			queues <- the_queue
			do_first <- the_queue
		case get_queues <- the_queue:
			Fprintln(Fo, "Noen leste på get_queues")
		}
	}
}

func ARQ(blow []Jobs, msg Dict) []Jobs {
	for i, job := range blow {
		if job.Ip == msg.Ip_order {
			blow[i].Dest, _ = AIM_Int(blow[i].Dest, msg.Floor)
			Fprintln(Fo, "La til i jobbkøen:", blow[i].Dest)
		}
	}
	return blow
}
