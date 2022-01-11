package game

import "time"

type Bonus struct {
	Title   string
	Message Message
	TTL     time.Duration
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
