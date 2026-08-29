package orbital

type Zone struct {
	ID          int
	Name        string
	IsAvailable bool
	StuffToFix  map[string]int
	Condition   string
}

func createZones() map[int]*Zone {
	return map[int]*Zone{
		361: {
			ID:          361,
			Name:        "Control room",
			IsAvailable: true,
			StuffToFix:  map[string]int{},
			Condition:   "stable",
		},

		362: {
			ID:          362,
			Name:        "Reactor",
			IsAvailable: false,
			StuffToFix: map[string]int{
				"accessCard": 1, //special tool in control room
				"metal":      5,
				"fuel":       2,
			},
			Condition: "unstable",
		},

		363: {
			ID:          363,
			Name:        "Communications",
			IsAvailable: false,
			StuffToFix: map[string]int{
				"metal":   2,
				"wire":    6,
				"medical": 3,
			},
			Condition: "unstable",
		},

		364: {
			ID:          364,
			Name:        "Storage",
			IsAvailable: true,
			StuffToFix:  map[string]int{},
			Condition:   "stable",
		},

		365: {
			ID:          365,
			Name:        "Life Support",
			IsAvailable: false,
			StuffToFix: map[string]int{
				"tool":  1, //special tool in communication
				"metal": 5,
				"fuel":  3,
			},
			Condition: "unstable",
		},

		366: {
			ID:          366,
			Name:        "Escape Module",
			IsAvailable: false,
			StuffToFix:  map[string]int{},
			Condition:   "unstable",
		},
	}
}
