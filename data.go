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
	RawDate  string `xml:"date"` //to do create importer for time.Time format directly through overiding the time.Time portion of unmarshal on System

	Date   time.Time
	Probes []Probe `xml:"probes>probe"`
}

type Probe struct {
	Name  string  `xml:"name"`
	Value float32 `xml:"value"`
	Type  string  `xml:"type"`
}

func NewSystem(input string) (System, error) {
	var system System
	err := xml.Unmarshal([]byte(input), &system)
	if err != nil {
		fmt.Println(err)
	}

	if system.Hostname == "" {
		err = fmt.Errorf("recieved blank value for hostname")
		return system, err
	}

	stringT := system.RawDate
	pLayout := "01/02/2006 03:04:05" // The reference time, in numerical order.
	parseTime, err := time.Parse(pLayout, stringT)

	if err != nil {
		fmt.Println(err)
	}
	system.Date = parseTime

	return system, err
}
