package data_storage

import (
	"errors"
	"fmt"
	"io"
	"time"
)

type Reading struct {
	Time  time.Time
	Name  string
	Value float32
	Type  string
}

type Store struct {
	Names    []string
	Readings []Reading
	time     time.Time
}

func NewStore() Store {
	store := Store{}
	return store
}

func NewReading(t time.Time, name string, value float32) Reading {
	reading := Reading{
		Time:  t,
		Name:  name,
		Value: value,
	}
	return reading
}

func (store *Store) AddTrackedNames(newName string) (Store, error) {
	//  cant figure out why I have to return as a pointer with out breasking stuff
	if len(store.Names) == 0 {
		store.Names = append(store.Names, newName)
	} else {
		matchingNames := 0
		for _, existingNames := range store.Names {
			if existingNames == newName {
				matchingNames++
				err := errors.New("entry already exists")
				return *store, err
			}
		}
		if matchingNames == 0 {
			store.Names = append(store.Names, newName)
		}
	}
	return *store, nil
}

func (store *Store) AddReading(newReading Reading) Store {
	store.AddTrackedNames(newReading.Name)

	for _, oldReadings := range store.Readings {
		if oldReadings.Time == newReading.Time {
			return *store
		} else {
			store.Readings = append(store.Readings, newReading)
		}
		return *store
	}

	store.Readings = append(store.Readings, newReading)
	return *store
}

func (store Store) PrintReadings(out io.Writer) error {
	s := store.Readings
	if len(s) <= 0 {
		err := errors.New("Tried to Print from an empty store")
		return err
	}
	for _, item := range s {
		fmt.Fprint(out, item.Name)
		fmt.Fprint(out, item.Time)
		fmt.Fprint(out, item.Value)

	}
	return nil
}

// search through System for names in store and return the names of items updated
func (store *Store) UpdateStore(scan System) ([]string, int, error) {
	//	lookingName := store.Names
	var foundNames []string
	var lastUpdate time.Time
	for _, lookingName := range store.Names {
		// has to be a better way than a nested for range
		for _, probeNames := range scan.Probes {

			switch {
			case lookingName == probeNames.Name:
				reading := NewReading(scan.Date, probeNames.Name, probeNames.Value)
				compareTimes := reading.Time.Compare(store.time) //// t comp u return 1 if t after u
				if compareTimes > 0 {
					store.AddReading(reading)
					foundNames = append(foundNames, lookingName)
					lastUpdate = reading.Time
				}
			}
		}
	}
	updates := len(foundNames)
	store.time = lastUpdate
	return foundNames, updates, nil
}

/*
probably better to incorporatr into update store as a flag if []reading is empty
func (store *Store) UpdateAllStore(scan System) ([]string, int, error) {
	// has to be a better way than a nested for range
	var addedNames []string
	for _, probeNames := range scan.Probes {

		//not tested yet
		reading := NewReading(time.Now(), probeNames.Name, probeNames.Value)
		store.AddReading(reading)
		addedNames = append(addedNames, probeNames.Name)
	}
	updates := len(addedNames)
	return addedNames, updates, nil
}
*/
