package storage

import (
	"testing"
)

func TestCheckStorageData(t *testing.T) {
	tests := []struct {
		name     string
		resource string
		amount   int
		storage  map[string]int
		wantErr  error
	}{
		{
			name:     "succesful case",
			resource: "wood",
			amount:   5,
			storage: map[string]int{
				"wood": 6,
			},
			wantErr: nil,
		},

		{
			name:     "incorrect name",
			resource: "woo",
			amount:   5,
			storage: map[string]int{
				"wood": 6,
			},
			wantErr: errNameIncorrect,
		},

		{
			name:     "not enough resource",
			resource: "wood",
			amount:   8,
			storage: map[string]int{
				"wood": 5,
			},
			wantErr: errResourceNotEnough,
		},

		{
			name:     "invalid amount",
			resource: "wood",
			amount:   -4,
			storage: map[string]int{
				"wood": 5,
			},
			wantErr: errAmountInvalid,
		},

		{
			name:     "edge invalid amount case",
			resource: "wood",
			amount:   0,
			storage: map[string]int{
				"wood": 5,
			},
			wantErr: errAmountInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckData(tt.resource, tt.amount, tt.storage)
			if err != tt.wantErr {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
