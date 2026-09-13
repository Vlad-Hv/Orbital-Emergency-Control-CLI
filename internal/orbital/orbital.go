package orbital

import (
	"fmt"
)

type Station struct {
	Energy      int
	Oxygen      int
	ZonesByID   map[int]*Zone
	CurrentZone *Zone
	SignalSent  bool
}

func Create() Station {
	zones := createZones()
	return Station{
		Energy:      100,
		Oxygen:      100,
		ZonesByID:   zones,
		CurrentZone: zones[361],
	}
}

func (s *Station) ChangeZone(ID int) {
	s.CurrentZone = s.ZonesByID[ID]
}

func (s *Station) FixZone(ID int) {
	s.ZonesByID[ID].Condition = "stable"
	s.ZonesByID[ID].IsAvailable = true
}

func (s *Station) EnergyHandler() {
	reactorID := 362
	if s.ZonesByID[reactorID].Condition == "unstable" {
		s.Energy -= 10
	} else {
		s.Energy = 100
	}
}

func (s *Station) OxygenHandler(step int) {
	lifeSupportID := 365
	condition := s.ZonesByID[lifeSupportID].Condition

	if condition == "unstable" {
		if step%3 == 0 {
			s.Oxygen -= 10
			fmt.Println("oxygen leak: oxygen -10")
		} else {
			s.Oxygen -= 5
		}
	}

	if condition == "stable" {
		s.Oxygen = 100
	}
}

func (o *Station) SendSignalSOS() {
	o.SignalSent = true
}
