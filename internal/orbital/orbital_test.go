package orbital

import (
	"testing"
)

func TestCreateStation(t *testing.T) {
	station := Create()

	wantEnergy := 100
	wantOxygen := 100
	wantCurrentZone := station.ZonesByID[361]

	if wantCurrentZone == nil {
		t.Fatal("expected zone ID 361, got nil")
	}

	if station.Energy != wantEnergy {
		t.Errorf("expected %d energy, actual %d", wantEnergy, station.Energy)
	}

	if station.Oxygen != wantOxygen {
		t.Errorf("expected %d oxygen, actual %d", wantOxygen, station.Oxygen)
	}

	if station.CurrentZone != wantCurrentZone {
		t.Error("expected zone ID 361")
	}
}

func TestChangeZone(t *testing.T) {
	station := Create()
	station.ChangeZone(362)

	wantCurrentZone := station.ZonesByID[362]

	if station.CurrentZone != wantCurrentZone {
		t.Errorf("expected current zone ID 362")
	}
}

func TestFixZone(t *testing.T) {
	var zoneID int = 363
	station := Create()
	station.CurrentZone = station.ZonesByID[zoneID]

	wantCondition := "stable"
	wantAvailability := true
	if station.CurrentZone.Condition != "unstable" || station.CurrentZone.IsAvailable == true {
		t.Fatal("zone is already stable or available")
	}

	station.FixZone(zoneID)
	if station.CurrentZone.Condition != wantCondition {
		t.Errorf("expected %v condition, got %v", wantCondition, station.CurrentZone.Condition)
	}

	if station.CurrentZone.IsAvailable != wantAvailability {
		t.Errorf("expected availability true, got %v", station.CurrentZone.IsAvailable)
	}
}

func TestEnergyHandler(t *testing.T) {
	tests := []struct {
		name       string
		condition  string
		energy     int
		wantEnergy int
	}{
		{name: "unstable condition", condition: "unstable", energy: 60, wantEnergy: 50},
		{name: "stable condition", condition: "stable", energy: 40, wantEnergy: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reactorID := 362
			station := Create()
			station.ZonesByID[reactorID].Condition = tt.condition
			station.Energy = tt.energy

			station.EnergyHandler()

			if station.Energy != tt.wantEnergy {
				t.Errorf("expected %d energy, got %d", tt.wantEnergy, station.Energy)
			}
		})
	}
}

func TestOxygenHandler(t *testing.T) {
	tests := []struct {
		name       string
		condition  string
		step       int
		oxygen     int
		wantOxygen int
	}{
		{
			name:       "unstable condition",
			condition:  "unstable",
			step:       1,
			oxygen:     60,
			wantOxygen: 55,
		},
		{
			name:       "unstable condition, oxygen leak",
			condition:  "unstable",
			step:       3,
			oxygen:     90,
			wantOxygen: 80,
		},
		{
			name:       "stable condition",
			condition:  "stable",
			step:       3,
			oxygen:     60,
			wantOxygen: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lifeSupportID := 365
			station := Create()

			station.ZonesByID[lifeSupportID].Condition = tt.condition
			station.Oxygen = tt.oxygen
			station.OxygenHandler(tt.step)
			if station.Oxygen != tt.wantOxygen {
				t.Errorf("expected %d oxygen, got %d", tt.wantOxygen, station.Oxygen)
			}
		})
	}
}

func TestSendSignalSOS(t *testing.T) {
	station := Create()

	station.SendSignalSOS()
	if !station.SignalSent {
		t.Error("expected signal sent")
	}
}
