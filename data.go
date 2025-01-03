package main

import (
	"encoding/xml"
	"fmt"
)
const input = `<Probe><hostname>Diva</hostname></Probe>`
type MyThing
func NewProbe() probe Probe) {
	var p Probe
	if err := xml.Unmarshal([]byte(input), &p); err != nil {
	}
	fmt.Println(p.Hostname)
	return p
}
