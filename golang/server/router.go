package server

import (
  "context"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/google/go-github/v35/github"

  "github.com/svo/our-source/golang/model"
  "github.com/svo/our-source/golang/repository"
)

type RouterFactory interface {
  Build() *gin.Engine
}

type GinRouterFactory struct {}

func (routerFactory GinRouterFactory) Build() *gin.Engine {
  router := gin.Default()
  context := context.Background()
  sessionContext := model.SessionContext{}.New("coconuts")
  clientFactory := func(httpClient *http.Client) *github.Client { return github.NewClient(httpClient) }
  gitHubContext := (&repository.GitHubContext{}).New(clientFactory, context, sessionContext)

  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  router.GET("/organisation/:organisation/team/:team/repository", func(c *gin.Context) {
    teamContext := model.TeamContext{}.New(c.Param("organisation"), c.Param("team"))
    c.JSON(200, gin.H{
      "message": gitHubContext.Select(teamContext),
    })
  })

  return router
}
