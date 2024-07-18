package usecase

import (
	"context"
	"fmt"

	"amartha.com/billing/model"
)

func (uc *Usecase) GetSchedule(ctx context.Context, userID int) (model.TransactionScheduleResponse, error) {

	userInfo, err := uc.Repository.GetUserInfo(ctx, userID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userInfo)
	paymentSchedule, _ := uc.Repository.GetTransactionScheduleByUserID(ctx, userID, 50)
	fmt.Println(paymentSchedule)

	response := model.TransactionScheduleResponse{
		Schedule: paymentSchedule,
		User:     userInfo,
	}
	return response, nil
}
