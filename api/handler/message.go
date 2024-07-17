package handler

import (
	"api/api/auth"
	pb "api/generated/content_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Message
// @Description Send Message
// @Tags message
// @Param info body content_service.MessageRequest true "info"
// @Success 200 {object} content_service.MessageResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/messages [post]
func (h *Handler) Message(ctx *gin.Context) {
	h.Logger.Info("SendMessage started")
	accessToken := ctx.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := pb.MessageRequest{}
	err = ctx.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.SenderId = id
	res, err := h.Content.Message(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, &res)
	h.Logger.Info("SendMessage ended")
}

// @Security ApiKeyAuth
// @Summary get Message
// @Description get Message
// @Tags message
// @Param limit query string false "Number of messages to fetch"
// @Param page query string false "Number of messages to omit"
// @Success 200 {object} content_service.GetMessageResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/messages [get]
func (h *Handler) GetMessage(ctx *gin.Context) {
	h.Logger.Info("GetMessages started")
	req := pb.GetMessageRequest{}
	limitStr := ctx.Query("limit")
	pageStr := ctx.Query("page")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Logger.Error(err.Error())
			return
		}
		req.Limit = int32(limit)
	} else {
		req.Limit = 10
	}

	if pageStr != "" {
		offset, err := strconv.Atoi(pageStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Logger.Error(err.Error())
			return
		}
		req.Page = int32(offset)
	} else {
		req.Page = 1
	}

	res, err := h.Content.GetMessage(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(200, &res)
	h.Logger.Info("GetMessages ended")
}
