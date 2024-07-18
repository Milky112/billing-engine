package usecase

import (
	"context"

	"amartha.com/billing/model"
	"amartha.com/billing/repository"
)

type UsecaseInterface interface {
	GetSchedule(ctx context.Context, productID int) (model.TransactionScheduleResponse, error)
	AcquireLoan(ctx context.Context, userID int) error
}

type Usecase struct {
	Repository repository.RepositoryInterface
}

type NewUsecaseOptions struct {
	Repository repository.RepositoryInterface
}

func NewUsecase(ops NewUsecaseOptions) *Usecase {
	return &Usecase{
		Repository: ops.Repository,
	}
}
