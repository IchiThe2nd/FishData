package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

const input = `<System\g><hostname>Diva</hostname></system>`

type System struct {
	Hostname string `xml:"hostname"`
	Serial   string `xml:"serial"`
	Timezone string `xml:"timezone"`
	RawDate  string `xml:"date"`

	Date   time.Time
	Probes []Probe `xml:"probes>probe"`
}

type Probe struct {
	Name  string  `xml:"name"`
	Value float32 `xml:"value"`
	Type  string  `xml:"type"`
}

func NewSystem(input string) (System, error) { //this is prroabbly name wrong now change to newScan
	var system System
	err := xml.Unmarshal([]byte(input), &system)
	if err != nil {
		fmt.Println(err)
	}

	if system.Hostname == "" {
		err = fmt.Errorf("recieved blank value for hostname")
		return system, err
	}

	system.Date = convTime(system.RawDate, "01/02/2006 03:04:05")
	return system, err
}

func convTime(input, layout string) time.Time {
	adate, err := time.Parse(layout, input)

	if err != nil {
		fmt.Println(err)
	}
	return adate
}
