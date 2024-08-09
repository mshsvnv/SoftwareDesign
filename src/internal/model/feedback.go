package model

import "time"

type Feedback struct {
	RacketID int       `json:"racket_id"`
	UserID   int       `json:"user_id"`
	Feedback string    `json:"feedback"`
	Date     time.Time `json:"date"`
	Rating   int       `json:"rating"`
}
