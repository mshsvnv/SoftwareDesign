package http

import (
	"net/http"
	"src/internal/dto"
	"src/internal/service"
	"src/pkg/logging"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedbackController struct {
	l       logging.Interface
	service service.IFeedbackService
}

func NewFeedbackController(l logging.Interface, service service.IFeedbackService) *FeedbackController {
	return &FeedbackController{
		l:       l,
		service: service,
	}
}

func (fc *FeedbackController) GetFeedbacksByRacketID(c *gin.Context) {

	racketID, _ := strconv.Atoi(c.Param("id"))

	feedbacks, err := fc.service.GetFeedbacksByRacketID(c, racketID)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"feedbacks": feedbacks,
	})
}

func (fc *FeedbackController) GetFeedbacksByUserID(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	feedbacks, err := fc.service.GetFeedbacksByUserID(c, userID)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"feedbacks": feedbacks,
	})
}

func (fc *FeedbackController) CreateFeedback(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req dto.CreateFeedbackReq

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		fc.l.Infof(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	req.UserID = userID

	feedback, err := fc.service.CreateFeedback(c, &req)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"feedback": feedback,
	})
}

func (fc *FeedbackController) DeleteFeedback(c *gin.Context) {

	userID, err := getUserID(c)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	racketID, _ := strconv.Atoi(c.Param("id"))

	var req dto.RemoveFeedbackReq

	req.UserID = userID
	req.RacketID = racketID

	err = fc.service.RemoveFeedback(c, &req)

	if err != nil {
		fc.l.Errorf(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"feedback": feedback,
	// })
}
