package handler

import (
	"errors"
	"fmt"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	router  *gin.Engine
	usecase usecase.ProductUsecase
}

func (ph *ProductHandler) createProduct(ctx *gin.Context) {
	var req model.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := model.ProductModel{
		ProductName:       req.ProductName,
		Description:       req.Description,
		Price:             req.Price,
		Stok:              req.Stok,
		CategoryProductId: req.CategoryProductId,
		Status:            req.Status,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err := ph.usecase.CreateProduct(&product)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ProductHandler.createProduct() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ProductHandler.CreateProduct() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot Insert product because error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ph *ProductHandler) getAllProduct(ctx *gin.Context) {
	product, err := ph.usecase.GetAllProduct()
	if err != nil {
		fmt.Printf("ProductHandler.getAllProduct() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving product data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func (ph *ProductHandler) getProductById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}
	product, err := ph.usecase.GetProductById(id)
	if err != nil {
		fmt.Printf("ProductHandler.getProductById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving product data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func NewProductHandler(r *gin.Engine, usecase usecase.ProductUsecase) *ProductHandler {
	handler := ProductHandler{
		router:  r,
		usecase: usecase,
	}
	r.POST("/product", handler.createProduct)
	r.GET("/product", handler.getAllProduct)
	r.GET("/product/:id", handler.getProductById)
	return &handler
}
