package server

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v35/github"

	"github.com/svo/our-source/golang/model"
	"github.com/svo/our-source/golang/repository"
)

func Init(factory RouterFactory) *gin.Engine {
	context := context.Background()
	sessionContext := model.SessionContext{}.New("coconuts")
	clientFactory := func(httpClient *http.Client) *github.Client { return github.NewClient(httpClient) }
	gitHubContext := (&repository.GoGitHubContext{}).New(clientFactory, context, sessionContext)
	return factory.Build(&gitHubContext)
}
