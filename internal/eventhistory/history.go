package history

type History []string

func CreateHistory() History {
	var history []string

	return history
}

func (h *History) Add(message string) {
	*h = append(*h, message)
}
