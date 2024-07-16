package http

import (
	"github.com/gin-gonic/gin"
	// httputils "course/internal/controller/http/utils"
	"src/internal/service"
	"src/pkg/logging"
)

type Controller struct {
	handler *gin.Engine
}

func NewRouter(handler *gin.Engine) *Controller {

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	// handler.OPTIONS("/*any", httputils.DisableCors)

	return &Controller{handler: handler}
}

// user
func (c *Controller) SetUserRoute(l logging.Interface, service service.IUserService, authService service.IAuthService) {

	a := NewUserController(l, service, authService)

	group := c.handler.Group("auth")

	group.POST("/register", a.Register)
	group.POST("/login", a.Login)
}

// product
func (c *Controller) SetProductRoute(l logging.Interface, service service.IRacketService) {

	a := NewRacketController(l, service)

	c.handler.GET("/rackets", a.ListsAllRackets)
	c.handler.GET("/rackets/:id", a.GetRacketByID)
}

// cart
func (c *Controller) SetCartRoute(l logging.Interface, service service.ICartService, userService service.IUserService, authService service.IAuthService) {

	a := NewCartController(l, service)
	b := NewUserController(l, userService, authService)

	group := c.handler.Group("api", b.UserIdentity)

	group.GET("/cart", a.GetMyCart)
}

// order
func (c *Controller) SetOrderRoute(l logging.Interface) {

}
