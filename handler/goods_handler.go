package handler

import (
	// "fmt"
	"fmt"
	"net/http"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsHandler interface {
}

type goodsHandlerImpl struct {
	router  *gin.Engine
	goodsUsecase usecase.GoodsUsecase
}

func (goodsHandler *goodsHandlerImpl) InsertGoods(ctx *gin.Context) {
	goods := &model.GoodsModel{}
	err := ctx.ShouldBindJSON(&goods)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	err = goodsHandler.goodsUsecase.InsertGoods(goods)
	if err != nil{
		fmt.Printf("error an cpHandler.cpUsecase.InsertCategoryProduct : %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data category product",
		})
		return
	}

	if goods.Status == "APPROVE" {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "APPROVE",
			"message": "TrxGoods inserted successfully",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "DENIED",
			"message": "TrxGoods failed to insert",
		})
	}

}

func (goodsHandler *goodsHandlerImpl) GetGoodsById(ctx *gin.Context) {
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

	goods, err := goodsHandler.goodsUsecase.GetGoodsById(id)
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
		"data":    goods,
	})
}

func NewGoodsHandler(router *gin.Engine,goodsUsecase usecase.GoodsUsecase) GoodsHandler {
	ghandler := goodsHandlerImpl{
		goodsUsecase: goodsUsecase,
	}

	router.POST("/goods", ghandler.InsertGoods)
	router.GET("/goods/:id", ghandler.GetGoodsById)
	return ghandler
}
