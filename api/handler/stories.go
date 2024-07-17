package handler

import (
	"api/api/auth"
	pb "api/generated/content_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary create stories
// @Description create new stories
// @Tags stories
// @Param info body content_service.CreateStoriesRequest true "info"
// @Success 200 {object} content_service.CreateStoriesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories [post]
func (h *Handler) CreateStories(ctx *gin.Context) {
	h.Logger.Info("CreateStory started")
	req := pb.CreateStoriesRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	accessToken := ctx.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.Content.CreateStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("CreateStory ended")
}

// @Security ApiKeyAuth
// @Summary Update stories
// @Description Update new stories
// @Tags stories
// @Param story_id path string true "story_id"
// @Param info body content_service.UpdateStoriesRequest true "info"
// @Success 200 {object} content_service.UpdateStoriesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [put]
func (h *Handler) UpdateStories(ctx *gin.Context) {
	h.Logger.Info("UpdateStory started")
	req := pb.UpdateStoriesRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.Id = ctx.Param("story_id")
	res, err := h.Content.UpdateStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("UpdateStory ended")
}

// @Security ApiKeyAuth
// @Summary Delete story
// @Description Delete new story
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} string "succesfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [delete]
func (h *Handler) DeleteStories(ctx *gin.Context) {
	h.Logger.Info("DeleteStory started")
	req := pb.DeleteStoriesRequest{}
	req.Id = ctx.Param("story_id")
	_, err := h.Content.DeleteStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "succesfully deleted"})
	h.Logger.Info("DeleteStory ended")
}


// @Security ApiKeyAuth
// @Summary get all story
// @Description gets stories
// @Tags stories
// @Param limit query string false "Number of stories to fetch"
// @Param page query string false "Number of stories to omit"
// @Success 200 {object} content_service.GetsStoriesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories [get]
func (h *Handler) GetsStories(ctx *gin.Context) {
	h.Logger.Info("Gets Stories started")
	req := pb.GetsStoriesRequest{}
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

	res, err := h.Content.GetsStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("Gets Stories ended")
}


// @Security ApiKeyAuth
// @Summary Get stories
// @Description Get stories by id
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} content_service.GetStoriesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [get]
func (h *Handler) GetStories(ctx *gin.Context){
	h.Logger.Info("GetStory started")
	req := pb.GetStoriesRequest{}
	req.Id = ctx.Param("story_id")
	if len(req.Id) <= 0 {

		h.Logger.Error("id is empty")
		ctx.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.Content.GetStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("GetStory ended")
}

// @Security ApiKeyAuth
// @Summary comment story
// @Description comment to story
// @Tags stories
// @Param story_id path string true "story_id"
// @Param info body content_service.CommentStoriesRequest true "story_id"
// @Success 200 {object} content_service.CommentStoriesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/comments [post]
func (h *Handler) CommentStories(ctx *gin.Context){
	h.Logger.Info("CommentStory started")
	accessToken := ctx.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.CommentStoriesRequest{}
	ctx.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	req.AuthorId = id
	req.StoryId = ctx.Param("story_id")
	if len(req.StoryId) <= 0 {

		h.Logger.Error("id is empty")
		ctx.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.Content.CommentStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("CommentStory ended")
}

// GetCommentsOfStory godoc
// @Security ApiKeyAuth
// @Summary comment of story
// @Description get comment of story
// @Tags stories
// @Param limit query string false "Number of stories to fetch"
// @Param page query string false "Number of stories to omit"
// @Success 200 {object} content_service.GetCommentStoriesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/comments [get]
func (h *Handler) GetCommentStories(ctx *gin.Context){
	h.Logger.Info("GetCommentsOfStory started")
	req := pb.GetCommentStoriesRequest{}
	// req.StoryId = c.Param("story_id")
	// if len(req.StoryId) <= 0 {

	// 	h.Log.Error("id is empty")
	// 	c.JSON(400, gin.H{"error": "id is empty"})

	// }
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

	res, err := h.Content.GetCommentStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("GetCommentsOfStory ended")
}

// @Security ApiKeyAuth
// @Summary Likes	 story
// @Description comment to story
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} content_service.LikeStoriesResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/like [post]
func (h *Handler) LikeStories(ctx *gin.Context){
	h.Logger.Info("Like started")
	// accessToken := ctx.GetHeader("Authorization")
	// id, err := auth.GetUserIdFromAccessToken(accessToken)
	// if err != nil {
	// 	h.Logger.Error(err.Error())
	// 	ctx.JSON(400, gin.H{"error": err.Error()})
	// }
	req := pb.LikeStoriesRequest{}
	req.StoriesId = ctx.Param("story_id")
	if len(req.StoriesId) == 0 {
		h.Logger.Error("id is empty")
		ctx.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.Content.LikeStories(ctx, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, &res)
	h.Logger.Info("Like ended")
}