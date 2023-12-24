package api

import (
	"github.com/athunlal/Note-Taking-Application/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, userHandler *handler.UserHandler) {

	user := r.Group("/user")
	{
		user.POST("/signup", userHandler.Register)
		user.POST("/login", userHandler.Login)

		user.POST("/note", userHandler.CreateNote)
		user.GET("/note", userHandler.GetNote)
		user.DELETE("/note", userHandler.DeleteNote)

	}
}
