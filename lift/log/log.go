package log

import (
	. "../.././functions"
	. "../.././network"
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
			if msg.Dir == "int" {
				for i, lift := range last_queue {
					if lift.Ip_order == msg.Ip_order {
						Println("lift.Floor = ", lift.Floor, "and msg.FLoor = ", msg.Floor)
						if lift.Floor != msg.Floor {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							job_queue = ARQ(job_queue, msg)
							Println("DA JOB2:", job_queue)
							last_queue[i].Dir = Determine_dir(job_queue, lift)
							the_queue = Queues{job_queue, ext_queue, last_queue}
							break
						} else if lift.Dir != "standby" {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							job_queue = ARQ(job_queue, msg)
							Println("DA JOB2:", job_queue)
							last_queue[i].Dir = Determine_dir(job_queue, lift)
							the_queue = Queues{job_queue, ext_queue, last_queue}
							break
						}
					}
				}
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
				the_queue = Queues{job_queue, ext_queue, last_queue}
			} else if msg.Floor >= M {
				last_queue, _ = AIM_Dict2(last_queue, msg)
			} else if msg.Dir == "standby" {
				var update bool
				last_queue, update = AIM_Dict(last_queue, msg)
				if update {
					get_at_floor <- msg
				}
			}
		case msg := <-slave_order:

			Println("FROM LAST:", the_queue)

			if msg.Dir == "int" {
				for j, lift := range last_queue {
					if lift.Ip_order == msg.Ip_order {
						if lift.Floor != msg.Floor {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							for i, job := range job_queue {
								if job.Ip == msg.Ip_order {
									job_queue[i].Dest, _ = AIM_Int(job_queue[i].Dest, msg.Floor)
								}
							}
							last_queue[j].Dir = Determine_dir(job_queue, lift)
							the_queue = Queues{job_queue, ext_queue, last_queue}
							slave_queues <- the_queue
						} else if lift.Dir != "standby" {
							job_queue, _ = AIM_Jobs(job_queue, msg.Ip_order)
							for i, job := range job_queue {
								if job.Ip == msg.Ip_order {
									job_queue[i].Dest, _ = AIM_Int(job_queue[i].Dest, msg.Floor)
								}
							}
							last_queue[j].Dir = Determine_dir(job_queue, lift)
							the_queue = Queues{job_queue, ext_queue, last_queue}
							slave_queues <- the_queue
						}
					}
				}
			} else if msg.Ip_order == "ext" {
				ext_queue, _ = AIM_Spice(ext_queue, msg.Floor, msg.Dir)
				the_queue = Queues{job_queue, ext_queue, last_queue}
				slave_queues <- the_queue
			} else if msg.Floor >= M {
				last_queue, _ = AIM_Dict2(last_queue, msg)
			} else if msg.Dir == "standby" {
				var update bool
				last_queue, update = AIM_Dict(last_queue, msg)
				if update {
					get_at_floor <- msg
				}
			}
		case msg := <-queues:
			the_queue = msg
			job_queue = msg.Int_queue
			ext_queue = msg.Ext_queue
		case do_first <- the_queue:

		case msg := <-get_queues:
			the_queue = msg
			slave_queues <- the_queue
		case get_queues <- the_queue:
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
