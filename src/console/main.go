package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rivo/tview"

	"src/config"
	"src/console/handler"
	"src/pkg/logging"
	"src/pkg/storage/postgres"
)

var (
	pages = tview.NewPages()
	app   = tview.NewApplication()
	form  = tview.NewForm()
	flex  = tview.NewFlex()
)

func main() {

	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	loggerFile, err := os.OpenFile(
		cfg.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer func(loggerFile *os.File) {
		err := loggerFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(loggerFile)

	logger := logging.New(cfg.Logger.Level, loggerFile)

	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.Database,
	))

	if err != nil {
		log.Fatal(err)
	}

	h := handler.CreateHandler(db, logger)

	pages.AddPage("Menu (guest)", h.CreateGuestMenu(flex, form, pages, app), true, true).
		AddPage("Register", form, true, true).
		AddPage("Login", form, true, true).
		AddPage("View the catalog", flex, true, true).
		AddPage("Finish", form, true, true)

	pages.AddPage("Menu (authorized guest)", h.CreateAuthorizedGuestMenu(flex, form, pages), true, true).
		AddPage("View the catalog", flex, true, true).
		AddPage("Add racket to cart", form, true, true).
		AddPage("Delete racket from cart", form, true, true).
		AddPage("View my cart", flex, true, true).
		AddPage("View my orders", flex, true, true).
		AddPage("Create an order", form, true, true).
		AddPage("Create a feedback", form, true, true).
		AddPage("Exit", form, true, true)

	pages.AddPage("Menu (admin)", h.CreateAdminMenu(flex, form, pages), true, true).
		AddPage("View suppliers", flex, true, true).
		AddPage("View the catalog", flex, true, true).
		AddPage("View users", flex, true, true).
		AddPage("Add racket", form, true, true).
		AddPage("Add supplier", form, true, true).
		AddPage("Edit racket quantity", form, true, true).
		AddPage("Edit user role", form, true, true).
		AddPage("Exit", form, true, true)

	pages.AddPage("Menu (seller)", h.CreateSellerMenu(flex, form, pages), true, true).
		AddPage("View the catalog", flex, true, true).
		AddPage("View all orders", flex, true, true).
		AddPage("Add racket", form, true, true).
		AddPage("Edit racket", form, true, true).
		AddPage("Edit order", form, true, true).
		AddPage("Exit", form, true, true)

	pages.SwitchToPage("Menu (guest)")

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}
