package routes

import (
	"net/http"
	"resto-backend-go-1/helpers"
	web_users "resto-backend-go-1/models/web/users"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func Routes() {
	r := gin.Default()

	userRouter := r.Group("/users")

	{
		userRouter.GET("/", func(ctx *gin.Context) {

		})

		userRouter.POST("/", func(ctx *gin.Context) {
			var registerInput web_users.RegisterUserRequest
			contentType := helpers.GetContentType(ctx)
			if contentType == appJSON {
				ctx.ShouldBindJSON(&registerInput)
			} else {
				ctx.ShouldBind(&registerInput)
			}

			web_users.RegisterResponse(ctx, http.StatusCreated, registerInput, "Success Register ")
		})
	}

	r.Run()
}
