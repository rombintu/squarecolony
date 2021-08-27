package game

type Bonus struct {
	Title   string
	Message Message
}

func NewBonus(
	title string,
	message Message,
) Bonus {
	return Bonus{
		Title:   title,
		Message: message,
	}
}
