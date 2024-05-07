package service

import (
	"context"
	"fmt"

	"src_new/internal/dto"
	"src_new/internal/model"
	repo "src_new/internal/repository"
	"src_new/pkg/utils"
)

//go:generate mockery --name=IFeedbackService
type IFeedbackService interface {
	CreateFeedback(ctx context.Context, req *dto.CreateFeedbackReq) (*model.Feedback, error)
	RemoveFeedback(ctx context.Context, req *dto.RemoveFeedbackReq) error
	UpdateFeedback(ctx context.Context, req *dto.UpdateFeedbackReq) error
	GetFeedbacksByUserID(ctx context.Context, userID int) ([]*model.Feedback, error)
	GetFeedbacksByRacketID(ctx context.Context, racketID int) ([]*model.Feedback, error)
}

type FeedbackService struct {
	repo repo.IFeedbackRepository
}

func NewFeedbackService(repo repo.IFeedbackRepository) *FeedbackService {
	return &FeedbackService{
		repo: repo,
	}
}

func (s *FeedbackService) CreateFeedback(ctx context.Context, req *dto.CreateFeedbackReq) (*model.Feedback, error) {

	var feedback model.Feedback

	utils.Copy(&feedback, req)

	err := s.repo.Create(ctx, &feedback)

	if err != nil {
		return nil, fmt.Errorf("CreateFeedback.GetFeedbackByID fail, %s", err)
	}

	return &feedback, nil

}

func (s *FeedbackService) RemoveFeedback(ctx context.Context, req *dto.RemoveFeedbackReq) error {

	feedback, err := s.repo.GetFeedback(ctx,
		&dto.GetFeedbackReq{
			RacketID: req.RacketID,
			UserID:   req.UserID,
		})

	if feedback == nil {
		return fmt.Errorf("RemoveFeedback.GetFeedbackByID fail, %s", err)
	}

	err = s.repo.Remove(ctx, req)

	if err != nil {
		return fmt.Errorf("RemoveFeedback.Remove fail, %s", err)
	}

	return nil
}

func (s *FeedbackService) UpdateFeedback(ctx context.Context, req *dto.UpdateFeedbackReq) error {

	feedback, err := s.repo.GetFeedback(ctx,
		&dto.GetFeedbackReq{
			RacketID: req.RacketID,
			UserID:   req.UserID,
		})

	if feedback == nil {
		return fmt.Errorf("UpdateFeedback.GetFeedbackByID fail, %s", err)
	}

	utils.Copy(&feedback, req)

	err = s.repo.Update(ctx, feedback)

	if err != nil {
		return fmt.Errorf("UpdateFeedback.Update fail, %s", err)
	}

	return nil
}

func (s *FeedbackService) GetFeedbacksByRacketID(ctx context.Context, racketID int) ([]*model.Feedback, error) {

	feedbacks, err := s.repo.GetFeedbacksByRacketID(ctx, racketID)

	if err != nil {
		return nil, fmt.Errorf("GetFeedbacksByRacketID.GetFeedbacksByRacketID fail, %s", err)
	}

	return feedbacks, nil
}

func (s *FeedbackService) GetFeedbacksByUserID(ctx context.Context, userID int) ([]*model.Feedback, error) {

	feedbacks, err := s.repo.GetFeedbacksByUserID(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("GetFeedbacksByUserID.GetFeedbacksByUserID fail, %s", err)
	}

	return feedbacks, nil
}
