package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetSchedule(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	fmt.Println("===== Get Handler ====")
	userIDStr := r.FormValue("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("===== Invalid User ID ====")

	}
	response, err := h.Usecase.GetSchedule(context.Background(), userID)
	fmt.Println("Get Schedule Being Called", response)

	e, _ := json.Marshal(response)
	w.Write(e)
	return err
}
