package storage

import (
	"testing"
)

func TestCreateStorage(t *testing.T) {
	storage := Create()
	wantLen := 4
	wantMap := map[string]int{
		"metal":   13,
		"fuel":    5,
		"wire":    6,
		"medical": 3,
	}

	if len(storage) != wantLen {
		t.Fatalf("expected storage lenth %d, got %d", wantLen, len(storage))
	}

	for resource, amount := range storage {
		if wantMap[resource] != amount {
			t.Errorf("expected amount %d resource %v, got %d", wantMap[resource], resource, amount)
		}
	}
}
