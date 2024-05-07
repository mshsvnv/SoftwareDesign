package dto

type CreateRacketReq struct {
	Brand    string
	Weight   string
	Balance  float32
	HeadSize float32
	Quantity int
	Price    int
}

type UpdateRacketReq struct {
	ID       int
	Brand    string
	Weight   string
	Balance  float32
	HeadSize float32
	Quantity int
	Price    int
}