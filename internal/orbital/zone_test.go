package orbital

import (
	"testing"
)

func TestZone(t *testing.T) {
	zones := createZones()
	wantAmount := 5

	if len(zones) != wantAmount {
		t.Errorf("expected zones amount %d, got %d", wantAmount, len(zones))
	}

	for i := 361; i <= 365; i++ {
		_, ok := zones[i]
		if !ok {
			t.Errorf("expected zone witb ID %d is exist", i)
		}
	}
}
