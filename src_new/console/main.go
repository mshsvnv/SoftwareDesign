package main

// import (
// 	"github.com/rivo/tview"

// 	"src_new/console/handler"
// )

// // const connURL = "postgresql://postgres:m20031504@localhost:5432/Shop"

// // var testDB *postgres.Postgres

// var (
// 	pages = tview.NewPages()
// 	app   = tview.NewApplication()
// 	form  = tview.NewForm()
// 	// list  = tview.NewList().ShowSecondaryText(true)
// )

// func main() {

// 	h := handler.CreateHandler()

// 	pages.AddPage("Menu (guest)", h.CreateGuestMenu(form, pages, app), true, true).
// 		AddPage("keke", form, true, true).
// 		AddPage("Login", form, true, true)

// 	pages.AddPage("Menu (admin)", h.CreateAdminMenu(form, pages, app), true, true).
// 		AddPage("Register", form, true, true).
// 		AddPage("Login", form, true, true)

// 	pages.SwitchToPage("Menu (guest)")

//		if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
//			panic(err)
//		}
//	}
