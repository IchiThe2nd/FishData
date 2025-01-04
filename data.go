package main

import (
	"encoding/xml"
	"fmt"
)

const input = `<System\g><hostname>Diva</hostname></system>`

type System struct {
	Hostname string `xml:"hostname"`
	Serial   string `xml:"serial"`
	Timezone string `xml:"timezone"`
	// Date     time.Time `xml:"date"`
}

func NewSystem(input string) (System, error) {
	var system System
	err := xml.Unmarshal([]byte(input), &system)
	if err != nil {
		fmt.Println(err)
	}
	if system.Hostname == "" {
		err = fmt.Errorf("recieved blank value for hostname")
	}
	//fmt.Println(system.Date.String())
	return system, err
}
