package interrupts

import . "fmt"
import "os"
import "os/signal"
import "syscall"

func Disable_interrupts(enable bool) {
	if enable {
		ctrl_c := make(chan os.Signal, 1)
		ctrl_z := make(chan os.Signal, 1)
		ctrl_q := make(chan os.Signal, 1)

		signal.Notify(ctrl_c, syscall.SIGINT)
		signal.Notify(ctrl_z, syscall.SIGTSTP)
		signal.Notify(ctrl_q, syscall.SIGQUIT)

		for {
			select {
			case <-ctrl_c:
				Println("Got CTRL-C: SIGINT. n00b ...")
			case <-ctrl_z:
				Println("Got CTRL-Z: SIGTSTP. n00b ...")
			case <-ctrl_q:
				Println("Got CTRL-\\: SIGQUIT. n00b ...")
			}
		}
	}
}
