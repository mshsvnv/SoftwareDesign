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

type DeleteRacketCartReq struct {
	RacketID string
	OrderID  string
}

type DeleteRacketOrderReq struct {
	OrderID string
	CartID  string
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
