package log

import (
	. "../.././formating"
	. "../.././functions"
	. "../.././network"
	. "fmt"
)

func Job_queues(master_order, slave_order, get_at_floor chan Dict, queues, get_queues, set_queues, slave_queues, do_first chan Queues) {

	Fo.WriteString("Entered Job_queues\n")

	last_queue := []Dict{}
	job_queue := []Jobs{}
	ext_queue := []Dict{}
	the_queue := Queues{job_queue, ext_queue, last_queue}

	for {
		select {
		case msg := <-master_order: // MASTER
			if msg.Dir == "int" { // Fikk internordre
				for i, lift := range last_queue {
					if lift.Ip_order == msg.Ip_order {
						if lift.Floor != msg.Floor {
							job_queue = ARQ(job_queue, msg)
							last_queue[i].Dir = Determine_dir(job_queue, lift)
							break
						} else if lift.Dir != "standby" {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							job_queue = ARQ(job_queue, msg)
							last_queue[i].Dir = Determine_dir(job_queue, lift)
							break
						}
					}
				}
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
			} else if msg.Floor >= M {
				last_queue, _ = AIM_Dict2(last_queue, msg)
			} else if msg.Dir == "standby" {
				if len(last_queue) != 0 {
					for _, last := range last_queue {
						if last.Ip_order != msg.Ip_order {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
						}
					}
				} else {
					job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
				}
				var update bool
				last_queue, update = AIM_Dict(last_queue, msg)
				if update {
					get_at_floor <- msg
				}
			} else if msg.Dir == "remove" {
				get_at_floor <- msg
			}
			the_queue = Queues{job_queue, ext_queue, last_queue}
		case msg := <-slave_order:
			if msg.Dir == "int" {
				for j, lift := range last_queue {
					if lift.Ip_order == msg.Ip_order {
						if lift.Floor != msg.Floor {
							for i, job := range job_queue {
								if job.Ip == msg.Ip_order {
									job_queue[i].Dest, _ = AIM_Int(job_queue[i].Dest, msg.Floor)
								}
							}
							last_queue[j].Dir = Determine_dir(job_queue, lift)
						} else if lift.Dir != "standby" {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							for i, job := range job_queue {
								if job.Ip == msg.Ip_order {
									job_queue[i].Dest, _ = AIM_Int(job_queue[i].Dest, msg.Floor)
								}
							}
							last_queue[j].Dir = Determine_dir(job_queue, lift)
						}
					}
				}
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
			} else if msg.Floor >= M {
				last_queue, _ = AIM_Dict2(last_queue, msg)
			} else if msg.Dir == "standby" {
				if len(last_queue) != 0 {
					for _, last := range last_queue {
						if last.Ip_order != msg.Ip_order {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
						}
					}
				} else {
					job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
				}
				var update bool
				last_queue, update = AIM_Dict(last_queue, msg)
				if update {
					get_at_floor <- msg
				}
			} else if msg.Dir == "remove" {
				get_at_floor <- msg
			}
			the_queue = Queues{job_queue, ext_queue, last_queue}
			slave_queues <- the_queue
		case msg := <-set_queues:
			Println("FROM ALGO:")
			Format_queues_term(msg)
			the_queue.Int_queue = msg.Int_queue
			the_queue.Ext_queue = msg.Ext_queue
			the_queue.Last_queue = msg.Last_queue
			// Println("TO LOG:")
			// Format_queues_term(the_queue)
			slave_queues <- the_queue
		case msg := <-queues:
			the_queue.Int_queue = msg.Int_queue
			the_queue.Ext_queue = msg.Ext_queue
			the_queue.Last_queue = msg.Last_queue
		case do_first <- the_queue: // DO FIRST
		case get_queues <- the_queue: // ALGO
			Println("TO ALGO:")
			Format_queues_term(the_queue)
		}
	}
}

func ARQ(blow []Jobs, msg Dict) []Jobs {
	for i, job := range blow {
		if job.Ip == msg.Ip_order {
			blow[i].Dest, _ = AIM_Int(blow[i].Dest, msg.Floor)
		}
	}
	return blow
}

func Determine_dir(job_queue []Jobs, last Dict) string {
	for _, job := range job_queue {
		if last.Ip_order == job.Ip {
			if len(job.Dest) != 0 {
				if job.Dest[0].Floor-last.Floor > 0 {
					return "up"
				} else if job.Dest[0].Floor-last.Floor < 0 {
					return "down"
				} else {
					return "standby"
				}
			} else {
				return "standby"
			}
		}
	}
	return "standby"
}
