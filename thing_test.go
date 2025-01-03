package dave

import "testing"

func TestFishData(t *testing.T) {

	// Given
	input := `<hostname>Diva</hostname>`
	excpected := "Diva"

	// When
	output := FishData(input)

	// Then
	if output.Hostname != excpected {
		t.Errorf("Expected %s, but got %s", excpected, output.Hostname)
	}
}