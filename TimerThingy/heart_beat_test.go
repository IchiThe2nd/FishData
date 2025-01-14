package timerthingy

import (
	"bytes"
	"testing"
)

// first test just printing out at all and then use time compare and print to verify
func Test_createTimer(t *testing.T) {
	duration := 5
	output := &bytes.Buffer{}
	updater := NewUpdater(duration, output)
	updater.Start()
	want := "a"
	got := output.String()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}

}
