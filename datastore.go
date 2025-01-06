package main

import (
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
	Readings []Reading
}

func NewStore() Store {
	store := Store{}
	// fmt.Print("created new store\n")
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

func (store *Store) AddReading(reading Reading) *Store {
	store.Readings = append(store.Readings, reading)

	return store
}
func (store Store) PrintReadings(out io.Writer) {
	s := store.Readings
	for _, item := range s {
		fmt.Fprint(out, item.Name)
		fmt.Fprint(out, item.Time)
		fmt.Fprint(out, item.Value)
	}
}
