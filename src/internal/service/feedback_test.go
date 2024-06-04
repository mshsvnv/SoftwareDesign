package service

import (
	"context"
	"errors"
	"src/internal/dto"
	"src/internal/model"
	"src/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type FeedbackServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.IFeedbackRepository
	service  IFeedbackService
}

func (suite *FeedbackServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewIFeedbackRepository(suite.T())
	suite.service = NewFeedbackService(nil, suite.mockRepo)
}

func TestFeedbackServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FeedbackServiceTestSuite))
}

// CreateFeedback
func (suite *FeedbackServiceTestSuite) TestCreateFeedBackFail() {

	req := &dto.CreateFeedbackReq{
		RacketID: 0,
		UserID:   1,
		Rating:   5,
	}

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(errors.New("error")).Times(1)

	feedback, err := suite.service.CreateFeedback(context.Background(), req)

	suite.Nil(feedback)
	suite.NotNil(err)
}

func (suite *FeedbackServiceTestSuite) TestCreateFeedBackSuccess() {

	req := &dto.CreateFeedbackReq{
		RacketID: 0,
		UserID:   1,
		Rating:   5,
	}

	suite.mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(nil).Times(1)

	feedback, err := suite.service.CreateFeedback(context.Background(), req)

	suite.NotNil(feedback)
	suite.Nil(err)
}

// RemoveFeedback
func (suite *FeedbackServiceTestSuite) TestRemoveFeedBackGetFail() {

	req := &dto.GetFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	suite.mockRepo.On("GetFeedback", mock.Anything, req).
		Return(nil, errors.New("error")).Times(1)

	err := suite.service.RemoveFeedback(context.Background(),
		&dto.RemoveFeedbackReq{
			RacketID: 0,
			UserID:   1,
		})

	suite.NotNil(err)
}
func (suite *FeedbackServiceTestSuite) TestRemoveFeedBackFail() {

	req := &dto.GetFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	reqRemove := &dto.RemoveFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	suite.mockRepo.On("GetFeedback", mock.Anything, req).
		Return(&model.Feedback{
			RacketID: 0,
			UserID:   1,
			Rating:   5,
		}, nil).Times(1)

	suite.mockRepo.On("Remove", mock.Anything, reqRemove).
		Return(errors.New("error"))

	err := suite.service.RemoveFeedback(context.Background(), reqRemove)

	suite.NotNil(err)
}

func (suite *FeedbackServiceTestSuite) TestRemoveFeedBackSuccess() {

	req := &dto.GetFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	reqRemove := &dto.RemoveFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	suite.mockRepo.On("GetFeedback", mock.Anything, req).
		Return(&model.Feedback{
			RacketID: 0,
			UserID:   1,
			Rating:   5,
		}, nil).Times(1)

	suite.mockRepo.On("Remove", mock.Anything, reqRemove).
		Return(nil)

	err := suite.service.RemoveFeedback(context.Background(), reqRemove)

	suite.Nil(err)
}

// UpdateFeedback
func (suite *FeedbackServiceTestSuite) TestUpdateFeedBackGetFail() {

	req := &dto.GetFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	suite.mockRepo.On("GetFeedback", mock.Anything, req).
		Return(nil, errors.New("error")).Times(1)

	err := suite.service.UpdateFeedback(context.Background(),
		&dto.UpdateFeedbackReq{
			RacketID: 0,
			UserID:   1,
		})

	suite.NotNil(err)
}

func (suite *FeedbackServiceTestSuite) TestUpdateFeedBackFail() {

	req := &dto.GetFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	reqUpdate := &dto.UpdateFeedbackReq{
		RacketID: 0,
		UserID:   1,
		Rating:   10,
	}

	suite.mockRepo.On("GetFeedback", mock.Anything, req).
		Return(&model.Feedback{
			RacketID: 0,
			UserID:   1,
			Rating:   reqUpdate.Rating,
		}, nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything, mock.Anything).
		Return(errors.New("error"))

	err := suite.service.UpdateFeedback(context.Background(), reqUpdate)

	suite.NotNil(err)
}

func (suite *FeedbackServiceTestSuite) TestUpdateFeedBackSuccess() {

	req := &dto.GetFeedbackReq{
		RacketID: 0,
		UserID:   1,
	}

	reqUpdate := &dto.UpdateFeedbackReq{
		RacketID: 0,
		UserID:   1,
		Rating:   10,
	}

	feedback := &model.Feedback{
		RacketID: 0,
		UserID:   1,
		Rating:   reqUpdate.Rating,
	}
	suite.mockRepo.On("GetFeedback", mock.Anything, req).
		Return(feedback, nil).Times(1)

	suite.mockRepo.On("Update", mock.Anything, feedback).
		Return(nil)

	err := suite.service.UpdateFeedback(context.Background(), reqUpdate)

	suite.Nil(err)
}

// GetFeedbacksByUserID
func (suite *FeedbackServiceTestSuite) TestGetFeedBackBuUserIDFail() {

	userID := 0

	suite.mockRepo.On("GetFeedbacksByUserID", mock.Anything, userID).
		Return(nil, errors.New("error"))

	feedback, err := suite.service.GetFeedbacksByUserID(context.Background(), userID)

	suite.NotNil(err)
	suite.Nil(feedback)
}

func (suite *FeedbackServiceTestSuite) TestGetFeedBackByUserIDSuccess() {

	userID := 0

	suite.mockRepo.On("GetFeedbacksByUserID", mock.Anything, userID).
		Return([]*model.Feedback{
			{
				RacketID: 0,
				UserID:   userID},
			{
				RacketID: 10,
				UserID:   userID},
		}, nil)

	feedback, err := suite.service.GetFeedbacksByUserID(context.Background(), userID)

	suite.Nil(err)
	suite.NotNil(feedback)
}

// GetFeedbacksByRacketID
func (suite *FeedbackServiceTestSuite) TestGetFeedBackBuRacketIDFail() {

	racketID := 0

	suite.mockRepo.On("GetFeedbacksByRacketID", mock.Anything, racketID).
		Return(nil, errors.New("error"))

	feedback, err := suite.service.GetFeedbacksByRacketID(context.Background(), racketID)

	suite.NotNil(err)
	suite.Nil(feedback)
}

func (suite *FeedbackServiceTestSuite) TestGetFeedBackByRacketIDSuccess() {

	racketID := 0

	suite.mockRepo.On("GetFeedbacksByRacketID", mock.Anything, racketID).
		Return([]*model.Feedback{
			{
				RacketID: racketID,
				UserID:   0},
			{
				RacketID: racketID,
				UserID:   1},
		}, nil)

	feedback, err := suite.service.GetFeedbacksByRacketID(context.Background(), racketID)

	suite.Nil(err)
	suite.NotNil(feedback)
}
