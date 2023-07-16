package middleware

import (
	"errors"
	"net/http"
	"strings"

	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
	"pinjam-modal-app/utils/authutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(userRepo repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Periksa token dari header Authorization
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Misalnya, format header Authorization: Bearer <token>
		splitToken := strings.Split(authorizationHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := splitToken[1]

		// Validasi token
		token, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Setel data pengguna pada konteks Gin
		user, err := getUserFromToken(token, userRepo)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi metode tanda tangan token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid token")
		}

		// Kembalikan kunci rahasia yang sama yang digunakan saat pembuatan token
		return []byte(authutil.TokenKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}

func getUserFromToken(token *jwt.Token, userRepo repository.UserRepo) (*model.UserModel, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("Invalid username in token claims")
	}

	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}
