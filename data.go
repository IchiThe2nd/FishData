package main

import (
	"encoding/xml"
	"fmt"
)

type Probe struct {
	Hostname string   `xml:"hostname"`
	XMLName  xml.Name `xml:"name"`

	//Name    string   `xml:"name"`
	Value float64 `xml:"value"`
	Type  string  `xml:"type"`
}

const input = `<Probe><hostname>Diva</hostname></Probe>`

func NewProbe() (probe Probe) {
	var p Probe
	if err := xml.Unmarshal([]byte(input), &p); err != nil {
		//panic(err)
	}
	fmt.Println(p.Hostname)
	//	fmt.Println(p.Name)
	return p
}
