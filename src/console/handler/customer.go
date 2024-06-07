package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/rivo/tview"

	"src/internal/dto"
	"src/internal/model"
	"src/pkg/utils"
)

func (h *Handler) CreateAnOrderForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.PlaceOrderReq{}
	info := &model.OrderInfo{}

	cart, _ := h.cartService.GetCartByID(context.Background(), curUser.ID)

	if cart.Quantity == 0 {
		form.AddTextView("Your cart is empty!", "", 30, 1, true, true)
	} else {

		form.AddInputField("DeliveryDate", "", 30, nil,
			func(date string) {
				t, _ := time.Parse("01/02/2006", date)
				info.DeliveryDate = time.Time(t)
			})

		form.AddInputField("Address", "", 30, nil,
			func(address string) {
				info.Address = address
			})

		form.AddInputField("Recepient Name", "", 30, nil,
			func(name string) {
				info.RecepientName = name
			})

		form.AddButton("Create", func() {

			req.OrderInfo = info
			req.UserID = curUser.ID

			err := h.orderService.CreateOrder(context.Background(), req)

			if err != nil {
				form.AddTextView("Order status: ", "In progress", 20, 1, true, true)
			} else {
				form.AddTextView("Order status: ", "Failed", 20, 1, true, true)
			}

			pages.SwitchToPage("Menu (authorized guest)")
		})
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (authorized guest)")
	})

	return form
}

func (h *Handler) CreateFeedbackForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	orders, err := h.orderService.GetMyOrders(context.Background(), curUser.ID)

	if err != nil {
		return form
	}

	if len(orders) != 0 {

		var racketsID []int

		for _, order := range orders {

			for _, line := range order.Lines {
				racketsID = append(racketsID, line.RacketID)
			}
		}

		racketsID = utils.UniqueValues(racketsID)

		var racketsID_ []string
		for _, id := range racketsID {
			racketsID_ = append(racketsID_, strconv.FormatInt(int64(id), 10))
		}

		req := &dto.CreateFeedbackReq{
			UserID: curUser.ID,
		}

		form.AddDropDown("Racket ID", racketsID_, 0,
			func(option string, optionIndex int) {
				d, _ := strconv.Atoi(option)
				req.RacketID = d
			})

		form.AddTextArea("Feedback", "", 40, 0, 200,
			func(description string) {
				req.Feedback = description
			})

		form.AddInputField("Rating", "", 40, nil,
			func(option string) {
				f, _ := strconv.ParseFloat(option, 64)
				req.Rating = float32(f)
			})

		form.AddButton("Create", func() {
			h.feedbackService.CreateFeedback(context.Background(), req)

			pages.SwitchToPage("Menu (authorized guest)")
		})
	} else {
		form.AddTextView("Buy racket first-then make a feedback!", "", 30, 1, true, true)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (authorized guest)")
	})

	return form
}

func (h *Handler) ViewMyCartForm(flex *tview.Flex, pages *tview.Pages) *tview.Flex {

	flex.SetDirection(tview.FlexRow)

	cart, err := h.cartService.GetCartByID(context.Background(), curUser.ID)

	if err != nil {
		return flex
	}

	flex.AddItem(tview.NewTextView().
		SetText(fmt.Sprintf("Quantity: %d	Total Price: %f\n",
			cart.Quantity,
			cart.TotalPrice)),
		1, 0, false)

	if cart.Quantity != 0 {

		table := tview.NewTable().SetBorders(true)
		rows := len(cart.Lines)

		for r := 0; r <= rows; r++ {

			if r == 0 {
				table.SetCell(r, 0,
					tview.NewTableCell("Racket ID"))
				table.SetCell(r, 1,
					tview.NewTableCell("Quantity"))
			} else {
				curRacket := cart.Lines[r-1]
				id := strconv.FormatInt(int64(curRacket.RacketID), 10)
				table.SetCell(r, 0,
					tview.NewTableCell(id))

				quantity := strconv.FormatInt(int64(curRacket.Quantity), 10)
				table.SetCell(r, 1,
					tview.NewTableCell(quantity))
			}
		}

		flex.AddItem(table, 5*rows, 10, false)

	} else {
		flex.AddItem(tview.NewTextView().SetText("Cart is empty!"), 2, 0, false)
	}

	button := tview.NewButton("Back").SetSelectedFunc(func() {

		pages.SwitchToPage("Menu (authorized guest)")
	})

	flex.AddItem(button, 1, 1, true)

	return flex
}

