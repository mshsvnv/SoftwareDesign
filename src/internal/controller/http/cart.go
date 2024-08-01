package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"
	"strconv"

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

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	cart, err := cc.service.GetCartByID(c, userID)
	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}

func (cc *CartController) AddRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req dto.AddRacketCartReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		cc.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserID = userID

	cart, err := cc.service.AddRacket(c, &req)
	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}

func (cc *CartController) RemoveRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req dto.RemoveRacketCartReq

	// if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
	// 	cc.l.Infof(err.Error())
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	racketID, _ := strconv.Atoi(c.Param("id"))

	req.RacketID = racketID
	req.UserID = userID

	cart, err := cc.service.RemoveRacket(c, &req)
	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}

func (cc *CartController) UpdateRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req dto.UpdateRacketCartReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		cc.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	racketID, _ := strconv.Atoi(c.Param("id"))

	req.RacketID = racketID
	req.UserID = userID

	cart, err := cc.service.UpdateRacket(c, &req)
	if err != nil {
		cc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
	}

	c.JSON(http.StatusOK, gin.H{
		"cart": cart,
	})
}
