package server

import "github.com/gin-gonic/gin"

type RouterFactory interface {
  Build() *gin.Engine
}

type GinRouterFactory struct {}

func (routerFactory GinRouterFactory) Build() *gin.Engine {
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  return router
}
