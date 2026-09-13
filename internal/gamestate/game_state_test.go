package state

import (
	"OStation/internal/orbital"
	"errors"
	"testing"
)

func TestGameState(t *testing.T) {
	tests := []struct {
		name    string
		energy  int
		oxygen  int
		wasSend bool
		wantErr error
	}{
		{
			name:    "success case",
			energy:  100,
			oxygen:  100,
			wasSend: false,
			wantErr: nil,
		},

		{
			name:    "energy zero case lose",
			energy:  0,
			oxygen:  100,
			wasSend: false,
			wantErr: energyLose,
		},

		{
			name:    "energy less than zero case lose",
			energy:  -1,
			oxygen:  100,
			wasSend: false,
			wantErr: energyLose,
		},

		{
			name:    "oxygen zero lose case",
			energy:  100,
			oxygen:  0,
			wasSend: false,
			wantErr: oxygenLose,
		},

		{
			name:    "oxygen less than zero lose case",
			energy:  100,
			oxygen:  -1,
			wasSend: false,
			wantErr: oxygenLose,
		},

		{
			name:    "win case",
			energy:  100,
			oxygen:  100,
			wasSend: true,
			wantErr: win,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			station := orbital.Create()

			station.Energy = tt.energy
			station.Oxygen = tt.oxygen
			station.SignalSent = tt.wasSend

			err := Handler(&station)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
