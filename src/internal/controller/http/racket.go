package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/model"
	"src/internal/service"
	"src/pkg/logging"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RacketController struct {
	l               logging.Interface
	racketService   service.IRacketService
	feedbackService service.IFeedbackService
	userService     service.IUserService
}

func NewRacketController(
	l logging.Interface,
	racketService service.IRacketService,
	feedbackService service.IFeedbackService,
	userService service.IUserService) *RacketController {
	return &RacketController{
		l:               l,
		racketService:   racketService,
		feedbackService: feedbackService,
		userService:     userService,
	}
}

func (r *RacketController) ListsAllRackets(c *gin.Context) {

	rackets, err := r.racketService.GetAllRackets(c)
	if err != nil {
		r.l.Errorf("%s", err.Error())
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
		r.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	feedbacks, err := r.feedbackService.GetFeedbacksByRacketID(c, racketID)
	if err != nil {
		r.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"racket":    racket,
		"feedbacks": feedbacks,
	})
}

func (r *RacketController) AddRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		r.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := r.userService.GetUserByID(c, userID)

	if user.Role != model.UserRoleAdmin {
		r.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req dto.CreateRacketReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		r.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	racket, err := r.racketService.CreateRacket(c, &req)
	if err != nil {
		r.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"racket": racket,
	})
}

func (r *RacketController) UpdateRacket(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		r.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := r.userService.GetUserByID(c, userID)

	if user.Role != model.UserRoleAdmin && user.Role != model.UserRoleSeller {
		r.l.Errorf("%s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var req dto.UpdateRacketReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		r.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	racketID, _ := strconv.Atoi(c.Param("id"))
	req.ID = racketID

	err = r.racketService.UpdateRacket(c, &req)
	if err != nil {
		r.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
