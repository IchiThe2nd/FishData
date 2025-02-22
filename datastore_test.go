package data_storage

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
			TrackedNames: []string{
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
			TrackedNames: []string{
				"TrackedProbe",
			},
		}
		aStore := NewStore()
		aStore.AddTrackedNames("TrackedProbe")
		got := aStore
		//		fmt.Println(aStore.TrackedNames)
		assertStore(t, got, want)
	})
	t.Run("Adding a reading updates TrackedNames", func(t *testing.T) {
		wantStore := Store{
			TrackedNames: []string{
				"Temp",
			},
		}
		store := NewStore()
		store.AddReading(testReading1)
		got := store.TrackedNames[0]
		want := wantStore.TrackedNames[0]
		assertMatching(t, got, want)
	})
	t.Run("does not add duplicated time readings", func(t *testing.T) {
		store := NewStore()
		aStore := store.AddReading(testReading1)
		aStore.AddReading(testReading1)
		aStore.AddReading(testReading1)
		aStore.AddReading(testReading1)
		if len(aStore.Readings) > 1 {
			t.Errorf(" repeat readings added to store, there are %v entries", len(aStore.Readings))
		}

	})
	// turn this inot table test.. test adding "" ,
	t.Run("Exisiting trackedTrackedNames do not create addition entries", func(t *testing.T) {
		aStore := NewStore()
		aStore.AddTrackedNames("aName")
		aStore.AddTrackedNames("bName")
		before := len(aStore.TrackedNames)
		_, err := aStore.AddTrackedNames("aName")
		if err == nil {
			t.Errorf(" should have recived error on adding a duplicated name")
			//fmt.Print(err)
		}
		after := len(aStore.TrackedNames)

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
	t.Run("updating a storeupdates and returns the store and an err:", func(t *testing.T) {
		aStore := NewStore()
		aStore.AddTrackedNames("bob")
		bStore := aStore
		assertMatchingTypes(t, aStore, bStore)
		cStore, err := aStore.AddTrackedNames("bill")
		assertMatchingTypes(t, bStore, cStore)
		if err != nil {
			t.Error("recievd error when not expected")
		}
		assertStore(t, aStore, cStore) //double check pointers arent acting up
	})
	//NEXT TEST when given a "scan" it should see if there is Probe with "Name" from slice of ((impklemnent this first)) names in the Store. If it finds it should create a reading and append to the store. if no names are found it should warn  data imported.
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
