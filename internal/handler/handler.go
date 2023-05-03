package handler

import (
	"net/http"

	"forum/internal/models"
	"forum/internal/render"
	"forum/internal/service"
	"forum/pkg/logger"
)

type Handler struct {
	service       *service.Service
	templateCache render.TemplateCache
	logger        *logger.Logger
}

func NewHandler(service *service.Service, tempCache render.TemplateCache) *Handler {
	return &Handler{
		service:       service,
		templateCache: tempCache,
		logger:        logger.GetLoggerInstance(),
	}
}

func (h *Handler) getUserFromContext(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		h.logger.PrintInfo("User is not authenticated")
		return nil
	}
	return user
}
