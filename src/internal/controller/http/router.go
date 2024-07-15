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
func (c *Controller) SetUserRoute(l logging.Interface, service service.IUserService) {

	a := NewUserController(l, service)

	c.handler.POST("/register", a.Register)
	c.handler.POST("/login", a.Login)
	c.handler.POST("/refresh", a.RefreshToken)
	c.handler.GET("/me")
}

// order
func (c *Controller) SetOrderRoute(l logging.Interface) {

}

// product
func (c *Controller) SetProductRoute(l logging.Interface, service service.IRacketService) {

	a := NewRacketController(l, service)

	c.handler.GET("/rackets", a.ListsAllRackets)
	c.handler.GET("/rackets/:id", a.GetRacketByID)
}

// cart
