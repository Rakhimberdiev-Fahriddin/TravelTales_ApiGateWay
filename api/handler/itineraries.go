package handler

import (
	"api/api/auth"
	pb "api/generated/content_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary create
// @Description create itineraries
// @Tags itineraries
// @Param info body content_service.CreateItinerariesRequest true "info"
// @Success 200 {object} content_service.CreateItinerariesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries [post]
func (h *Handler) CreateItineraries(ctx *gin.Context) {
	h.Logger.Info("Create Itineraries started")

	accesToken := ctx.GetHeader("Authorization")

	id, err := auth.GetUserIdFromAccessToken(accesToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(400, err.Error())
		return
	}

	req := pb.CreateItinerariesRequest{}
	err = ctx.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.AuthorId = id

	res, err := h.Content.CreateItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error("there")
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, res)
	h.Logger.Info("Ended")
}

// @Security ApiKeyAuth
// @Summary update
// @Description update itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Param info body content_service.UpdateItinerariesRequest true "info"
// @Success 200 {object} content_service.UpdateItinerariesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [put]
func (h *Handler) UpdateItineraries(ctx *gin.Context) {
	req := pb.UpdateItinerariesRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(400, err.Error())
		return
	}

	req.Id = ctx.Param("itinerary_id")
	if len(req.Id) <= 0 {

		h.Logger.Error("id is empty")
		ctx.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.Content.UpdateItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("UpdateItineraries ended")
}

// @Security ApiKeyAuth
// @Summary Delete
// @Description Delete itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Success 200 {object} string "successfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [delete]
func (h *Handler) DeleteItineraries(ctx *gin.Context) {
	h.Logger.Info("DeleteItineraries started")
	req := pb.DeleteItinerariesRequest{}
	req.Id = ctx.Param("itinerary_id")
	if len(req.Id) <= 0 {
		h.Logger.Error("no id")
		ctx.JSON(400, gin.H{"error": "id is empty"})
	}
	_, err := h.Content.DeleteItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "successfully deleted"})
	h.Logger.Info("DeleteItineraries ended")
}

// GetItineraries godoc
// @Security ApiKeyAuth
// @Summary Get
// @Description Gets Itineraries
// @Tags itineraries
// @Param limit query string false "Number of itineraries to fetch"
// @Param page query string false "Number of itineraries to omit"
// @Success 200 {object} content_service.GetsItinerariesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries [get]
func (h *Handler) GetsItineraries(ctx *gin.Context) {
	h.Logger.Info("GetItineraries started")
	req := pb.GetsItinerariesRequest{}
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

	res, err := h.Content.GetsItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("GetItineraries ended")
}

// @Security ApiKeyAuth
// @Summary Get
// @Description Get itineraries by id
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Success 200 {object} content_service.GetItinerariesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [get]
func (h *Handler) GetItineraries(ctx *gin.Context) {
	h.Logger.Info("GetItinerariesById started")
	req := pb.GetItinerariesRequest{}
	req.Id = ctx.Param("itinerary_id")
	if len(req.Id) <= 0 {
		h.Logger.Error("id is empty")
		ctx.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.Content.GetItineraries(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("GetItinerariesById ended")
}
