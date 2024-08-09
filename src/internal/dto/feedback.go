package dto

import "time"

type CreateFeedbackReq struct {
	RacketID int    `json:"racket_id"`
	UserID   int    `json:"user_id"`
	Feedback string `json:"feedback"`
	Rating   int    `json:"rating"`
}

type RemoveFeedbackReq struct {
	RacketID int
	UserID   int
}

type GetFeedbackReq struct {
	RacketID int
	UserID   int
}

type UpdateFeedbackReq struct {
	RacketID int
	UserID   int
	Feedback string
	Date     time.Time
	Rating   int
}
