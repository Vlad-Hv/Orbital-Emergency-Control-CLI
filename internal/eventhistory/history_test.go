package history

import (
	"testing"
)

func TestAdd(t *testing.T) {
	history := Create()
	message := "history package testing"
	wantLen := 1

	if history != nil {
		t.Fatal("expected empty history")
	}

	history.Add(message)

	if len(history) != wantLen {
		t.Fatalf("expected lenth %d, got %d", wantLen, len(history))
	}
	if history[0] != message {
		t.Errorf("expected message %v, got %v", message, history[0])
	}
}
