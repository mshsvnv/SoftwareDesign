package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (a *AuthController) UserIdentity(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "empty auth header"})
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid auth header"})
	}

	userID, err := a.authService.ParseToken(headerParts[1])

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()})
	}

	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (int, error) {

	id, ok := c.Get(userCtx)
	if !ok {
		return 0, fmt.Errorf("%s", "userID not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, fmt.Errorf("%s", "userID is of invalid type")
	}

	return idInt, nil
}
