package dto

type AddRacketCartReq struct {
	UserID   int
	RacketID int
	Quantity int
}

type RemoveRacketCartReq struct {
	UserID   int
	RacketID int
}

type UpdateRacketCartReq struct {
	UserID   int
	RacketID int
	Quantity int
}
