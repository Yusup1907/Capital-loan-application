package handler

import (
	"log"
	"net/http"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	router  *gin.Engine
	usecase usecase.LoanApplicationUsecase
}

func (lh *LoanHandler) CreateLoanApplication(ctx *gin.Context) {
	var req model.LoanJoinRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	loan := model.LoanApplicationModel{
		CustomerId:     req.CustomerId,
		LoanDate:       req.LoanDate,
		DueDate:        req.DueDate,
		CategoryLoanId: req.CategoryLoanId,
		Amount:         req.Amount,
		Description:    req.Description,
		Status:         req.Status,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = lh.usecase.CreateLoanApplication(&loan)
	if err != nil {
		log.Println("Failed to create loan application:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Failed to create loan application",
		})
		return
	}

	if loan.Status == model.LoanStatusApprove {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  model.LoanStatusApprove,
			"message": "Loan application created successfully",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  model.LoanStatusDenied,
			"message": "Failed to create loan application",
		})
	}
}

func NewLoanApplicationHandler(r *gin.Engine, usecase usecase.LoanApplicationUsecase) *LoanHandler {
	handler := LoanHandler{
		router:  r,
		usecase: usecase,
	}
	r.POST("/loan", handler.CreateLoanApplication)
	// r.GET("/product", handler.getAllProduct)
	// r.GET("/product/:id", handler.getProductById)
	// r.PUT("/product/:id", handler.updateProduct)
	// r.DELETE("/product/:id", handler.deleteProduct)

	return &handler
}
