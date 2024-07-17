package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) GetSchedule(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	fmt.Println("===== Get Handler ====")
	response, err := h.Usecase.GetSchedule(context.Background(), 1)
	fmt.Println("Get Schedule Being Called", response)

	e, _ := json.Marshal(response)
	w.Write(e)
	return err
}
