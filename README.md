Go!Elevator
===========
By Christoffer Ramstad-Evensen & Erlend Hestnes

This project conserns a realtime system for controlling elevators on a network. The systems is based on N elevators covering M floors communicating over TCP. UDP is used to keep track of which elevators are on the network. Drivers and source code are included. 

![](https://raw.github.com/oldgeezr/sanntid/master/figures/lifts.png)

System illustration

Main.go
--------

Main.go creates a phoenix program that listens to the master program. The idea behind this solution is to keep track of the master program and clone a new master from the slave program if master crashes. Main.go initiates goelevator.go

The elevator queue system
--------------------------

![](https://raw.github.com/oldgeezr/sanntid/master/figures/elevator_queues.png)

The queue-system is designed in such a way that every elevator should know what each elevator is doing at all times. So that if one elevator goes down and comes back again, it will continue. This is also the case if master elevator goes down.

The elevator algorithm
------------------------

![](https://raw.github.com/oldgeezr/sanntid/master/figures/elevator_algorithm.png)

The algorithm is only executed on the master elevator. The algorithm triggers when an elevator reaches a new floor or if a stop signal is sent. The algorithm checks for internal and external orders for the elevator that triggered it. If it finds an external order in the same direction that the elevator is going, it will append it to the job_queue, then remove it.

Project package structure
------------------------

![](https://raw.github.com/oldgeezr/sanntid/master/figures/project_package_structure.png)
