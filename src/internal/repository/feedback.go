package repository

import (
	"context"
	"src/internal/dto"
	"src/internal/model"
)

//go:generate mockery --name=IFeedbackRepository
type IFeedbackRepository interface {
	Create(ctx context.Context, feedback *model.Feedback) error
	Update(ctx context.Context, feedback *model.Feedback) error
	Remove(ctx context.Context, req *dto.RemoveFeedbackReq) error
	GetFeedback(ctx context.Context, req *dto.GetFeedbackReq) (*model.Feedback, error)
	GetFeedbacksByUserID(ctx context.Context, id int) ([]*model.Feedback, error)
	GetFeedbacksByRacketID(ctx context.Context, id int) ([]*model.Feedback, error)
}
