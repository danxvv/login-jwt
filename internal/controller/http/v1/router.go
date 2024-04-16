package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"login-user/internal/usecase"
)

func NewRouter(handler *gin.Engine, u usecase.User, v *validator.Validate) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())

	h := handler.Group("/api/v1")
	{
		newUserRoutes(h, u, v)
	}
}
