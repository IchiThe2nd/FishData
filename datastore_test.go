package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_Data_Store(t *testing.T) {

	var aTime = time.Now()
	var testReading1 = NewReading(aTime, "Temp", 10.0)

	t.Run("creates a reading", func(t *testing.T) {
		want := testReading1
		//when
		got := NewReading(aTime, "Temp", 10.0)
		assertMatching(t, got, want)
	})
	t.Run("attaches a reading to a store with data", func(t *testing.T) {
		//	testReading1 := NewReading(time, "Temp", 10.0)
		got := NewStore()
		got.AddReading(testReading1)
		want := Store{
			Readings: []Reading{
				testReading1,
			},
			Names: []string{
				"Temp",
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

		testReading2 := NewReading(aTime, "Temp", 20.0)
		store := NewStore()
		store.AddReading(testReading1)
		store.AddReading(testReading2)
		want := "Temp"
		got := store.Readings[0].Name

		assertMatching(t, got, want)
	})
	t.Run("Creator Creates a store ", func(t *testing.T) {
		store := Store{}
		got := NewStore()
		assertStore(t, got, store)
	})
	t.Run("Store has a list for probes being tracked", func(t *testing.T) {
		//testReading1 := NewReading(time, "Temp", 10.0)
		want := Store{
			Names: []string{
				"TrackedProbe",
			},
		}
		aStore := NewStore()
		aStore.AddTrackedNames("TrackedProbe")
		got := aStore
		//		fmt.Println(aStore.Names)
		assertStore(t, got, want)
	})
	t.Run("Adding a reading updates Names", func(t *testing.T) {
		wantStore := Store{
			Names: []string{
				"Temp",
			},
		}
		store := NewStore()
		store.AddReading(testReading1)
		got := store.Names[0]
		want := wantStore.Names[0]
		assertMatching(t, got, want)
	})

	t.Run("Exisiting trackedNames do not create addition entries", func(t *testing.T) {
		aStore := NewStore()
		aStore.AddTrackedNames("aName")
		aStore.AddTrackedNames("bName")
		before := len(aStore.Names)
		aStore.AddTrackedNames("aName")
		after := len(aStore.Names)

		if before != after {
			t.Errorf("starting store length was %v and after adding an existing name store length was %v", before, after)
		}
	})
	t.Run("Prints out contents of a reading", func(t *testing.T) {
		store := NewStore()
		store.AddReading(testReading1)
		buffer := &bytes.Buffer{}

		store.PrintReadings(buffer)
		wantedString := []string{
			testReading1.Name,
			testReading1.Time.String(),

			"10",
		}
		want := strings.Join(wantedString, "")
		got := buffer.String()
		assertMatching(t, got, want)
	})
	t.Run(" Printing from an empty store errors out", func(t *testing.T) {
		store := NewStore()
		out := &bytes.Buffer{}
		err := store.PrintReadings(out)
		if err == nil {
			t.Error("Tried to print from empty Store")
		}
	})

	//NEXT TEST when given a "system" it should see if there is Probe with "Name" from slice of ((impklemnent this first)) names in the Store. If it finds it should create a reading and append to the store. if no names are found it should warn  data imported.
}

func assertStore(t *testing.T, got Store, want Store) {
	t.Helper()
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
