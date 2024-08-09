package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	l       logging.Interface
	service service.IOrderService
}

func NewOrderController(l logging.Interface, service service.IOrderService) *OrderController {
	return &OrderController{
		l:       l,
		service: service,
	}
}

func (o *OrderController) CreateOrder(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		o.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req dto.PlaceOrderReq
	if err := c.ShouldBindBodyWithJSON(&req); c.Request.Body == nil || err != nil {
		o.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserID = userID

	if err := o.service.CreateOrder(c, &req); err != nil {
		o.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error 2"})
		return
	}
}

func (o *OrderController) GetMyOrders(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		o.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	orders, err := o.service.GetMyOrders(c, userID)

	if err != nil {
		o.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}
