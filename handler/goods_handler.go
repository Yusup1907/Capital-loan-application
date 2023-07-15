package handler

import (
	// "fmt"
	"fmt"
	"log"
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

func(goodsHandler *goodsHandlerImpl) GetAllTrxGoods(ctx *gin.Context){
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = 1
		}
	
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			limit = 10
		}
	
		loanGoods, err := goodsHandler.goodsUsecase.GetAllTrxGoods(page, limit)
		if err != nil {
			log.Println("Failed to create loan application:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Failed to retrieve loan applications",
			})
			return
		}
	
		response := make([]model.LoanGoodsModel, 0)
		for _, loanGood := range loanGoods {
			response = append(response, *loanGood)
		}
	
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
	
func NewGoodsHandler(srv *gin.Engine,goodsUsecase usecase.GoodsUsecase) GoodsHandler {
	ghandler := goodsHandlerImpl{
		router: srv,
		goodsUsecase: goodsUsecase,
	}

	srv.POST("/goods", ghandler.InsertGoods)
	srv.GET("/goods/:id", ghandler.GetGoodsById)
	srv.GET("/goods", ghandler.GetAllTrxGoods)
	return ghandler
}
