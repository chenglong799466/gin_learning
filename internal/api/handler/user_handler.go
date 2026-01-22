package handler

import (
	"gin_learning/internal/api/response"
	"gin_learning/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.InvalidRequest(c, err.Error())
		return
	}

	userReq := service.CreateUserReq{
		Name: req.Name,
	}

	_, err := h.userService.CreateUser(&userReq)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.Success(c)
}
