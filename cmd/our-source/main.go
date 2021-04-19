package main

import "github.com/gin-gonic/gin"

func main() {
  _ = setupRouter().Run()
}

func setupRouter() *gin.Engine {
  router := gin.Default()

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  return router
}