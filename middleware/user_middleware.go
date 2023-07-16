package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		Email := session.Get("Email")
		userRole := session.Get("UserRole")
		if userRole == nil && Email == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Access Denied",
			})
			ctx.Abort()
			return
		}
		if userRole == "User" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Access Denied",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
