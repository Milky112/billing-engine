package handler

import "amartha.com/billing/usecase"

type Handler struct {
	Usecase usecase.UsecaseInterface
}

type NewServerOptions struct {
	Usecase usecase.UsecaseInterface
}

func NewServer(opts NewServerOptions) *Handler {
	return &Handler{
		Usecase: opts.Usecase,
	}
}
