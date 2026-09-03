package state

import (
	"OStation/internal/orbital"
	"errors"
)

func GameState(orbitalStation *orbital.OrbitalStation) error {
	switch {
	case orbitalStation.Energy <= 0:
		return errors.New("\n\nYOU LOST\nReason: u lost all your energy")

	case orbitalStation.Oxygen <= 0:
		return errors.New("\n\nYOU LOST\nReason: u lost all your oxygen")

	case orbitalStation.WasSignalSent:
		return errors.New("\n\nYOU WIN\nReason: u fixed all troubles and sent an emergy call")
	default:
		return nil
	}
}
