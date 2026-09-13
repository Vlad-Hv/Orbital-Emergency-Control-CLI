package state

import (
	"OStation/internal/orbital"
	"errors"
)

var (
	energyLose = errors.New("\n\nYOU LOST\nReason: u lost all your energy")
	oxygenLose = errors.New("\n\nYOU LOST\nReason: u lost all your oxygen")
	win        = errors.New("\n\nYOU WIN\nReason: u fixed all troubles and sent an emergy call")
)

func Handler(s *orbital.Station) error {
	switch {
	case s.Energy <= 0:
		return energyLose

	case s.Oxygen <= 0:
		return oxygenLose

	case s.SignalSent:
		return win

	default:
		return nil
	}
}
