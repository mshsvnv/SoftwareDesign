package model

import "time"

type Feedback struct {
	RacketID int
	UserID   int
	Feedback string    `json:"feedback"`
	Date     time.Time `json:"date"`
	Rating   float32   `json:"rating"`
}
