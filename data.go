package main

import (
	"encoding/xml"
	"fmt"
)

const input = `<System\g><hostname>Diva</hostname></system>`

type System struct {
	Hostname string `xml:"hostname"`
}

func NewSystem(input string) System {
	var system System
	err := xml.Unmarshal([]byte(input), &system)
	if err != nil {
		fmt.Println(err)
	}
	return system
}
