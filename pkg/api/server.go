package api

import (
	"github.com/athunlal/Note-Taking-Application/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())

	return &ServerHttp{
		Engine: engine,
	}
}

func (ser *ServerHttp) Start() {
	ser.Engine.Run(":8080")
}
