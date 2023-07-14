package handler

import (
	"log"
	"net/http"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	router  *gin.Engine
	usecase usecase.LoanApplicationUsecase
}

func (lh *LoanHandler) createLoanApplication(ctx *gin.Context) {
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
		CustomerId:      req.CustomerId,
		LoanDate:        req.LoanDate,
		DueDate:         req.DueDate,
		CategoryLoanId:  req.CategoryLoanId,
		Amount:          req.Amount,
		Description:     req.Description,
		Status:          req.Status,
		RepaymentStatus: req.RepaymentStatus,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
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

func (lh *LoanHandler) getLoanApplications(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	loanApplications, err := lh.usecase.GetLoanApplications(page, limit)
	if err != nil {
		log.Println("Failed to create loan application:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Failed to retrieve loan applications",
		})
		return
	}

	response := make([]model.LoanApplicationJoinModel, 0)
	for _, loanApplication := range loanApplications {
		response = append(response, *loanApplication)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

func (lh *LoanHandler) getLoanApplicationById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid ID",
		})
		return
	}

	loanApplication, err := lh.usecase.GetLoanApplicationById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Failed to retrieve loan application",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    loanApplication,
	})
}

func NewLoanApplicationHandler(r *gin.Engine, usecase usecase.LoanApplicationUsecase) *LoanHandler {
	handler := LoanHandler{
		router:  r,
		usecase: usecase,
	}
	r.POST("/loan", handler.createLoanApplication)
	r.GET("/loan", handler.getLoanApplications)
	r.GET("/loan/:id", handler.getLoanApplicationById)
	// r.PUT("/product/:id", handler.updateProduct)
	// r.DELETE("/product/:id", handler.deleteProduct)

	return &handler
}
