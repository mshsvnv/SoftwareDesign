package dto

type AddRacketCartReq struct {
	UserID   int
	RacketID int `json:"racket_id"`
	Quantity int `json:"quantity"`
}

type RemoveRacketCartReq struct {
	UserID   int `json:"user_id"`
	RacketID int `json:"racket_id"`
}

type UpdateRacketCartReq struct {
	UserID   int `json:"user_id"`
	RacketID int `json:"racket_id"`
	Quantity int `json:"quantity"`
}
