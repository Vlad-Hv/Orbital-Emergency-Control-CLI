package orbital

type OrbitalStation struct {
	//добавить структуру станции и дальше по плану
	Energy      int
	Oxygen      int
	Zones       map[int]*Zone
	CurrentZone *Zone
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

func (orbital *OrbitalStation) ChangeZone(ID int) {
	orbital.CurrentZone = orbital.Zones[ID]
}

func (orbital *OrbitalStation) FixZone(ID int) {
	orbital.Zones[ID].Condition = "stable"
	orbital.Zones[ID].IsAvailable = true
}

func (orbital *OrbitalStation) EnergyHandler() {
	if orbital.Zones[362].Condition == "unstable" {
		orbital.Energy -= 10
	} else {
		orbital.Energy = 100
	}
}
