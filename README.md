Go!Elevator
===========
By Christoffer Ramstad-Evensen & Erlend Hestnes


Main.go
--------

![](https://raw.github.com/oldgeezr/sanntid/coffee/figures/elevator_main.png)

Main.go creates almost all the channels the program uses, and creates the necessary go-threads for the setup. 


The elevator queue system
--------------------------

![](https://raw.github.com/oldgeezr/sanntid/coffee/figures/elevator_queues.png)

The queue-system is designed in such a way that every elevator should know what each elevator is doing at all times.

The elevator log module
-------------------------

![](https://raw.github.com/oldgeezr/sanntid/coffee/figures/elevator_log_module.png)

This module maintains the elevator-queues.

The elevator algorithm
------------------------

![](https://raw.github.com/oldgeezr/sanntid/coffee/figures/elevator_algorithm.png)

The algorithm is only executed on the master elevator. The algorithm triggers when an elevator reaches a new floor.
