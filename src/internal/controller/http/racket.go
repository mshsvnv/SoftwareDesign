package http

import (
	"net/http"
	"src/internal/service"
	"src/pkg/logging"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RacketController struct {
	l               logging.Interface
	racketService   service.IRacketService
	feedbackService service.IFeedbackService
}

func NewRacketController(
	l logging.Interface,
	racketService service.IRacketService,
	feedbackService service.IFeedbackService) *RacketController {
	return &RacketController{
		l:               l,
		racketService:   racketService,
		feedbackService: feedbackService,
	}
}

func (r *RacketController) ListsAllRackets(c *gin.Context) {

	rackets, err := r.racketService.GetAllRackets(c)
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

	racketID, _ := strconv.Atoi(c.Param("id"))

	racket, err := r.racketService.GetRacketByID(c, racketID)
	if err != nil {
		r.l.Errorf("failed to GetRacketByID: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	feedbacks, err := r.feedbackService.GetFeedbacksByRacketID(c, racketID)
	if err != nil {
		r.l.Errorf("failed to GetFeedbacksByRacketID: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"racket":    racket,
		"feedbacks": feedbacks,
	})
}
