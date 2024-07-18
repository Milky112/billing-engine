package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"amartha.com/billing/model"
	"github.com/julienschmidt/httprouter"
)

// AcquireLoan function to acquire loan for customer with amount Rp 5.000.000 and interest 10% for 50 period of payment
func (h *Handler) AcquireLoan(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	fmt.Println("===== Get Handler ====")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Println(err)

	var t model.AcquireLoanRequest
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	userID, err := strconv.Atoi(t.UserID)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}

	err = h.Usecase.AcquireLoan(context.Background(), userID)

	return err
}

func (h *Handler) MakePayment(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	fmt.Println("===== Get Handler ====")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Println(err)

	var paymentRequest model.MakePaymentRequest
	err = json.Unmarshal(body, &paymentRequest)
	if err != nil {
		panic(err)
	}

	// err = h.Usecase.AcquireLoan(context.Background(), )

	return err
}
