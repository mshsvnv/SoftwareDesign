package main

import (
	"src_new/console/handler"
	"src_new/pkg/storage/postgres"

	"github.com/rivo/tview"
)

var (
	pages = tview.NewPages()
	app   = tview.NewApplication()
	form  = tview.NewForm()
	flex  = tview.NewFlex()
)

func main() {

	// logger := logging.GetLogger()
	// logger.Info("create connection")

	db, err := postgres.New("postgresql://postgres:admin@localhost:5432/Shop")

	if err != nil {
		panic(err)
	}

	h := handler.CreateHandler(db)

	pages.AddPage("Menu (guest)", h.CreateGuestMenu(flex, form, pages, app), true, true).
		AddPage("Register", form, true, true).
		AddPage("Login", form, true, true).
		AddPage("View the catalog", flex, true, true).
		AddPage("Finish", form, true, true)

	pages.AddPage("Menu (authorized guest)", h.CreateAuthorizedGuestMenu(flex, form, pages), true, true).
		AddPage("View the catalog", flex, true, true).
		AddPage("Add racket to cart", form, true, true).
		AddPage("View my cart", flex, true, true).
		AddPage("View my orders", flex, true, true).
		AddPage("Create an order", form, true, true).
		AddPage("Create a feedback", form, true, true).
		AddPage("Exit", form, true, true)

	pages.AddPage("Menu (admin)", h.CreateAdminMenu(flex, form, pages), true, true).
		AddPage("Add racket", form, true, true).
		AddPage("Add supplier", form, true, true).
		AddPage("Remove racket", form, true, true).
		AddPage("Remove supplier", form, true, true).
		AddPage("Edit racket", form, true, true).
		AddPage("Edit supplier", form, true, true).
		AddPage("Edit order", form, true, true).
		AddPage("Exit", form, true, true)

	pages.SwitchToPage("Menu (guest)")

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
