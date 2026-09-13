package authorization

import (
	"testing"
)

func TestModifyData(t *testing.T) {
	password := "12345"
	wantPassword := 12345
	passwordInt, err := modifyData(password)

	if err != nil {
		t.Fatal("unexpected error")
	}

	if passwordInt != wantPassword {
		t.Errorf("expected password %d, got %v", wantPassword, passwordInt)
	}
}
