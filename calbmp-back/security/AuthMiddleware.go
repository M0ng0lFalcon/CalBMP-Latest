package security

import (
	"calbmp-back/Database"
	"calbmp-back/Res"
	"calbmp-back/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			Res.FailMsg(ctx, "[!] Insufficient permissions")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			Res.FailMsg(ctx, "[!] Insufficient permissions")
			ctx.Abort()
			return
		}

		// 验证通过后获取 claim 中的userID
		userId := claims.UserId
		DB := Database.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			Res.FailMsg(ctx, "[!] Insufficient permissions")
			ctx.Abort()
			return
		}

		// 用户存在 user 信息写入上下文
		ctx.Set("user", user)
		ctx.Set("username", user.Username)
		ctx.Next()
	}
}
