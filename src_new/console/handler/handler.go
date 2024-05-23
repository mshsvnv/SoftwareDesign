package handler

import (
	"context"

	"github.com/rivo/tview"

	"src_new/internal/model"
	repo "src_new/internal/repository/postgres"
	"src_new/internal/service"

	"src_new/pkg/logging"
	"src_new/pkg/storage/postgres"
	"src_new/pkg/utils"
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

	logger.Infof("here i am")

	userRepo.Create(context.Background(), &model.User{
		Email:    "admin@mail.ru",
		Password: utils.HashAndSalt([]byte("admin")),
		Role:     model.UserRoleAdmin,
	})

	curUser = &model.User{}
	// curUser.Role = model.UserRoleAdmin

	return &Handler{
		userService:     *service.NewUserService(userRepo),
		supplierService: *service.NewSupplierService(supplierRepo),
		racketService:   *service.NewRacketService(racketRepo, supplierRepo),
		cartService:     *service.NewCartService(cartRepo, racketRepo),
		orderService:    *service.NewOrderService(orderRepo, cartRepo),
		feedbackService: *service.NewFeedbackService(feedbackRepo),
		logger:          logger,
	}
}

func (h *Handler) CreateGuestMenu(flex *tview.Flex, form *tview.Form, pages *tview.Pages, exitFunc *tview.Application) *tview.List {

	return tview.NewList().
		AddItem("Register", "", '1', func() {
			form.Clear(true)
			h.RegisterForm(form, pages)
			pages.SwitchToPage("Register")
		}).
		AddItem("Login", "", '2', func() {
			form.Clear(true)
			h.LoginForm(form, pages)
			pages.SwitchToPage("Login")
		}).
		AddItem("View the catalog", "", '3', func() {
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
		AddItem("View the catalog", "", '1', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewCatalogForm(flex, pages)
			pages.SwitchToPage("View the catalog")
		}).
		AddItem("Add racket to cart", "", '2', func() {
			form.Clear(true)
			h.AddRacketToCartForm(form, pages)
			pages.SwitchToPage("Add racket to cart")
		}).
		AddItem("View my cart", "", '3', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewMyCartForm(flex, pages)
			pages.SwitchToPage("View my cart")
		}).
		AddItem("View my orders", "", '4', func() {
			flex.Clear()
			h.ViewMyOrdersForm(flex, pages)
			pages.SwitchToPage("View my orders")
		}).
		AddItem("Create an order", "", '5', func() {
			form.Clear(true)
			h.CreateAnOrderForm(form, pages)
			pages.SwitchToPage("Create an order")
		}).
		AddItem("Create a feedback", "", '6', func() {
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
		AddItem("View suppliers", "", '1', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewSuppliersForm(flex, pages)
			pages.SwitchToPage("View suppliers")
		}).
		AddItem("View the catalog", "", '2', func() {
			form.Clear(true)
			flex.Clear()
			h.ViewCatalogForm(flex, pages)
			pages.SwitchToPage("View the catalog")
		}).
		AddItem("Add racket", "", '3', func() {
			form.Clear(true)
			h.AddRacketForm(form, pages)
			pages.SwitchToPage("Add racket")
		}).
		AddItem("Add supplier", "", '4', func() {
			form.Clear(true)
			h.AddSupplierForm(form, pages)
			pages.SwitchToPage("Add supplier")
		}).
		AddItem("Remove racket", "", '5', func() {
			form.Clear(true)
			h.RemoveRacketForm(form, pages)
			pages.SwitchToPage("Remove racket")
		}).
		AddItem("Remove supplier", "", '6', func() {
			form.Clear(true)
			h.RemoveSupplierForm(form, pages)
			pages.SwitchToPage("Remove supplier")
		}).
		AddItem("Edit racket", "", '7', func() {
			form.Clear(true)
			h.EditRacketForm(form, pages)
			pages.SwitchToPage("Remove racket")
		}).
		AddItem("Edit supplier", "", '8', func() {
			form.Clear(true)
			h.EditSupplierForm(form, pages)
			pages.SwitchToPage("Remove supplier")
		}).
		AddItem("Exit", "", '*', func() {
			curUser = nil
			pages.SwitchToPage("Menu (guest)")
		})
}
