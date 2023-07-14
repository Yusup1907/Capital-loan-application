package handler

import (
	"pinjam-modal-app/manager"
	"pinjam-modal-app/middleware"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	srv            *gin.Engine
}

func (s *server) Run() {
	s.srv.Use(middleware.LoggerMiddleware())
	NewCategoryLoanHandler(s.srv, s.usecaseManager.GetCategoryLoanUsecase())
	NewUserHandler(s.srv, s.usecaseManager.GetUserUsecase())
	NewLoginHandler(s.srv, s.usecaseManager.GetLoginUsecase())

	s.srv.Run()
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	engine := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            engine,
	}
}
