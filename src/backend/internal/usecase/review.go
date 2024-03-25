package usecase

import (
	"context"

	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/Roongkun/software-eng-ii/internal/repository"
	"github.com/Roongkun/software-eng-ii/internal/repository/postgres"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReviewUseCase struct {
	ReviewRepo repository.Review
}

func NewReviewUseCase(db *bun.DB) *ReviewUseCase {
	return &ReviewUseCase{
		ReviewRepo: postgres.NewReviewDB(db),
	}
}

func (p *ReviewUseCase) FindByUserId(ctx context.Context, userId uuid.UUID) ([]*model.Review, error) {
	return p.ReviewRepo.FindByUserId(ctx, userId)
}
