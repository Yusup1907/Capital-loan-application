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
	NewUserHandler(s.engine, s.usecaseManager.GetUserUsecase())
	NewLoginHandler(s.engine, s.usecaseManager.GetLoginUsecase())
	NewCustomerHandler(s.engine, s.usecaseManager.GetCustomerUsecase())
	NewProductHandler(s.engine, s.usecaseManager.GetProductUsecase())
	NewCategoryProductHandler(s.engine, s.usecaseManager.GetCategoryProductUsecase())
	NewGoodsHandler(s.engine, s.usecaseManager.GetGoodsUsecase())
	NewCategoryLoanHandler(s.engine, s.usecaseManager.GetCategoryLoanUsecase())
	NewLoanApplicationHandler(s.engine, s.usecaseManager.GetLoanAppUsecase())

	s.engine.Use(middleware.LoggerMiddleware())

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
