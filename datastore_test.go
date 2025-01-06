package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_Data_Store(t *testing.T) {
	time := time.Now() //testing time
	t.Run("creates a reading", func(t *testing.T) {
		want := Reading{
			Time:  time,
			Name:  "Temp",
			Value: 5.0,
		}
		//when
		got := NewReading(time, "Temp", 5.0)
		assertMatching(t, got, want)
	})
	t.Run("attaches a reading to a store with data", func(t *testing.T) {
		testReading1 := NewReading(time, "Temp", 10.0)
		got := NewStore()
		got.AddReading(testReading1)
		want := Store{
			Readings: []Reading{
				testReading1,
			},
		}

		lenOfStore := len(got.Readings)
		numOfReadings := 1
		if lenOfStore != numOfReadings {
			t.Errorf("store things it has %d readings but should have %d", lenOfStore, numOfReadings)
		}
		assertStore(t, got, want)
	})
	t.Run("Retains valid data after attaching a reading", func(t *testing.T) {

		testReading1 := NewReading(time, "Temp", 10.0)
		testReading2 := NewReading(time, "Temp", 20.0)
		store := NewStore()
		store.AddReading(testReading1)
		store.AddReading(testReading2)
		want := "Temp"
		got := store.Readings[0].Name

		assertMatching(t, got, want)
	})
	t.Run("Creates a store ", func(t *testing.T) {
		//testReading1 := NewReading(time, "Temp", 10.0)
		store := Store{}
		got := NewStore()
		assertStore(t, got, store)
	})
	t.Run("Prints out contents of a reading", func(t *testing.T) {
		store := NewStore()
		testReading1 := NewReading(time, "Temp", 10.0)
		store.AddReading(testReading1)
		buffer := &bytes.Buffer{}
		//		testReading2 := NewReading(time, "Temp", 14.0)
		//		store.AddReading(testReading2)

		store.PrintReadings(buffer)
		wantedString := []string{
			testReading1.Name,
			testReading1.Time.String(),

			"10",
		}
		want := strings.Join(wantedString, "")
		got := buffer.String()
		//		fmt.Printf("\n Got string is %v\n", got)
		//		fmt.Printf("\n want string is %v\n", want)
		assertMatching(t, got, want)
		//want := "10"
		//assertMatching(t, got, want)
	})
}

func assertStore(t *testing.T, got Store, want Store) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v and %v do not reflect equal\n", got, want)
	}
}

/*
   N – Null
   Z – Zero
   O – One
   M – Many (or More complex)
   B – Boundary Behaviors
   I – Interface definition
   E – Exercise Exceptional behavior
   S – Simple Scenarios, Simple Solutions
*/
