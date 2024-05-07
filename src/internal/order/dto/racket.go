package dto

type Racket struct {
	ID    string
	Brand string
	Price float64
}

type UpdateRacketReq struct {
	ID       string
	Quantity uint
}
