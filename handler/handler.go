package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/service"
)

type Handler struct {
	AuthHandler AuthController
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		AuthHandler: *NewAuthController(service.Auth, logger),
	}
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}
