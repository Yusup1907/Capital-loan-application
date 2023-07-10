package handler

import (
	"errors"
	"fmt"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

type CategoryLoanHandler struct {
	router  *gin.Engine
	usecase usecase.CategoryLoanUsecase
}

func (ch *CategoryLoanHandler) InsertCategoryLoan(ctx *gin.Context) {
	var req model.InsertCategoryLoan
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryLoan := model.CategoryLoanModel{
		CategoryLoanName: req.CategoryLoanName,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := ch.usecase.InsertCategoryLoan(&categoryLoan)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CategoryLoanHandler.InsertCategoryLoan() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CategoryLoanHandler.InsertCategoryLoan() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot insert category loan due to an error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewCategoryLoanHandler(r *gin.Engine, usecase usecase.CategoryLoanUsecase) *CategoryLoanHandler {
	handler := CategoryLoanHandler{
		router:  r,
		usecase: usecase,
	}

	r.POST("/category-loan", handler.InsertCategoryLoan)
	return &handler
}
