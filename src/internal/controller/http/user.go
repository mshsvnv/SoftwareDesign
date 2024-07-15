package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	l       logging.Interface
	service service.IUserService
}

func NewUserController(l logging.Interface, service service.IUserService) *UserController {
	return &UserController{
		l:       l,
		service: service,
	}
}

func (u *UserController) Login(c *gin.Context) {

	var req dto.LoginReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		u.l.Infof("error")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	_, err := u.service.Login(c, &req)
	if err != nil {
		u.l.Infof("error")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "you are logined",
	})
}

func (u *UserController) Register(c *gin.Context) {

	var req dto.RegisterReq
	if err := c.ShouldBindBodyWithJSON(&req); c.Request.Body == nil || err != nil {
		u.l.Infof("error")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	_, err := u.service.Register(c, &req)

	if err != nil {
		u.l.Infof("error")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "you are registered",
	})
}

func (u *UserController) RefreshToken(c *gin.Context) {

}
