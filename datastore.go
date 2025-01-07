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

func (store *Store) AddReading(reading Reading) *Store {
	store.AddTrackedNames(reading.Name)
	store.Readings = append(store.Readings, reading)
	return store
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
