package http

import (
	"net/http"
	"src/internal/service"
	"src/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	l       logging.Interface
	service service.ICartService
}

func NewCartController(l logging.Interface, service service.ICartService) *CartController {
	return &CartController{
		l:       l,
		service: service,
	}
}

func (cc *CartController) GetMyCart(c *gin.Context) {

	userID, _ := c.Get(userCtx)

	c.JSON(http.StatusOK, gin.H{
		"id": userID,
	})

	// userID, err := strconv.Atoi(c.Param("id"))

	// if err != nil {
	// 	cc.l.Infof("userID: %d %s", userID, err.Error())
	// }

	// cart, err := cc.service.GetCartByID(c, userID)
	// if err != nil {
	// 	cc.l.Errorf("failed to list fullInfoCards: %s", err.Error())
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"cart": cart,
	// })
}
