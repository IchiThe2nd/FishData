package timerthingy // thing to schedule when to pull data from a variable source (assume in mememory stuff first then for either a file or directly from other site)

import (
	"fmt"
	"io"
	"time"
)

/* a timer that checks something  every so often.
step 1. make a timer >> make a test for a timer pinging a ch.
*/

type updater struct {
	frequency int
	// source input data we want or location
	output io.Writer // probably going to be a write call or something later.
}

func NewUpdater(duration int, output io.Writer) updater {
	updater := updater{
		frequency: duration,
		output:    output,
	}
	return updater
}

func (u updater) Start() error {
	//fmt.Fprint(u.output, "a") //just using a proxy for the actual update functionA
	freq := u.frequency
	//fmt.Printf("%d\n", freq)
	for i := freq / 5; i > 0; i-- {
		stoptimer := time.AfterFunc(5*time.Millisecond, u.dummyAction)
		//fmt.Printf("%d\n", i)
		defer stoptimer.Stop()
	}
	time.Sleep(30 * time.Millisecond)
	//time.Sleep(1 * time.Second)
	return nil
}

func (u updater) dummyAction() {
	//fmt.Println("dummyAction fired")
	fmt.Fprint(u.output, "a") //just using a proxy for the actual update functionA
}
