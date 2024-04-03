package dto

type Racket struct {
	ID       string
	Brand    string
	Quantity uint
	Price    float64
}

type ListRacketReq struct {
	Brand     string
	OrderBy   string
	OrderDesc bool
}

type CreateRacketReq struct {
	Brand    string
	Quantity uint
	Price    float64
}

type UpdateRacketReq struct {
	Brand    string
	Quantity uint
	Price    float64
}
