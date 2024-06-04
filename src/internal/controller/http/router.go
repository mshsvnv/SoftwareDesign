package http

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	httputils "course/internal/controller/http/utils"
// )

// type Controller struct {
// 	handler *gin.Engine
// }

// func NewRouter(handler *gin.Engine) *Controller {
// 	handler.Use(gin.Logger())
// 	handler.Use(gin.Recovery())

// 	handler.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

// 	handler.OPTIONS("/*any", httputils.DisableCors)

// 	return &Controller{handler: handler}
// }
