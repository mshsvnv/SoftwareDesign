package model

import "time"

type Feedback struct {
	RacketID int
	UserID   int
	Feedback string
	Date     time.Time
	Rating   float32
}
