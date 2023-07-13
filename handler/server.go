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
<<<<<<< HEAD
	engine         *gin.Engine
}

func (s *server) Run() {
	NewProductHandler(s.engine, s.usecaseManager.GetProductUsecase())
	NewCategoryProductHandler(s.srv, s.usecaseManager.GetCategoryProductUsecase())
	NewGoodsHandler(s.srv, s.usecaseManager.GetGoodsUsecase())

	s.engine.Run()
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

<<<<<<< HEAD
	engine := gin.Default()

	return &server{
		usecaseManager: usecase,
		engine:         engine,
	}

}
