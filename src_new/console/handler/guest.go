package handler

import (
	"context"
	"strconv"

	"github.com/rivo/tview"

	"src_new/internal/dto"
	"src_new/internal/model"
)

func (h *Handler) RegisterForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.RegisterReq{}

	form.AddInputField("Name", "", 30, nil,
		func(name string) {
			req.Name = name
		})

	form.AddInputField("Surname", "", 30, nil,
		func(surname string) {
			req.Surname = surname
		})

	form.AddInputField("Email", "", 30, nil,
		func(email string) {
			req.Email = email
		})

	form.AddInputField("Password", "", 30, nil,
		func(password string) {
			req.Password = password
		})

	states := []string{"Customer", "Admin"}

	form.AddDropDown("Role", states, 0,
		func(option string, optionIndex int) {
			req.Role = option
		})

	form.AddButton("Register", func() {

		_, err := h.userService.Register(context.Background(), req)

		if err != nil {
			pages.SwitchToPage("Menu (guest)")
			return
		}

		pages.SwitchToPage("Menu (guest)")
	})

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (guest)")
	})

	return form
}

func (h *Handler) LoginForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.LoginReq{}

	form.AddInputField("Email", "", 30, nil,
		func(email string) {
			req.Email = email
		})

	form.AddInputField("Password", "", 30, nil,
		func(password string) {
			req.Password = password
		})

	form.AddButton("Login", func() {

		user, err := h.userService.Login(context.Background(), req)

		if err != nil {
			pages.SwitchToPage("Menu (guest)")
			return
		}

		if user.Role == "Admin" {
			pages.SwitchToPage("Menu (admin)")
			return
		}

		curUser = user
		pages.SwitchToPage("Menu (authorized guest)")
	})

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (guest)")
	})

	return form
}

func (h *Handler) ViewCatalogForm(flex *tview.Flex, pages *tview.Pages) *tview.Flex {

	flex.SetDirection(tview.FlexRow)

	rackets, err := h.racketService.GetAllRackets(context.Background())

	if err != nil {
		return flex
	}

	table := tview.NewTable().SetBorders(true)
	rows := len(rackets)

	for r := 0; r <= rows; r++ {

		if r == 0 {
			table.SetCell(r, 0,
				tview.NewTableCell("Num"))
			table.SetCell(r, 1,
				tview.NewTableCell("Brand"))
			table.SetCell(r, 2,
				tview.NewTableCell("Weight"))
			table.SetCell(r, 3,
				tview.NewTableCell("Balance"))
			table.SetCell(r, 4,
				tview.NewTableCell("HeadSize"))
			table.SetCell(r, 5,
				tview.NewTableCell("Price"))
		} else {
			curR := r - 1
			id := strconv.FormatInt(int64(rackets[curR].ID), 10)
			table.SetCell(r, 0,
				tview.NewTableCell(id))

			table.SetCell(r, 1,
				tview.NewTableCell(rackets[curR].Brand))

			weight := strconv.FormatFloat(float64(rackets[curR].Weight), 'f', -1, 32)
			table.SetCell(r, 2,
				tview.NewTableCell(weight))

			balance := strconv.FormatFloat(float64(rackets[curR].Balance), 'f', -1, 32)
			table.SetCell(r, 3,
				tview.NewTableCell(balance))

			headSize := strconv.FormatFloat(float64(rackets[curR].HeadSize), 'f', -1, 32)
			table.SetCell(r, 4,
				tview.NewTableCell(headSize))

			price := strconv.FormatFloat(float64(rackets[curR].Price), 'f', -1, 32)
			table.SetCell(r, 5,
				tview.NewTableCell(price))
		}
	}

	button := tview.NewButton("Back").SetSelectedFunc(func() {

		if curUser != nil {

			if curUser.Role == model.UserRoleAdmin {
				pages.SwitchToPage("Menu (admin)")
			} else if curUser.Role == model.UserRoleCustomer {
				pages.SwitchToPage("Menu (authorized guest)")
			}

		} else {
			pages.SwitchToPage("Menu (guest)")
		}
	})

	flex.AddItem(table, 6*rows, 10, false)
	flex.AddItem(button, 1, 1, true)

	return flex
}
