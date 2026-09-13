package history

type History []string

func Create() History {
	var history []string

	return history
}

func (h *History) Add(message string) {
	*h = append(*h, message)
}
