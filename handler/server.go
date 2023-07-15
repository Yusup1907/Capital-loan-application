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
	NewCustomerHandler(s.engine, s.usecaseManager.GetCustomerUsecase())
	NewProductHandler(s.engine, s.usecaseManager.GetProductUsecase())
	NewCategoryProductHandler(s.engine, s.usecaseManager.GetCategoryProductUsecase())
	NewGoodsHandler(s.engine, s.usecaseManager.GetGoodsUsecase())
	NewCategoryLoanHandler(s.engine, s.usecaseManager.GetCategoryLoanUsecase())
	NewLoanApplicationHandler(s.engine, s.usecaseManager.GetLoanAppUsecase())

	s.engine.Run(":8080")
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
