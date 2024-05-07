package repository

import (
	"context"
	"src_new/internal/dto"
	"src_new/internal/model"
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

// type FeedbackRepository struct {
// 	db []*model.Feedback
// }

// func NewFeedbackRepository() *FeedbackRepository {
// 	return &FeedbackRepository{}
// }

// func (r *FeedbackRepository) Create(ctx context.Context, feedback *model.Feedback) error {

// 	r.db = append(r.db, feedback)

// 	return nil
// }

// func (r *FeedbackRepository) Update(ctx context.Context, feedback *model.Feedback) error {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].UserID == feedback.UserID && r.db[i].RacketID == feedback.RacketID {

// 			r.db[i] = feedback

// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("no feedback")
// }

// func (r *FeedbackRepository) Remove(ctx context.Context, req *dto.RemoveFeedbackReq) error {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].UserID == req.UserID && r.db[i].RacketID == req.RacketID {

// 			r.db = append(r.db[:i], r.db[i+1:]...)
// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("no feedback")
// }

// func (r *FeedbackRepository) GetFeedback(ctx context.Context, req *dto.GetFeedbackReq) (*model.Feedback, error) {

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].UserID == req.UserID && r.db[i].RacketID == req.RacketID {

// 			return r.db[i], nil
// 		}
// 	}

// 	return nil, fmt.Errorf("no feedback")
// }

// func (r *FeedbackRepository) GetFeedbackByRacketID(ctx context.Context, racketID int) ([]*model.Feedback, error) {

// 	var feedbacks []*model.Feedback

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].RacketID == racketID {
// 			feedbacks = append(feedbacks, r.db[i], nil)
// 		}
// 	}

// 	return feedbacks, nil
// }

// func (r *FeedbackRepository) GetFeedbackByUserID(ctx context.Context, userID int) ([]*model.Feedback, error) {

// 	var feedbacks []*model.Feedback

// 	for i := 0; i < len(r.db); i++ {

// 		if r.db[i].RacketID == userID {
// 			feedbacks = append(feedbacks, r.db[i], nil)
// 		}
// 	}

// 	return feedbacks, nil
// }
