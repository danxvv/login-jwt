package app

import (
	"github.com/gin-gonic/gin"
	"login-user/config"
	v1 "login-user/internal/controller/http/v1"
	"login-user/internal/entity"
	"login-user/internal/usecase"
	"login-user/internal/usecase/repository"
	"login-user/pkg/httpserver"
	"login-user/pkg/postgres"
)

func RunApp(cfg *config.Config) {
	p, _ := postgres.NewPostgresConnection(cfg)
	err := p.DB.AutoMigrate(&entity.User{})
	if err != nil {
		return
	}

	defer p.Close()
	userRepo := repository.NewUserRepo(p)
	userUseCase := usecase.NewUserUsecase(userRepo)

	ginHandler := gin.New()

	validator := v1.GetValidator()

	v1.NewRouter(ginHandler, userUseCase, validator)
	httpServer := httpserver.New(ginHandler, cfg.HTTP.Port)
	httpServer.Start()
}
