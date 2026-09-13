package orbital

import (
	"errors"
	"fmt"
)

var (
	errOptionInvalid       = errors.New("invalid option")
	errResourceUnexist     = errors.New("not enough resource in inventory:")
	errNotEnoughMaterial   = errors.New("you donot have enough material in the inventory")
	errReactorBroken       = errors.New("cannot fix this location, when reactor isnot working")
	errLifeSupportBroken   = errors.New("cannot fix the communication, when life support isnot working")
	errCommunicationBroken = errors.New("cannot send an emergy signal, when communication isnot working")
)

func ZoneMenuValidate(option int, station Station) error {
	if station.CurrentZone.Condition == "unstable" {
		if option < 1 || option > 3 {
			return errOptionInvalid
		}
	} else {
		if option < 1 || option == 2 || option > 3 {
			return errOptionInvalid
		}
	}

	return nil
}

func FixValidate(required map[string]int, inventory map[string]int) error {
	for resource, amount := range required {
		invResource, ok := inventory[resource]
		if !ok {
			return fmt.Errorf("%w %s", errResourceUnexist, resource)
		}
		if invResource < amount {
			return errNotEnoughMaterial
		}
	}
	return nil
}

func CommunicationFixValidate(station Station) error {
	const reactorID int = 362
	const lifeSupportID int = 365

	reactor := station.ZonesByID[reactorID]
	lifeSupport := station.ZonesByID[lifeSupportID]

	if reactor.Condition == "unstable" {
		return errReactorBroken
	}

	if lifeSupport.Condition == "unstable" {
		return errLifeSupportBroken
	}
	return nil
}

func Signal(station Station) error {
	const communicationID int = 363
	communication := station.ZonesByID[communicationID]

	if communication.Condition == "unstable" {
		return errCommunicationBroken
	}

	return nil
}

func LifeSupportFixValidate(station Station) error {
	const reactorID int = 362
	reactor := station.ZonesByID[reactorID]

	if reactor.Condition == "unstable" {
		return errReactorBroken
	}
	return nil
}
