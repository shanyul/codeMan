package middleware

import (
	"chujian-api/internal/controller"
	"chujian-api/internal/service"
	"chujian-api/pkg/app"
	"chujian-api/pkg/exception"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authority() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			data  interface{}
			token = c.GetHeader("Token")
		)

		if token == "" {
			app.Response(c, http.StatusUnauthorized, exception.InvalidParams, data, "")
			c.Abort()
			return
		}

		var authService service.AuthService
		user := authService.IsLogin(token)
		if user.UserId == 0 {
			app.Response(c, http.StatusUnauthorized, exception.ErrorTokenExpire, data, "")
			c.Abort()
			return
		}

		userInfo := controller.UserInfo
		userInfo.UserId = user.UserId
		userInfo.Username = user.Username
		userInfo.Nickname = user.Nickname
		userInfo.Phone = user.Phone
		userInfo.SessionKey = user.SessionKey

		c.Next()
	}
}
