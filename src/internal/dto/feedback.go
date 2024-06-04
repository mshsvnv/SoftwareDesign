package dto

import "time"

type CreateFeedbackReq struct {
	RacketID int
	UserID   int
	Feedback string
	Date     time.Time
	Rating   float32
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
	Rating   float32
}
