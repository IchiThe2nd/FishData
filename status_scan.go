package data_storage

import (
	"encoding/xml"
	"fmt"
	"time"
)

const input = `<Scan\g><hostname>Diva</hostname></scan>`

type Scan struct {
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

func NewScan(input string) (Scan, error) { //this is prroabbly name wrong now change to newScan
	var scan Scan
	err := xml.Unmarshal([]byte(input), &scan)
	if err != nil {
		fmt.Println(err)
		return scan, err
	}

	if scan.Hostname == "" {
		err = fmt.Errorf("recieved blank value for hostname")
		return scan, err
	}

	scan.Date = convTime(scan.RawDate, "01/02/2006 03:04:05")
	return scan, err
}

func convTime(input, layout string) time.Time {
	adate, err := time.Parse(layout, input)

	if err != nil {
		fmt.Println(err)
	}
	return adate
}
