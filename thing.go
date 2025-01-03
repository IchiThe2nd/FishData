package dave

import (
	"encoding/xml"
	"fmt"
)

type MyThing struct {
	Hostname string   `xml:"hostname"`
}

func FishData(input string) MyThing {
	var thing MyThing

	err := xml.Unmarshal([]byte(input), &thing)
	if err != nil {
		fmt.Println(err)
	}

	return thing
}