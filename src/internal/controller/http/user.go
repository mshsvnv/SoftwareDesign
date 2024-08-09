package http

import (
	"net/http"
	"src/internal/service"
	"src/pkg/logging"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	l            logging.Interface
	userService  service.IUserService
	cartService  service.ICartService
	orderService service.IOrderService
}

func NewUserController(
	l logging.Interface,
	userService service.IUserService,
	cartService service.ICartService,
	orderService service.IOrderService) *UserController {
	return &UserController{
		l:            l,
		userService:  userService,
		cartService:  cartService,
		orderService: orderService,
	}
}

func (u *UserController) GetMyProfile(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := u.userService.GetUserByID(c, userID)
	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *UserController) GetMyOrders(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	orders, err := u.orderService.GetMyOrders(c, userID)

	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func (u *UserController) GetUserByID(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Param("id"))

	user, err := u.userService.GetUserByID(c, userID)

	if err != nil {
		u.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
