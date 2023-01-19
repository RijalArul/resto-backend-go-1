package web_users

import "github.com/gin-gonic/gin"

type UserRegsisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Birthday string `json:"birthday"`
}

func RegisterResponse(ctx *gin.Context, statusCode int, data interface{}, message string) {
	ctx.JSON(statusCode, gin.H{
		"data":    data,
		"message": message,
	})
}
