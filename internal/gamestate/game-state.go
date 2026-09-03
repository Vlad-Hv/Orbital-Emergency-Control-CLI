package state

import (
	"OStation/internal/orbital"
	"errors"
)

func GameState(orbitalStation *orbital.OrbitalStation) error {
	if orbitalStation.Energy <= 0 {
		return errors.New("\n\nYOU LOST\nReason: u lost all your energy")
	}

	if orbitalStation.Oxygen <= 0 {
		return errors.New("\n\nYOU LOST\nReason: u lost all your oxygen")
	}

	return nil
}
