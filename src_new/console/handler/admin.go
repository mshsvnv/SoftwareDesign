package handler

import (
	"context"
	"strconv"

	"github.com/rivo/tview"

	"src_new/internal/dto"
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
		form.AddTextView("No supplier(", "Add supplier first!", 30, 1, true, true)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) EditSupplierForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	suppliers, err := h.supplierService.GetAllSuppliers(context.Background())

	if err != nil {
		return form
	}

	if len(suppliers) != 0 {

		var emails []string
		for _, supplier := range suppliers {
			emails = append(emails, supplier.Email)
		}

		req := &dto.UpdateSupplierReq{}

		form.AddDropDown("Email", emails, 30,
			func(email string, optionIndex int) {
				req.Email = email
			})

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

		form.AddButton("Add", func() {

			err := h.supplierService.UpdateSupplier(context.Background(), req)

			if err != nil {
				pages.SwitchToPage("Menu (admin)")
				return
			}

			pages.SwitchToPage("Menu (admin)")
		})
	} else {
		form.AddTextView("No suppliers to edit!", "", 30, 1, false, false)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) EditRacketForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	req := &dto.UpdateRacketReq{}

	suppliers, err := h.supplierService.GetAllSuppliers(context.Background())

	if err != nil {
		pages.SwitchToPage("Menu (admin)")
		return form
	}

	rackets, err := h.racketService.GetAllRackets(context.Background())

	if err != nil {
		pages.SwitchToPage("Menu (admin)")
		return form
	}

	if len(suppliers) != 0 && len(rackets) != 0 {

		var ids []string

		for _, racket := range rackets {
			ids = append(ids, strconv.FormatInt(int64(racket.ID), 10))
		}

		form.AddDropDown("Racket ID", ids, 0,
			func(id string, optionIndex int) {
				d, _ := strconv.Atoi(id)
				req.ID = d
			})

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

			err := h.racketService.UpdateRacket(context.Background(), req)

			if err != nil {
				pages.SwitchToPage("Menu (admin)")
				return
			}

			pages.SwitchToPage("Menu (admin)")
		})
	} else {
		form.AddTextView("No racket to edit", "", 30, 1, true, true)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) RemoveSupplierForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	email := ""

	suppliers, err := h.supplierService.GetAllSuppliers(context.Background())

	if err != nil {
		pages.SwitchToPage("Menu (admin)")
		return form
	}

	if len(suppliers) != 0 {

		suppliersEmails := []string{}

		for _, supplier := range suppliers {
			suppliersEmails = append(suppliersEmails, supplier.Email)
		}

		form.AddDropDown("Supplier Email: ", suppliersEmails, 0,
			func(option string, optionIndex int) {
				email = option
			})

		form.AddButton("Delete", func() {
			err = h.supplierService.RemoveSupplier(context.Background(), email)

			if err != nil {
				pages.SwitchToPage("Menu (admin)")
			}
		})
	} else {
		form.AddTextView("No supplier(", "Add supplier first!", 30, 1, true, true)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}

func (h *Handler) RemoveRacketForm(form *tview.Form, pages *tview.Pages) *tview.Form {

	id := ""

	rackets, err := h.racketService.GetAllRackets(context.Background())

	if err != nil {
		pages.SwitchToPage("Menu (admin)")
		return form
	}

	if len(rackets) != 0 {

		racketsID := []string{}

		for _, racket := range rackets {
			racketsID = append(racketsID, strconv.Itoa(racket.ID))
		}

		form.AddDropDown("Racket ID: ", racketsID, 0,
			func(option string, optionIndex int) {
				id = option
			})

		form.AddButton("Delete", func() {

			dig, _ := strconv.Atoi(id)
			err = h.racketService.RemoveRacket(context.Background(), dig)

			if err != nil {
				pages.SwitchToPage("Menu (admin)")
			}
		})
	} else {
		form.AddTextView("No racket(", "Add racket first!", 30, 1, true, true)
	}

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (admin)")
	})

	return form
}
