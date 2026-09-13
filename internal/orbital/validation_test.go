package orbital

import (
	"errors"
	"testing"
)

// положительный сценари с анстейбл, положительный с стэйбл, положительный на грани(3 и 1), и когба стейбл меньше 1, больше 1,
func TestZoneValidate(t *testing.T) {
	tests := []struct {
		name      string
		option    int
		condition string
		wantErr   error
	}{
		{name: "succesful least option", option: 1, condition: "unstable", wantErr: nil},
		{name: "succesful case", option: 2, condition: "unstable", wantErr: nil},
		{name: "succesful unstable biggest option", option: 3, condition: "unstable", wantErr: nil},
		{name: "validate unstable zero option", option: 0, condition: "unstable", wantErr: errOptionInvalid},
		{name: "validate big option", option: 4, condition: "unstable", wantErr: errOptionInvalid},
		{name: "succesful least case", option: 1, condition: "stable", wantErr: nil},
		{name: "succesful biggest case", option: 3, condition: "stable", wantErr: nil},
		{name: "validate stable zero option", option: 0, condition: "stable", wantErr: errOptionInvalid},
		{name: "validate stable biggest option", option: 4, condition: "stable", wantErr: errOptionInvalid},
		{name: "validate stable second option", option: 2, condition: "stable", wantErr: errOptionInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reactorID := 362
			station := Create()
			station.CurrentZone = station.ZonesByID[reactorID]

			station.CurrentZone.Condition = tt.condition
			err := ZoneMenuValidate(tt.option, station)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestFixValidate(t *testing.T) {
	tests := []struct {
		name      string
		resource  map[string]int
		inventory map[string]int
		wantErr   error
	}{
		{
			name: "succesful case",
			resource: map[string]int{
				"wood":  4,
				"metal": 5,
			},

			inventory: map[string]int{
				"wood":  4,
				"metal": 6,
			},

			wantErr: nil,
		},

		{
			name: "edge resources case",
			resource: map[string]int{
				"wood":  4,
				"metal": 8,
			},

			inventory: map[string]int{
				"wood":  4,
				"metal": 8,
			},

			wantErr: nil,
		},

		{
			name: "not enough resource in inventory",
			resource: map[string]int{
				"wood":  4,
				"metal": 6,
			},

			inventory: map[string]int{
				"metal": 7,
			},

			wantErr: errResourceUnexist,
		},

		{
			name: "not enough resource amount",
			resource: map[string]int{
				"wood":  5,
				"metal": 6,
			},

			inventory: map[string]int{
				"wood":  5,
				"metal": 3,
			},

			wantErr: errNotEnoughMaterial,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := FixValidate(tt.resource, tt.inventory)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestCommunicationValidate(t *testing.T) {
	tests := []struct {
		name                 string
		reactorCondition     string
		lifeSupportCondition string
		wantErr              error
	}{
		{name: "succesful case", reactorCondition: "stable", lifeSupportCondition: "stable", wantErr: nil},
		{name: "reactor is unstable", reactorCondition: "unstable", lifeSupportCondition: "stable", wantErr: errReactorBroken},
		{name: "life support is unstable", reactorCondition: "stable", lifeSupportCondition: "unstable", wantErr: errLifeSupportBroken},
		{name: "reactor and lifeSupport are unstable", reactorCondition: "unstable", lifeSupportCondition: "unstable", wantErr: errReactorBroken},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reactorID := 362
			lifeSupportID := 365
			station := Create()

			station.ZonesByID[reactorID].Condition = tt.reactorCondition
			station.ZonesByID[lifeSupportID].Condition = tt.lifeSupportCondition

			err := CommunicationFixValidate(station)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}

}

func TestSignal(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		wantErr   error
	}{
		{name: "sucessful case", condition: "stable", wantErr: nil},
		{name: "communication is unstable", condition: "unstable", wantErr: errCommunicationBroken},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			communicationID := 363
			station := Create()
			station.ZonesByID[communicationID].Condition = tt.condition

			err := Signal(station)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestLifeSupport(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		wantErr   error
	}{
		{name: "sucessful case", condition: "stable", wantErr: nil},
		{name: "reactor is unstable", condition: "unstable", wantErr: errReactorBroken},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reactorID := 362
			station := Create()
			station.ZonesByID[reactorID].Condition = tt.condition

			err := LifeSupportFixValidate(station)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
