package handler

import (
	"context"
	"strconv"

	"github.com/rivo/tview"

	"src/internal/dto"
	"src/internal/model"
)

func (h *Handler) AddSupplierForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.CreateSupplierReq{}

	form.AddInputField("Name", "", 30, nil,
		func(name string) {
			req.Name = name
		})

	form.AddInputField("Phone", "", 30, nil,
		func(phone string) {
			req.Phone = phone
		})

	form.AddInputField("Town", "", 30, nil,
		func(town string) {
			req.Town = town
		})

	form.AddInputField("Email", "", 30, nil,
		func(email string) {
			req.Email = email
		})

	form.AddButton("Add", func() {

		_, err := h.supplierService.CreateSupplier(context.Background(), req)

		if err != nil {
			pages.SwitchToPage("Menu (admin)")
			return
		}

		pages.SwitchToPage("Menu (admin)")
	})

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) AddRacketForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.CreateRacketReq{}

	suppliers, err := h.supplierService.GetAllSuppliers(context.Background())

	if err != nil {
		pages.SwitchToPage("Menu (admin)")
		return form
	}

	if len(suppliers) != 0 {

		form.AddInputField("Brand", "", 30, nil,
			func(name string) {
				req.Brand = name
			})

		suppliersEmails := []string{}

		for _, supplier := range suppliers {
			suppliersEmails = append(suppliersEmails, supplier.Email)
		}

		form.AddDropDown("Supplier Email", suppliersEmails, 0,
			func(option string, optionIndex int) {
				req.SupplierEmail = option
			})

		form.AddInputField("Weight", "", 30, nil,
			func(weight string) {

				f, err := strconv.ParseFloat(weight, 32)

				if err != nil {
					panic(err)
				}

				req.Weight = float32(f)
			})

		form.AddInputField("Balance", "", 30, nil,
			func(balance string) {

				f, err := strconv.ParseFloat(balance, 32)

				if err != nil {
					panic(err)
				}

				req.Balance = float32(f)
			})

		form.AddInputField("HeadSize", "", 30, nil,
			func(headsize string) {

				f, err := strconv.ParseFloat(headsize, 32)

				if err != nil {
					panic(err)
				}

				req.HeadSize = float32(f)
			})

		form.AddInputField("Quantity", "", 30, nil,
			func(quantity string) {

				f, err := strconv.Atoi(quantity)

				if err != nil {
					panic(err)
				}

				req.Quantity = f
			})

		form.AddInputField("Price", "", 30, nil,
			func(price string) {

				f, err := strconv.ParseFloat(price, 32)

				if err != nil {
					panic(err)
				}

				req.Price = float32(f)
			})

		form.AddButton("Add", func() {

			_, err := h.racketService.CreateRacket(context.Background(), req)

			if err != nil {
				pages.SwitchToPage("Menu (admin)")
				return
			}

			pages.SwitchToPage("Menu (admin)")
		})
	} else {
		form.AddTextView("No racket!", "Add supplier first!", 30, 1, true, true)
	}

	form.AddButton("Back", func() {

		if curUser.Role == model.UserRoleAdmin {
			pages.SwitchToPage("Menu (admin)")
		} else {
			pages.SwitchToPage("Menu (seller)")
		}
	})

	return form
}

func (h *Handler) EditUserRoleForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := dto.UpdateRoleReq{}

	form.AddInputField("Email", "", 30, nil,
		func(email string) {
			req.Email = email
		})

	roles := []string{"Customer", "Admin", "Seller"}

	form.AddDropDown("Role", roles, 0,
		func(option string, optionIndex int) {

			if option == "Customer" {
				req.Role = model.UserRoleCustomer
			} else if option == "Seller" {
				req.Role = model.UserRoleSeller
			} else {
				req.Role = model.UserRoleAdmin
			}
		})

	form.AddButton("Edit", func() {

		h.userService.UpdateRole(context.Background(), req)

		pages.SwitchToPage("Menu (admin)")
	})

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) EditRacketStatusForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.UpdateRacketReq{}

	rackets, err := h.racketService.GetAllRackets(context.Background())

	if err != nil {
		if curUser.Role == model.UserRoleAdmin {
			pages.SwitchToPage("Menu (admin)")
		} else {
			pages.SwitchToPage("Menu (seller)")
		}
		return form
	}

	if len(rackets) != 0 {

		var ids []string

		for _, racket := range rackets {
			ids = append(ids, strconv.FormatInt(int64(racket.ID), 10))
		}

		form.AddDropDown("Racket ID", ids, 0,
			func(id string, optionIndex int) {
				d, _ := strconv.Atoi(id)
				req.ID = d
			})

		form.AddInputField("Quantity", "", 30, nil,
			func(quantity string) {
				d, _ := strconv.Atoi(quantity)
				req.Quantity = d
			})

		form.AddButton("Edit", func() {

			h.racketService.UpdateRacket(context.Background(), req)

			if curUser.Role == model.UserRoleAdmin {
				pages.SwitchToPage("Menu (admin)")
			} else {
				pages.SwitchToPage("Menu (seller)")
			}
		})
	} else {
		form.AddTextView("No racket to edit", "", 30, 1, true, true)
	}

	form.AddButton("Back", func() {

		if curUser.Role == model.UserRoleAdmin {
			pages.SwitchToPage("Menu (admin)")
		} else {
			pages.SwitchToPage("Menu (seller)")
		}
	})

	return form
}

