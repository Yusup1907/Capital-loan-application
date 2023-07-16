package handler

import (
	"net/http"

	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"

	utils "pinjam-modal-app/utils/authutil"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginUsecase usecase.LoginUsecase
}

// Handler untuk route /login
func (l *LoginHandler) Login(c *gin.Context) {
	var loginData model.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	// Simulasi pengecekan kredensial
	// Ganti dengan logika validasi sesuai kebutuhan Anda
	if loginData.Email == "" || loginData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Memanggil usecase Login
	token, err := utils.GenerateToken(loginData.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (l *LoginHandler) Logout(c *gin.Context) {
	err := l.LoginUsecase.Logout(c) // Sertakan konteks c saat memanggil Logout
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Logout successful",
	})
}

func NewLoginHandler(router *gin.Engine, loginUsecase usecase.LoginUsecase) {
	// Inisialisasi cookie store
	store := cookie.NewStore([]byte("secret-key")) // Ganti dengan kunci rahasia yang lebih kuat

	// Set pengaturan cookie store
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   true,
	})

	// Inisialisasi handler dengan penggunaan sesi
	loginHandler := &LoginHandler{
		LoginUsecase: loginUsecase,
	}

	router.Use(sessions.Sessions("session-name", store))

	router.POST("/login", loginHandler.Login)
	router.POST("/logout", loginHandler.Logout)
}
