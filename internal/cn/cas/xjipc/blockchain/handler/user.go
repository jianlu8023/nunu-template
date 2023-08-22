package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"nunu-template/internal/cn/cas/xjipc/blockchain/pkg/request"
	"nunu-template/internal/cn/cas/xjipc/blockchain/service"
	"nunu-template/pkg/helper/resp"
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
	req := new(request.RegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.Register(ctx, req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	token, err := h.userService.Login(ctx, &req)
	if err != nil {
		resp.HandleError(ctx, http.StatusUnauthorized, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, gin.H{
		"accessToken": token,
	})
}

func (h *userHandler) GetProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)
	if userId == "" {
		resp.HandleError(ctx, http.StatusUnauthorized, 1, "unauthorized", nil)
		return
	}

	user, err := h.userService.GetProfile(ctx, userId)
	if err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, user)
}

func (h *userHandler) UpdateProfile(ctx *gin.Context) {
	userId := GetUserIdFromCtx(ctx)

	var req request.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, errors.Wrap(err, "invalid request").Error(), nil)
		return
	}

	if err := h.userService.UpdateProfile(ctx, userId, &req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}
