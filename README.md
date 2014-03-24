Go!Elevator
===========
By Christoffer Ramstad-Evensen & Erlend Hestnes


<!-- Main.go
--------

![](https://raw.github.com/oldgeezr/sanntid/coffee/figures/elevator_main.png)

Main.go creates almost all the channels the program uses, and creates the necessary go-threads for the program. In essence, main.go is a simple state machine.

-->

The elevator queue system
--------------------------

![](https://github.com/oldgeezr/sanntid/blob/master/figures/elevator_queues.png)

The queue-system is designed in such a way that every elevator should know what each elevator is doing at all times. So that if one elevator goes down and comes back again, it will continue. This is also the case if master elevator goes down.

<!-- The elevator log module
-------------------------

![](https://raw.github.com/oldgeezr/sanntid/coffee/figures/elevator_log_module.png)

This module maintains the elevator-queues.

-->

The elevator algorithm
------------------------

![](https://raw.github.com/oldgeezr/sanntid/coffee/figures/elevator_algorithm.png)

The algorithm is only executed on the master elevator. The algorithm triggers when an elevator reaches a new floor or if a stop signal is sent. The algorithm checks for internal and external orders for the elevator that triggered it. If it finds an external order in the same direction that the elevator is going, it will append it to the job_queue, then remove it.

Project package structure
------------------------

![](https://github.com/oldgeezr/sanntid/blob/master/figures/project_package_structure.png)
