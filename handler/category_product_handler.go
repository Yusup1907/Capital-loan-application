package handler

import (
	"errors"
	"fmt"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryProductHandler interface{

}

type categoryProductHandler struct {
	cpUsecase usecase.CategoryProductUsecase
}

func (cpHandler *categoryProductHandler) InsertCategoryProduct(ctx *gin.Context){
	cp := &model.CategoryProductModel{}
	err := ctx.ShouldBindJSON(&cp)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid Data JSON",
		})
		return
	}

	if cp.CategoryProductName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Category Product Name Tidak Boleh Kosong",
		})
		return
	}

	err = cpHandler.cpUsecase.InsertCategoryProduct(cp)
	if err != nil{
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ServiceHandler.InsertService() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		}else {
		fmt.Printf("error an cpHandler.cpUsecase.InsertCategoryProduct : %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data category product",
		})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (cpHandler *categoryProductHandler) GetCategoryProductById(ctx *gin.Context){
	idText := ctx.Param("id")
	if idText == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	cp, err := cpHandler.cpUsecase.GetCategoryProductById(id)
	if err != nil {
		fmt.Printf(" cpHandler.cpUsecase.GetCategoryProductById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data category product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cp,
	})
}

func NewCategoryProductHandler(srv *gin.Engine, cpUsecase usecase.CategoryProductUsecase) CategoryProductHandler {
	handler := &categoryProductHandler{
		cpUsecase: cpUsecase,
	}

	srv.POST("/category_product", handler.InsertCategoryProduct)
	srv.GET("/category_product/:id", handler.GetCategoryProductById)
	return handler
}
