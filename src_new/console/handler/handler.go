package handler

import "github.com/rivo/tview"

type Handler struct {
}

func CreateHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateGuestMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {

	return tview.NewList().
		AddItem("Register", "", '1', func() {
			form.Clear(true)
			h.RegisterEmployeeForm(form, pages)
			pages.SwitchToPage("Register")
		}).
		AddItem("Login", "", '2', func() {
			form.Clear(true)
			h.RegisterEmployeeForm(form, pages)
			pages.SwitchToPage("Register")
		})
}

func (h *Handler) CreateAdminMenu(form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {

	return tview.NewList().
		AddItem("Register", "", '1', func() {
			form.Clear(true)
			h.RegisterEmployeeForm(form, pages)
			pages.SwitchToPage("Register")
		})
}

func (h *Handler) RegisterEmployeeForm(form *tview.Form, pages *tview.Pages) {

	form.AddButton("Back", func() {
		pages.SwitchToPage("Menu (guest)")
	})
}
