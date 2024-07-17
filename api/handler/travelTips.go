package handler

import (
	"api/api/auth"
	pb "api/generated/content_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Add Travel Tips
// @Description create tips
// @Tags tips
// @Param info body content_service.AddTravelTipsRequest true "destination_id"
// @Success 200 {object} content_service.AddTravelTipsResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/travel-tips [post]
func (h *Handler) AddTravelTips(ctx *gin.Context) {
	h.Logger.Info("CreateTips started")
	accessToken := ctx.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := pb.AddTravelTipsRequest{}
	err = ctx.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	req.AuthorId = id
	res, err := h.Content.AddTravelTips(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("CreateTips ended")
}

// @Security ApiKeyAuth
// @Summary get
// @Description get tips
// @Tags tips
// @Param limit query string false "Number of messages to fetch"
// @Param page query string false "Number of messages to omit"
// @Param category query string false "category"
// @Success 200 {object} content_service.GetTravelTipsResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/travel-tips [get]
func (h *Handler) GetTravelTips(ctx *gin.Context) {
	h.Logger.Info("GetTips started")

	req := pb.GetTravelTipsRequest{}
	req.Category = ctx.Query("category")
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
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Logger.Error(err.Error())
			return
		}
		req.Page = int32(page)
	} else {
		req.Page = 1
	}

	res, err := h.Content.GetTravelTips(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("GetTips ended")
}

// @Security ApiKeyAuth
// @Summary best
// @Description get user
// @Tags users
// @Param user_id path string true "user_id"
// @Success 200 {object} content_service.StatisticsUserResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/users/{user_id}/statistics [get]
func (h *Handler) StatisticsUser(ctx *gin.Context) {
	h.Logger.Info("GetUserStat started")
	req := pb.StatisticsUserRequest{}
	req.UserId = ctx.Param("user_id")
	if len(req.UserId) <= 0 {
		h.Logger.Error("id is empty")
		ctx.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.Content.StatisticsUser(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("GetUserStat ended")
}

// @Security ApiKeyAuth
// @Summary top places
// @Description get top places
// @Tags top
// @Success 200 {object} content_service.GetTrendingDestinationsResponce
// @Failure 500 {object} string "Server error"
// @Router /api/v1/trending-destinations [get]
func (h *Handler) GetTrendingDestinations(ctx *gin.Context) {
	h.Logger.Info("TopDestinations started")
	req := pb.GetTrendingDestinationsRequest{}
	res, err := h.Content.GetTrendingDestinations(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	fmt.Println(res)
	ctx.JSON(200, &res)
	h.Logger.Info("TopDestinations ended")
}
