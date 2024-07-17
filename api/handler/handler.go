package handler

import (
	"api/generated/auth_service"
	"api/generated/content_service"
	"log/slog"
)

type Handler struct {
	Logger  *slog.Logger
	Content content_service.ContentClient
	User auth_service.AuthServiceClient
}

func NewHandler(user auth_service.AuthServiceClient,contentClient content_service.ContentClient, log *slog.Logger) *Handler {
	return &Handler{
		User: user,
		Content: contentClient,
		Logger:  log,
	}
}
