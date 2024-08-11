package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	l           logging.Interface
	authService service.IAuthService
}

func NewAuthController(
	l logging.Interface,
	authService service.IAuthService) *AuthController {
	return &AuthController{
		l:           l,
		authService: authService,
	}
}

func (a *AuthController) Login(c *gin.Context) {

	var req dto.LoginReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error 1"})
		return
	}

	token, err := a.authService.Login(c, &req)
	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error 2"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) Register(c *gin.Context) {

	var req dto.RegisterReq
	if err := c.ShouldBindBodyWithJSON(&req); c.Request.Body == nil || err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := a.authService.Register(c, &req)

	if err != nil {
		a.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
