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

func (store *Store) AddTrackedNames(newName string) *Store {
	store.Names = append(store.Names, newName)
	return store
}

func (store *Store) AddReading(reading Reading) *Store {
	store.Names = append(store.Names, reading.Name)
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
