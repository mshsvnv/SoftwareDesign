package http

import (
	"src/internal/service"
	"src/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	handler *gin.Engine
}

func NewRouter(handler *gin.Engine) *Controller {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	return &Controller{handler: handler}
}

func (c *Controller) SetAuthRoute(l logging.Interface, authService service.IAuthService) {

	authController := NewAuthController(l, authService)

	api := c.handler.Group("auth")

	api.POST("/register", authController.Register)
	api.POST("/login", authController.Login)
}

func (c *Controller) SetRacketRoute(l logging.Interface, racketService service.IRacketService, feedbackService service.IFeedbackService) {

	racketController := NewRacketController(l, racketService, feedbackService)

	c.handler.GET("/rackets", racketController.ListsAllRackets)
	c.handler.GET("/rackets/:id", racketController.GetRacketByID)
}

func (c *Controller) SetUserRoute(
	l logging.Interface,
	cartService service.ICartService,
	authService service.IAuthService,
	userService service.IUserService,
	orderService service.IOrderService) {

	cartController := NewCartController(l, cartService)
	authController := NewAuthController(l, authService)
	userController := NewUserController(l, userService, cartService, orderService)

	c.handler.GET("/user/:id", userController.GetUserByID)

	api := c.handler.Group("api", authController.UserIdentity)

	api.GET("/profile", userController.GetMyProfile)

	api.GET("/cart", cartController.GetMyCart)
	api.POST("/cart", cartController.AddRacket)
	api.PUT("/cart/:id", cartController.UpdateRacket)
	api.DELETE("/cart/:id", cartController.RemoveRacket)

	api.GET("/orderlist", userController.GetMyOrders)
}

// order
func (c *Controller) SetOrderRoute(
	l logging.Interface,
	authService service.IAuthService,
	orderService service.IOrderService) {

	authController := NewAuthController(l, authService)
	orderController := NewOrderController(l, orderService)

	api := c.handler.Group("api", authController.UserIdentity)

	api.POST("/order", orderController.CreateOrder)
	api.GET("/orders", orderController.GetMyOrders)
}

// feedback
func (c *Controller) SetFeedbackRoute(
	l logging.Interface,
	authService service.IAuthService,
	feedbackService service.IFeedbackService) {

	authController := NewAuthController(l, authService)
	feedbackController := NewFeedbackController(l, feedbackService)

	c.handler.GET("/feedbacks/:id", feedbackController.GetFeedbacksByRacketID)

	api := c.handler.Group("api", authController.UserIdentity)

	api.GET("/feedbacks", feedbackController.GetFeedbacksByUserID)
	api.POST("/feedback", feedbackController.CreateFeedback)
	api.DELETE("/feedback/:id", feedbackController.DeleteFeedback)
}
