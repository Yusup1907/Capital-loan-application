package handler

import (
	"fmt"
	"log"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usrUsecase usecase.UserUsecase
}

func (h *UserHandler) registerUser(c *gin.Context) {
	// Parsing request body
	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Validasi dan registrasi pengguna
	if err := h.usrUsecase.RegisterUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan response dengan data pengguna yang terdaftar
	c.JSON(http.StatusOK, user)
}

// func (usrHandler *UserHandler) UpadteUser(ctx *gin.Context) {
// 	usr := &model.UserModel{}
// 	err := ctx.ShouldBindJSON(&usr)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"success":      false,
// 			"errorMessage": "Invalid JSON data",
// 		})
// 		return
// 	}

// 	if len(usr.UserName) > 15 {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"success":      false,
// 			"errorMessage": "Panjang Nama tidak boleh lebih dari 15 karakter",
// 		})
// 		return
// 	}
// 	// Mengisi nilai createdAt dan updatedAt
// 	now := time.Now()
// 	usr.CreatedAt = now
// 	usr.UpdatedAt = now

// 	err = usrHandler.usrUsecase.UpadateUser(usr)
// 	if err != nil {
// 		if appErr, ok := err.(*apperror.AppError); ok {
// 			fmt.Printf("UserHandler.InsertUser() 1: %v\n", appErr)
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"success":      false,
// 				"errorMessage": appErr.Error(), // Menggunakan appErr.Error() untuk mendapatkan pesan error
// 			})
// 		} else {
// 			fmt.Printf("UserHandler.InsertUser() 2: %v\n", err)
// 			ctx.JSON(http.StatusInternalServerError, gin.H{
// 				"success":      false,
// 				"errorMessage": "Terjadi kesalahan ketika menyimpan data User",
// 			})
// 		}
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 	})
// }

func (usrHandler *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID tidak boleh kosong",
		})
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID harus berupa angka",
		})
		return
	}

	usr, err := usrHandler.usrUsecase.GetUserById(userId)
	if err != nil {
		fmt.Printf("UserHandler.GetUserById(): %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}
func (usrhandler *UserHandler) GetAllUser(ctx *gin.Context) {
	users, err := usrhandler.usrUsecase.GetAllUser()
	if err != nil {
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Failed to retrieve user data")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := gin.H{
		"success": true,
		"data":    users,
	}
	ctx.JSON(http.StatusOK, successResponse)
}
func (usrHandler *UserHandler) DeleteUser(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		err := apperror.NewAppError(http.StatusBadRequest, "ID cannot be empty")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		err := apperror.NewAppError(http.StatusBadRequest, "ID must be a number")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := usrHandler.usrUsecase.GetUserById(id)
	if err != nil {
		log.Printf("UserHandler.DeleteUser(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to delete user")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		err := apperror.NewAppError(http.StatusNotFound, "User not found")
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	err = usrHandler.usrUsecase.DeleteUser(user)
	if err != nil {
		log.Printf("UserHandler.User(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to delete User")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	successResponse := gin.H{
		"success": true,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func NewUserHandler(srv *gin.Engine, usrUsecase usecase.UserUsecase) *UserHandler {
	usrHandler := &UserHandler{
		usrUsecase: usrUsecase,
	}
	srv.POST("/user", usrHandler.registerUser)
	// srv.GET("/user", usrHandler.GetAllUser)
	// srv.GET("/user/:name", usrHandler.GetUserByName)
	// srv.GET("/user/id/:id", usrHandler.GetUserById)
	// srv.PUT("/user", usrHandler.UpadteUser)
	// srv.DELETE("/user/:id", usrHandler.DeleteUser)

	return usrHandler
}
