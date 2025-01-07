package main

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

func (store *Store) UpdateStore(scan System) (string, error) {
	lookingName := store.Names[0]
	for i, probeNames := range scan.Probes {
		if lookingName == probeNames.Name {
			fmt.Printf("found %v at scan.Probes[%v]", lookingName, i)
			return lookingName, nil
		}
	}
	fmt.Printf("\nlooking for %v\n", lookingName)
	return "", nil
}
