package usecase

import (
	"context"
	"errors"
	"fmt"

	"amartha.com/billing/model"
)

func (uc *Usecase) AcquireLoan(ctx context.Context, userID int) error {

	userInfo, err := uc.Repository.GetUserInfo(ctx, userID)
	if err != nil {
		fmt.Println(err)
	}

	if userInfo.Status == 1 {
		return errors.New("user already have loan, cannot acquire another loan")
	}

	amount := 5000000
	interest := 500000

	transactionID, err := uc.Repository.InsertTransaction(ctx, userID, amount, interest)
	if err != nil {
		fmt.Println(err)
	}

	uc.Repository.InsertLoanSchedule(ctx, userID, transactionID)
	return nil
}

func (uc *Usecase) MakePayment(ctx context.Context, payment model.MakePaymentRequest) error {
	userInfo, err := uc.Repository.GetUserInfo(ctx, payment.UserID)
	if err != nil {
		fmt.Println(err)
	}

	if userInfo.Status == 0 {
		return errors.New("No more loan to pay")
	}

	scheduleList, err := uc.Repository.GetTransactionScheduleByUserID(ctx, payment.UserID, payment.PaymentPeriod)

	if err != nil {
		fmt.Println(err)
	}

	loanPayment := 0
	transactionID := 0
	for _, data := range scheduleList {
		loanPayment += int(data.PaymentAmount)
		transactionID = data.TransactionID
	}

	fmt.Println(scheduleList)

	err = uc.Repository.UpdateTransactionByID(ctx, loanPayment, transactionID)
	fmt.Println(err)
	
	//Get 2 oldest Payment

	return nil
}
