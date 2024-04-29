package main

import (
	"context"
	"fmt"
	"src_new/internal/dto"
	"src_new/internal/repository"
	"src_new/internal/service"
)

func main() {
	repo := repository.NewUserRepository()
	serv := service.NewUserService(repo)

	ctx := context.Background()

	_, _ = serv.Register(ctx, &dto.RegisterReq{
		Name:     "Masha",
		Surname:  "Savinova",
		Email:    "mail@.ru",
		Password: "123456",
	})

	repoCart := repository.NewCartRepository()
	servCart := service.NewCartService(repoCart)

	_, err := servCart.GetCartByUserID(ctx, 0)

	if err != nil {
		fmt.Println(err.Error())
	}

	cart, _ := servCart.AddRacket(ctx, &dto.AddRacketReq{
		UserID:   0,
		RacketID: 1,
		Quantity: 2,
	})

	for _, line := range cart.Lines {
		fmt.Printf("Racket ID: %d Quantity: %d\n", line.RacketID, line.Quantity)
	}

	cart, _ = servCart.AddRacket(ctx, &dto.AddRacketReq{
		UserID:   0,
		RacketID: 2,
		Quantity: 2,
	})

	// cart, err := servCart.GetCartByUserID(ctx, 0)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	for _, line := range cart.Lines {
		fmt.Printf("Racket ID: %d Quantity: %d\n", line.RacketID, line.Quantity)
	}
}
