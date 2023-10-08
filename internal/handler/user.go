package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nunu-template/internal/pkg/request/user"
	"nunu-template/internal/pkg/response"
	"nunu-template/internal/service"
)

type UserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type userHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) UserHandler {
	return &userHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *userHandler) Register(ctx *gin.Context) {
	req := new(user.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.HandleError(ctx, http.StatusOK, response.ErrBadRequest, nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		response.HandleError(ctx, http.StatusOK, err, nil)
		return
	}

	response.HandleSuccess(ctx, nil)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var req user.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, http.StatusOK, response.ErrBadRequest, nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		response.HandleError(ctx, http.StatusOK, response.ErrUnauthorized, nil)
		return
	}

	response.HandleSuccess(ctx, gin.H{
		"accessToken": token,
	})
}

func (h *userHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		response.HandleError(ctx, http.StatusOK, response.ErrUnauthorized, nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		response.HandleError(ctx, http.StatusOK, err, nil)
		return
	}

	response.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req user.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleError(ctx, http.StatusOK, response.ErrBadRequest, nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		response.HandleError(ctx, http.StatusOK, response.ErrUpdateUser, nil)
		return
	}

	response.HandleSuccess(ctx, nil)
}
