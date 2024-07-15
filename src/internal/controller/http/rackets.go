package http

import (
	"net/http"
	"src/internal/service"
	"src/pkg/logging"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RacketController struct {
	l       logging.Interface
	service service.IRacketService
}

func NewRacketController(l logging.Interface, service service.IRacketService) *RacketController {
	return &RacketController{
		l:       l,
		service: service,
	}
}

func (r *RacketController) ListsAllRackets(c *gin.Context) {

	rackets, err := r.service.GetAllRackets(c)
	if err != nil {
		r.l.Errorf("failed to list fullInfoCards: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rackets": rackets,
	})
}

func (r *RacketController) GetRacketByID(c *gin.Context) {

	// cacheKey := c.Request.URL.RequestURI()
	racketID, _ := strconv.Atoi(c.Param("id"))

	racket, err := r.service.GetRacketByID(c, racketID)
	if err != nil {
		r.l.Errorf("failed to list fullInfoCards: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to list rackets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"racket": racket,
	})
}
