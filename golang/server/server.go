package server

import (
	"context"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v35/github"

	"github.com/svo/our-source/golang/configuration"
	"github.com/svo/our-source/golang/model"
	"github.com/svo/our-source/golang/repository"
	"github.com/svo/our-source/golang/transformer"
)

func Init(factory RouterFactory) *gin.Engine {
	context := context.Background()
	sessionContext := model.SessionContext{}.New(configuration.EnvironmentConfiguration{}.GetAccessToken())
	clientFactory := func(httpClient *http.Client) *github.Client { return github.NewClient(httpClient) }
	gitHubContext := (&repository.GoGitHubContext{}).New(clientFactory, new(transformer.RepositoryTransformer), context, sessionContext)
	return factory.Build(&gitHubContext)
}
