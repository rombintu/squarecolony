package game

import "time"

type Bonus struct {
	Title     string
	Data      []byte
	Timestemp time.Time
}
