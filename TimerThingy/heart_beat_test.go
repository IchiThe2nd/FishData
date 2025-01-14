package timerthingy

import (
	"bytes"
	"testing"
)

/*
 */
func Test_timers_properties(t *testing.T) {
	tests := map[string]struct {
		input  int
		result string
	}{
		"fires one time after 5 seconds": {
			input:  5,
			result: "a",
		},
		"fires two times": {
			input:  10,
			result: "aa",
		},
		"fires three times": {
			input:  15,
			result: "aaa",
		},
	}

	for name, test := range tests {

		t.Run(name, func(t *testing.T) {

			t.Parallel()

			duration := test.input
			want := test.result
			output := &bytes.Buffer{}
			updater := NewUpdater(duration, output)
			updater.Start()

			got := output.String()

			if got != want {
				t.Errorf("got %v but wanted %v", got, want)
			}
		})
	}

}
