package server

import "github.com/gin-gonic/gin"

func Init(factory RouterFactory) *gin.Engine {
  return factory.Build()
}