func (h *Handler) ViewUsersForm(flex *tview.Flex, pages *tview.Pages) *tview.Flex {

	users, err := h.userService.GetAllUsers(context.Background())

	if err != nil {
		pages.SwitchToPage("Menu (admin)")
		return flex
	}

	table := tview.NewTable().SetBorders(true)
	rows := len(users)

	for r := 0; r <= rows; r++ {

		if r == 0 {
			table.SetCell(r, 0,
				tview.NewTableCell("№"))
			table.SetCell(r, 1,
				tview.NewTableCell("Email"))
			table.SetCell(r, 2,
				tview.NewTableCell("Name"))
			table.SetCell(r, 3,
				tview.NewTableCell("Surname"))
			table.SetCell(r, 4,
				tview.NewTableCell("Role"))
		} else {
			user := users[r-1]

			id := strconv.FormatInt(int64(r), 10)
			table.SetCell(r, 0,
				tview.NewTableCell(id))

			table.SetCell(r, 1,
				tview.NewTableCell(user.Email))

			table.SetCell(r, 2,
				tview.NewTableCell(user.Name))

			table.SetCell(r, 3,
				tview.NewTableCell(user.Surname))

			table.SetCell(r, 4,
				tview.NewTableCell(string(user.Role)))
		}
	}

	button := tview.NewButton("Back").SetSelectedFunc(func() {
		pages.SwitchToPage("Menu (admin)")
	})

	flex.AddItem(table, 6*rows, 10, false)
	flex.AddItem(button, 1, 1, true)

	return flex
}

func (h *Handler) ViewSuppliersForm(flex *tview.Flex, pages *tview.Pages) *tview.Flex {

	flex.SetDirection(tview.FlexRow)

	suppliers, err := h.supplierService.GetAllSuppliers(context.Background())

	if err != nil {
		return flex
	}

	table := tview.NewTable().SetBorders(true)
	rows := len(suppliers)

	for r := 0; r <= rows; r++ {

		if r == 0 {
			table.SetCell(r, 0,
				tview.NewTableCell("№"))
			table.SetCell(r, 1,
				tview.NewTableCell("Email"))
			table.SetCell(r, 2,
				tview.NewTableCell("Name"))
			table.SetCell(r, 3,
				tview.NewTableCell("Phone"))
			table.SetCell(r, 4,
				tview.NewTableCell("Town"))
		} else {
			supplier := suppliers[r-1]

			id := strconv.FormatInt(int64(r+1), 10)
			table.SetCell(r, 0,
				tview.NewTableCell(id))

			table.SetCell(r, 1,
				tview.NewTableCell(supplier.Email))

			table.SetCell(r, 2,
				tview.NewTableCell(supplier.Name))

			table.SetCell(r, 3,
				tview.NewTableCell(supplier.Phone))

			table.SetCell(r, 4,
				tview.NewTableCell(supplier.Town))
		}
	}

	button := tview.NewButton("Back").SetSelectedFunc(func() {
		pages.SwitchToPage("Menu (admin)")
	})

	flex.AddItem(table, 6*rows, 10, false)
	flex.AddItem(button, 1, 1, true)

	return flex
}

func (h *Handler) ViewCatalogAllForm(flex *tview.Flex, pages *tview.Pages) *tview.Flex {

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
				tview.NewTableCell("ID"))
			table.SetCell(r, 1,
				tview.NewTableCell("Brand"))
			table.SetCell(r, 2,
				tview.NewTableCell("Seller Email"))
			table.SetCell(r, 3,
				tview.NewTableCell("Weight"))
			table.SetCell(r, 4,
				tview.NewTableCell("Balance"))
			table.SetCell(r, 5,
				tview.NewTableCell("HeadSize"))
			table.SetCell(r, 6,
				tview.NewTableCell("Quantity"))
			table.SetCell(r, 7,
				tview.NewTableCell("Price"))
			table.SetCell(r, 8,
				tview.NewTableCell("Avaliable"))
		} else {
			racket := rackets[r-1]

			id := strconv.FormatInt(int64(racket.ID), 10)
			table.SetCell(r, 0,
				tview.NewTableCell(id))

			table.SetCell(r, 1,
				tview.NewTableCell(racket.Brand))

			table.SetCell(r, 2,
				tview.NewTableCell(racket.SupplierEmail))

			weight := strconv.FormatFloat(float64(racket.Weight), 'f', -1, 32)
			table.SetCell(r, 3,
				tview.NewTableCell(weight))

			balance := strconv.FormatFloat(float64(racket.Balance), 'f', -1, 32)
			table.SetCell(r, 4,
				tview.NewTableCell(balance))

			headSize := strconv.FormatFloat(float64(racket.HeadSize), 'f', -1, 32)
			table.SetCell(r, 5,
				tview.NewTableCell(headSize))

			quantity := strconv.FormatInt(int64(racket.Quantity), 10)
			table.SetCell(r, 6,
				tview.NewTableCell(quantity))

			price := strconv.FormatFloat(float64(racket.Price), 'f', -1, 32)
			table.SetCell(r, 7,
				tview.NewTableCell(price))

			s := "false"
			if racket.Avaliable {
				s = "true"
			}
			table.SetCell(r, 8,
				tview.NewTableCell(s))
		}
	}

	button := tview.NewButton("Back").SetSelectedFunc(func() {

		if curUser.Role == model.UserRoleAdmin {
			pages.SwitchToPage("Menu (admin)")
		} else {
			pages.SwitchToPage("Menu (seller)")
		}
	})

	flex.AddItem(table, 6*rows, 10, false)
	flex.AddItem(button, 1, 1, true)

	return flex
}