func (h *Handler) ViewMyOrdersForm(flex *tview.Flex, pages *tview.Pages) *tview.Flex {

	flex.SetDirection(tview.FlexRow)

	orders, err := h.orderService.GetMyOrders(context.Background(), curUser.ID)

	if err != nil {
		return flex
	}

	if len(orders) != 0 {

		table := tview.NewTable().SetBorders(true)
		rows := len(orders)

		for r := 0; r <= rows; r++ {

			if r == 0 {
				table.SetCell(r, 0,
					tview.NewTableCell("№"))
				table.SetCell(r, 1,
					tview.NewTableCell("Total Price"))
				table.SetCell(r, 2,
					tview.NewTableCell("Creation Date"))
				table.SetCell(r, 3,
					tview.NewTableCell("Status"))
			} else {
				order := orders[r-1]

				id := strconv.FormatInt(int64(r-1), 10)
				table.SetCell(r, 0,
					tview.NewTableCell(id))

				totalPrice := strconv.FormatInt(int64(order.TotalPrice), 10)
				table.SetCell(r, 1,
					tview.NewTableCell(totalPrice))

				creationDate := order.CreationDate.String()
				table.SetCell(r, 2,
					tview.NewTableCell(creationDate))

				table.SetCell(r, 3,
					tview.NewTableCell(string(order.Status)))
			}
		}

		flex.AddItem(table, 5*rows, 10, false)

	} else {
		flex.AddItem(tview.NewTextView().SetText("There is no order!"), 2, 0, false)
	}

	button := tview.NewButton("Back").SetSelectedFunc(func() {
		pages.SwitchToPage("Menu (authorized guest)")
	})

	flex.AddItem(button, 1, 1, true)

	return flex
}

func (h *Handler) AddRacketToCartForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.AddRacketCartReq{
		UserID: curUser.ID,
	}

	rackets, err := h.racketService.GetAllAvaliableRackets(context.Background())

	if err != nil {
		pages.SwitchToPage("Menu (authorized guest)")
		return form
	}

	if len(rackets) != 0 {

		racketsID := []string{}

		for _, racket := range rackets {
			racketsID = append(racketsID, strconv.FormatInt(int64(racket.ID), 10))
		}

		form.AddDropDown("Racket ID", racketsID, 0,
			func(option string, optionIndex int) {
				req.RacketID, _ = strconv.Atoi(option)
			})

		form.AddInputField("Quantity", "", 30, nil,
			func(quantity string) {
				d, _ := strconv.Atoi(quantity)
				req.Quantity = d
			})

		form.AddButton("Add", func() {

			h.cartService.AddRacket(context.Background(), req)

			pages.SwitchToPage("Menu (authorized guest)")
		})
	} else {
		form.AddTextView("No rackets to add!", "", 20, 1, true, true)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (authorized guest)")
	})

	return form
}

func (h *Handler) DeleteRacketFromCartForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.RemoveRacketCartReq{
		UserID: curUser.ID,
	}

	cart, err := h.cartService.GetCartByID(context.Background(), curUser.ID)

	if err != nil {
		pages.SwitchToPage("Menu (authorized guest)")
		return form
	}

	racketsID := []string{}

	for _, racket := range cart.Lines {

		racketsID = append(racketsID, strconv.FormatInt(int64(racket.RacketID), 10))
	}

	if len(racketsID) != 0 {

		form.AddDropDown("Racket ID", racketsID, 0,
			func(option string, optionIndex int) {
				req.RacketID, _ = strconv.Atoi(option)
			})

		form.AddButton("Remove", func() {

			h.cartService.RemoveRacket(context.Background(), req)

			pages.SwitchToPage("Menu (authorized guest)")
		})
	} else {
		form.AddTextView("No rackets to delete!", "", 20, 1, true, true)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (authorized guest)")
	})

	return form
}
