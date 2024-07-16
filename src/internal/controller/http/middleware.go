package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (a *UserController) UserIdentity(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
	}

	userID, err := a.authService.ParseToken(headerParts[1])
	a.l.Infof("user id is: %d", userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userID)
}
