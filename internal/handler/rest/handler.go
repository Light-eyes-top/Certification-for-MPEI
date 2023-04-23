package rest

import (
	"certification/internal/service"
)

type Handler struct {
	sc *service.Service
}

func NewHandler(sc *service.Service) *Handler {
	return &Handler{sc: sc}
}
