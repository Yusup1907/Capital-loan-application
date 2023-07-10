package handler

import (
	"fmt"
	"net/http"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryProductHandler struct {
	cpUsecase usecase.CategoryProductUsecase
}

func (cpHandler *CategoryProductHandler) InsertCategoryProduct(ctx *gin.Context){
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
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
		return
	}

	err = cpHandler.cpUsecase.InsertCategoryProduct(cp)
	if err != nil{
		fmt.Printf("error an cpHandler.cpUsecase.InsertCategoryProduct : %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data category product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (cpHandler *CategoryProductHandler) GetCategoryProductById(ctx *gin.Context){
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

func NewCategoryProductHandler(srv *gin.Engine, cpUsecase usecase.CategoryProductUsecase) CategoryProductHandler{
	handler := &CategoryProductHandler{
		cpUsecase: cpUsecase,
	}

	srv.POST("/category_product", handler.InsertCategoryProduct)
	srv.GET("/category_product/:id", handler.GetCategoryProductById)
	return handler
}