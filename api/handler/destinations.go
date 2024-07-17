package handler

import (
	pb "api/generated/content_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary get
// @Description Gets Destinations
// @Tags destinations
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Param query query string false "query"
// @Success 200 {object} content_service.GetsDestinationsResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/destinations [get]
func (h *Handler) GetsDestinations(ctx *gin.Context) {
	h.Logger.Info("Create Destinations srated")

	req := pb.GetsDestinationsRequest{}

	limitStr := ctx.Query("limit")
	pageStr := ctx.Query("page")
	req.Query = ctx.Query("query")

	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)

	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}
	req.Page = int32(page)
	req.Limit = int32(limit)

	res, err := h.Content.GetsDestinations(ctx, &req)
	if err != nil {
		h.Logger.Error("Failed gets destinations", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
	h.Logger.Info("Gets Destinations End")
}

// @Security ApiKeyAuth
// @Summary Get
// @Description Get destination by id
// @Tags destinations
// @Param destination_id path string true "destination_id"
// @Success 200 {object} content_service.GetDestinationsResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/destinations/{destination_id} [get]
func (h *Handler) GetDestinations(ctx *gin.Context) {
	h.Logger.Info("Get Destination started")

	req := pb.GetDestinationsRequest{}
	req.Id = ctx.Param("destination_id")

	if len(req.Id) == 0 {
		h.Logger.Error("not found")
		ctx.JSON(400, gin.H{"error": "not id"})
		return
	}

	res, err := h.Content.GetDestinations(ctx, &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
	h.Logger.Info("Ended")
}
