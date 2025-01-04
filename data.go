package main

import (
	"encoding/xml"
	"fmt"
)

const input = `<Probe><hostname>Diva</hostname></Probe>`

type Probe struct {
	Hostname string `xml:"hostname"`
}

func NewProbe(input string) Probe {
	var probe Probe
	err := xml.Unmarshal([]byte(input), &probe)
	if err != nil {
		fmt.Println(err)
	}
	return probe
}
