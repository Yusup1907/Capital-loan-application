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
	srv *gin.Engine
}

func (s *server) Run() {
	NewCategoryProductHandler(s.srv, s.usecaseManager.GetCategoryProductUsecase())

	s.srv.Run(":8080")
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv: srv,
	}
}