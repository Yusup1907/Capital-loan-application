package handler

import (
	"pinjam-modal-app/manager"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
}

func (s *server) Run() {
	NewProductHandler(s.engine, s.usecaseManager.GetProductUsecase())

	s.engine.Run()
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	engine := gin.Default()

	return &server{
		usecaseManager: usecase,
		engine:         engine,
	}

}