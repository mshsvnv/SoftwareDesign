package handler

import (
	"context"

	"github.com/rivo/tview"

	"src/internal/model"
	repo "src/internal/repository/postgres"
	"src/internal/service"

	"src/pkg/logging"
	"src/pkg/storage/postgres"
	"src/pkg/utils"
)

type Handler struct {
	userService     service.UserService
	supplierService service.SupplierService
	racketService   service.RacketService
	cartService     service.CartService
	orderService    service.OrderService
	feedbackService service.FeedbackService
	logger          *logging.Logger
}

var curUser *model.User

func CreateHandler(db *postgres.Postgres, logger *logging.Logger) *Handler {

	userRepo := repo.NewUserRepository(db)
	supplierRepo := repo.NewSupplierRepository(db)
	racketRepo := repo.NewRacketRepository(db)
	cartRepo := repo.NewCartRepository(db)
	orderRepo := repo.NewOrderRepository(db)
	feedbackRepo := repo.NewFeedbackRepository(db)

	userRepo.Create(context.Background(), &model.User{
		Email:    "admin",
		Password: utils.HashAndSalt([]byte("admin")),
		Role:     model.UserRoleAdmin,
	})

	curUser = &model.User{}

	return &Handler{
		userService:     *service.NewUserService(logger, userRepo),
		supplierService: *service.NewSupplierService(logger, supplierRepo),
		racketService:   *service.NewRacketService(logger, racketRepo, supplierRepo),
		cartService:     *service.NewCartService(logger, cartRepo, racketRepo),
		orderService:    *service.NewOrderService(logger, orderRepo, cartRepo, racketRepo),
		feedbackService: *service.NewFeedbackService(logger, feedbackRepo),
		logger:          logger,
	}
}

func (h *Handler) CreateGuestMenu(flex *tview.Flex, form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {

	return tview.NewList().
		AddItem("Register", "", '-', func() {
			form.Clear(true)
			h.RegisterForm(form, pages)
			pages.SwitchToPage("Register")
		}).
		AddItem("Login", "", '-', func() {
			form.Clear(true)
			h.LoginForm(form, pages)
			pages.SwitchToPage("Login")
		}).
		AddItem("View the catalog", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewCatalogForm(flex, pages)
			pages.SwitchToPage("View the catalog")
		}).
		AddItem("Finish", "", '*', func() {
			exitFunc.Stop()
		})
}

func (h *Handler) CreateAuthorizedGuestMenu(flex *tview.Flex, form *tview.Form, pages *tview.Pages) *tview.List {

	return tview.NewList().
		AddItem("View the catalog", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewCatalogForm(flex, pages)
			pages.SwitchToPage("View the catalog")
		}).
		AddItem("Add racket to cart", "", '-', func() {
			form.Clear(true)
			h.AddRacketToCartForm(form, pages)
			pages.SwitchToPage("Add racket to cart")
		}).
		AddItem("Delete racket from cart", "", '-', func() {
			form.Clear(true)
			h.DeleteRacketFromCartForm(form, pages)
			pages.SwitchToPage("Delete racket from cart")
		}).
		AddItem("View my cart", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewMyCartForm(flex, pages)
			pages.SwitchToPage("View my cart")
		}).
		AddItem("View my orders", "", '-', func() {
			flex.Clear()
			h.ViewMyOrdersForm(flex, pages)
			pages.SwitchToPage("View my orders")
		}).
		AddItem("Create an order", "", '-', func() {
			form.Clear(true)
			h.CreateAnOrderForm(form, pages)
			pages.SwitchToPage("Create an order")
		}).
		AddItem("Create a feedback", "", '-', func() {
			form.Clear(true)
			h.CreateFeedbackForm(form, pages)
			pages.SwitchToPage("Create a feedback")
		}).
		AddItem("Exit", "", '*', func() {
			curUser = nil
			pages.SwitchToPage("Menu (guest)")
		})
}

func (h *Handler) CreateAdminMenu(flex *tview.Flex, form *tview.Form, pages *tview.Pages) *tview.List {

	return tview.NewList().
		AddItem("View suppliers", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewSuppliersForm(flex, pages)
			pages.SwitchToPage("View suppliers")
		}).
		AddItem("View the catalog", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewCatalogAllForm(flex, pages)
			pages.SwitchToPage("View the catalog")
		}).
		AddItem("View users", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewUsersForm(flex, pages)
			pages.SwitchToPage("View users")
		}).
		AddItem("Add racket", "", '-', func() {
			form.Clear(true)
			h.AddRacketForm(form, pages)
			pages.SwitchToPage("Add racket")
		}).
		AddItem("Add supplier", "", '-', func() {
			form.Clear(true)
			h.AddSupplierForm(form, pages)
			pages.SwitchToPage("Add supplier")
		}).
		AddItem("Edit racket quantity", "", '-', func() {
			form.Clear(true)
			h.EditRacketStatusForm(form, pages)
			pages.SwitchToPage("Edit racket quantity")
		}).
		AddItem("Edit user role", "", '-', func() {
			form.Clear(true)
			h.EditUserRoleForm(form, pages)
			pages.SwitchToPage("Edit user role")
		}).
		AddItem("Exit", "", '*', func() {
			curUser = nil
			pages.SwitchToPage("Menu (guest)")
		})
}

func (h *Handler) CreateSellerMenu(flex *tview.Flex, form *tview.Form, pages *tview.Pages) *tview.List {

	return tview.NewList().
		AddItem("View the catalog", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewCatalogAllForm(flex, pages)
			pages.SwitchToPage("View the catalog")
		}).
		AddItem("View all orders", "", '-', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewAllOrdersForm(flex, pages)
			pages.SwitchToPage("View all orders")
		}).
		AddItem("Add racket", "", '-', func() {
			form.Clear(true)
			h.AddRacketForm(form, pages)
			pages.SwitchToPage("Add racket")
		}).
		AddItem("Edit racket quantity", "", '-', func() {
			form.Clear(true)
			h.EditRacketStatusForm(form, pages)
			pages.SwitchToPage("Edit racket quantity")
		}).
		AddItem("Edit order status", "", '-', func() {
			form.Clear(true)
			h.EditOrderStatusForm(form, pages)
			pages.SwitchToPage("Edit order status")
		}).
		AddItem("Exit", "", '*', func() {
			curUser = nil
			pages.SwitchToPage("Menu (guest)")
		})
}
