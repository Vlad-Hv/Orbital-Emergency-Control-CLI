package orbital

import (
	"fmt"
)

type OrbitalStation struct {
	//добавить структуру станции и дальше по плану
	Energy        int
	Oxygen        int
	Zones         map[int]*Zone
	CurrentZone   *Zone
	WasSignalSent bool
}

func CreateStation() OrbitalStation {
	zones := createZones()
	return OrbitalStation{
		Energy:      100,
		Oxygen:      100,
		Zones:       zones,
		CurrentZone: zones[361],
	}
}

func (o *OrbitalStation) ChangeZone(ID int) {
	o.CurrentZone = o.Zones[ID]
}

func (o *OrbitalStation) FixZone(ID int) {
	o.Zones[ID].Condition = "stable"
	o.Zones[ID].IsAvailable = true
}

func (o *OrbitalStation) EnergyHandler() {
	if o.Zones[362].Condition == "unstable" {
		o.Energy -= 10
	} else {
		o.Energy = 100
	}
}

func (o *OrbitalStation) OxygenHandler(step int) {
	if o.Zones[365].Condition == "unstable" {
		if step%3 == 0 {
			o.Oxygen -= 10
			fmt.Println("oxygen leak: oxygem -10")
		} else {
			o.Oxygen -= 5
		}
	} else {
		o.Oxygen = 100
	}
}

func (o *OrbitalStation) SendEmergySignal() {
	o.WasSignalSent = true
}
