package server

import (
	"github.com/gin-gonic/gin"

	"github.com/svo/our-source/golang/model"
	"github.com/svo/our-source/golang/repository"
)

type RouterFactory interface {
	Build(gitHubContext repository.GitHub) *gin.Engine
}

type GinRouterFactory struct{}

func (routerFactory GinRouterFactory) Build(gitHubContext repository.GitHub) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/organisation/:organisation/team/:team/repository", func(c *gin.Context) {
		teamContext := model.TeamContext{}.New(c.Param("organisation"), c.Param("team"))
		c.JSON(200, gitHubContext.Select(teamContext))
	})

	return router
}
