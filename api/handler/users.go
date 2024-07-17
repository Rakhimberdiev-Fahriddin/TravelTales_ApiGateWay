package handler

import (
	"api/api/auth"
	pb "api/generated/auth_service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// @Summary get user
// @Description you can see your profile
// @Tags uscrers
// @Success 200 {object} auth_service.GetProfileResponce
// @Failure 401 {object} string "Invalid token"
// @Failure 500 {object} string "error while reading from server"
// @Router /api/v1/users/profile [get]
func (h *Handler) GetUserProfile(ctx *gin.Context) {
	h.Logger.Info("Profile is working")
	accessToken := ctx.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	req := pb.GetProfileRequest{
		Id: id,
	}
	res, err := h.User.GetUserProfile(ctx,&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error1": err.Error()})
	}
	ctx.JSON(http.StatusOK,res)
	h.Logger.Info("Profile ended")
}

// @Security ApiKeyAuth
// @Summary Update user
// @Description you can update your profile
// @Tags users
// @Param userinfo body auth_service.UpdateProfileRequest true "info"
// @Success 200 {object} auth_service.UpdateProfileResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "error while reading from server"
// @Router /api/v1/users/profile [put]
func (h *Handler) UpdateUserProfile(ctx *gin.Context) {
	h.Logger.Info("UserProfileUpdate is working")
	accessToken := ctx.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := pb.UpdateProfileRequest{UserId: id}
	if err := ctx.BindJSON(&req); err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.User.UpdateUserProfile(ctx,&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
	h.Logger.Info("UserProfileUpdate ended")
}

// @Security ApiKeyAuth
// @Summary List users
// @Description you can see all users
// @Tags users
// @Param limit query string false "Number of users to fetch"
// @Param page query string false "Number of users to omit"
// @Success 200 {object} auth_service.ListProfileResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "error while reading from server"
// @Router /api/v1/users [get]
func (h *Handler) ListUsersProfile(ctx *gin.Context) {
	h.Logger.Info("List Users is working")
	req := pb.ListProfileRequest{}
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

	res, err := h.User.ListUsersProfile(ctx,&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
	h.Logger.Info("GetAllUsers ended")
}

// @Security ApiKeyAuth
// @Summary delete user
// @Description you can delete your profile
// @Tags users
// @Param user_id path string true "user_id"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "error while reading from server"
// @Router /api/v1/users/{user_id} [delete]
func (h *Handler) DeleteUserProfile(ctx *gin.Context) {
	h.Logger.Info("Delete is working")
	id := ctx.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user id is incorrect"})
		return
	}
	req := pb.DeleteProfileRequest{Id: id}
	_, err = h.User.DeleteUserProfile(ctx,&req)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
	h.Logger.Info("Delete ended")
}

// @Security ApiKeyAuth
// @Summary Activities user
// @Description you can see your profile activity
// @Tags users
// @Param user_id path string true "user_id"
// @Success 200 {object} auth_service.ActivityResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "error while reading from server"
// @Router /api/v1/users/{user_id}/activity [get]
func (h *Handler) ActivityUser(ctx *gin.Context) {
	h.Logger.Info("Activity is working")
	id := ctx.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id is incorrect"})
		return
	}
	req := pb.ActivityRequest{
		UserId: id,
	}

	res, err := h.User.ActivityUser(ctx,&req)

	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, res)
	h.Logger.Info("Activity ended")
}

// @Security ApiKeyAuth
// @Summary follow user
// @Description you can follow another user
// @Tags users
// @Param user_id path string true "user_id"
// @Success 200 {object} auth_service.FollowResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "error while reading from server"
// @Router /api/v1/users/{user_id}/follow [post]
func (h *Handler) FollowUser(ctx *gin.Context) {
	h.Logger.Info("Follow is working")
	id := ctx.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id is incorrect"})
	}

	accessToken := ctx.GetHeader("Authorization")
	idFollower, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
	}
	req := pb.FollowRequest{
		FollowerId:  idFollower,
		FollowingId: id,
	}
	res, err := h.User.FollowUser(ctx,&req)
	fmt.Println(res,err)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, res)
	h.Logger.Info("Follow ended")
}

// @Security ApiKeyAuth
// @Summary get followers
// @Description you can see your followers
// @Tags users
// @Param user_id path string true "user_id"
// @Param limit query string false "Number of users to fetch"
// @Param page query string false "Number of users to omit"
// @Success 200 {object} auth_service.FollowersResponce
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "error while reading from server"
// @Router /api/v1/users/{user_id}/followers [get]
func (h *Handler) FollowersUsers(ctx *gin.Context) {
	h.Logger.Info("Followers is working")
	id := ctx.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id is incorrect"})
	}
	req := pb.FollowersRequest{UserId: id}

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

	res, err := h.User.FollowersUsers(ctx,&req)
	if err != nil {
		h.Logger.Error(err.Error())
	}
	ctx.JSON(http.StatusOK, res)
	h.Logger.Info("Followers ended")
}
