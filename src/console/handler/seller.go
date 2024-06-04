package handler

import (
	"context"
	"src/internal/dto"
	"src/internal/model"
	"strconv"

	"github.com/rivo/tview"
)

func (h *Handler) ViewAllOrdersForm(flex *tview.Flex, pages *tview.Pages) *tview.Flex {

	flex.SetDirection(tview.FlexRow)

	orders, err := h.orderService.GetAllOrders(context.Background())

	if err != nil {
		return flex
	}

	table := tview.NewTable().SetBorders(true)
	rows := len(orders)

	for r := 0; r <= rows; r++ {
		if r == 0 {
			table.SetCell(r, 0,
				tview.NewTableCell("User ID"))

			table.SetCell(r, 1,
				tview.NewTableCell("Order ID"))

			table.SetCell(r, 2,
				tview.NewTableCell("Status"))

			table.SetCell(r, 3,
				tview.NewTableCell("Total Price"))

			table.SetCell(r, 4,
				tview.NewTableCell("Creation Date"))
		} else {

			order := orders[r-1]

			id := strconv.FormatInt(int64(order.UserID), 10)
			table.SetCell(r, 0,
				tview.NewTableCell(id))

			id = strconv.FormatInt(int64(order.ID), 10)
			table.SetCell(r, 1,
				tview.NewTableCell(id))

			table.SetCell(r, 2,
				tview.NewTableCell(string(order.Status)))

			total_price := strconv.FormatFloat(float64(order.TotalPrice), 'f', -1, 32)
			table.SetCell(r, 3,
				tview.NewTableCell(total_price))

			table.SetCell(r, 4,
				tview.NewTableCell(order.CreationDate.String()))
		}
	}

	button := tview.NewButton("Back").SetSelectedFunc(func() {

		pages.SwitchToPage("Menu (seller)")
	})

	flex.AddItem(table, 6*rows, 10, false)
	flex.AddItem(button, 1, 1, true)

	return flex
}

func (h *Handler) EditOrderStatusForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.UpdateOrder{}

	form.AddInputField("Order ID", "", 30, nil,
		func(id string) {
			d, _ := strconv.Atoi(id)
			req.OrderID = d
		})

	statuses := []string{
		"Done",
		"Cancelled"}

	form.AddDropDown("Order Status", statuses, 0,
		func(option string, optionIndex int) {
			req.Status = model.OrderStatus(option)
		})

	form.AddButton("Edit", func() {

		h.orderService.UpdateOrder(context.Background(), req)

		pages.SwitchToPage("Menu (seller)")
	})

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (seller)")
	})

	return form
}
