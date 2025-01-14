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
	eternal   bool
	// source input data we want or location
	output io.Writer // probably going to be a write call or something later.
}

func NewUpdater(duration int, output io.Writer, eternal bool) updater {
	updater := updater{
		frequency: duration,
		output:    output,
	}
	return updater
}

func (u updater) Start() error {
	freq := u.frequency

	for range time.Tick(time.Millisecond * 5) {
		u.dummyAction()
		freq = freq - 5
		if freq == 0 {
			return nil
		}
	}
	return nil
}

func (u updater) dummyAction() {
	//fmt.Println("dummyAction fired")
	fmt.Fprint(u.output, "a") //just using a proxy for the actual update functionA
}
