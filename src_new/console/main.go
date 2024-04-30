package main

import (
	"context"
	"fmt"
	"src_new/internal/dto"
	"src_new/internal/model"
	"src_new/internal/repository"
	"src_new/internal/service"
)

func main() {

	ctx := context.Background()

	orderrepo := repository.NewOrderRepository()
	racketrepo := repository.NewRacketRepository()

	orderserv := service.NewOrderService(orderrepo,
		racketrepo)

	_ = racketrepo.Create(ctx, &model.Racket{
		ID:       0,
		Quantity: 10,
		Price:    100,
	})

	_ = racketrepo.Create(ctx, &model.Racket{
		ID:       1,
		Quantity: 10,
		Price:    200,
	})

	_, _ = orderserv.CreateOrder(ctx, &dto.PlaceOrderReq{
		UserID: 0,
		Lines: []*dto.PlaceOrderLineReq{{
			RacketID: 0,
			Quantity: 1,
		},
			{
				RacketID: 1,
				Quantity: 2,
			},
		},
	})

	_, _ = orderserv.CreateOrder(ctx, &dto.PlaceOrderReq{
		UserID: 0,
		Lines: []*dto.PlaceOrderLineReq{{
			RacketID: 0,
			Quantity: 10,
		},
			{
				RacketID: 1,
				Quantity: 2,
			},
		},
	})

	// if err == nil {
	// 	fmt.Printf("Hehe, %f\n", order.TotalPrice)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// for _, line := range order.Lines {
	// 	fmt.Printf("%d, %d, %f\n", line.RacketID, line.Quantity, line.Price)
	// }

	orders, _ := orderserv.GetMyOrders(ctx, 0)

	for _, order := range orders {
		fmt.Printf("%d %f\n", order.ID, order.TotalPrice)

		for _, racket := range order.Lines {
			fmt.Printf("	ID %d Price %f Quantitiy %d\n", racket.RacketID, racket.Price, racket.Quantity)
		}
	}
}
