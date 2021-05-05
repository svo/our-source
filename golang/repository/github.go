package repository

import (
	"context"
	"net/http"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"

	"github.com/svo/our-source/golang/model"
	"github.com/svo/our-source/golang/transformer"
)

type GitHubRepositoryByTeamLister interface {
	ListTeamReposBySlug(ctx context.Context, org, slug string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
}

type GitHub interface {
	Select(teamContext model.TeamContext) []model.Repository
}

type GoGitHubContext struct {
	context                      context.Context
	sessionContext               model.SessionContext
	gitHubRepositoryByTeamLister GitHubRepositoryByTeamLister
	repositoryTransformer        transformer.RepositoryToModelTransformer
}

func (gitHubContext *GoGitHubContext) Select(teamContext model.TeamContext) []model.Repository {
	repository, _, _ := gitHubContext.gitHubRepositoryByTeamLister.ListTeamReposBySlug(gitHubContext.context, teamContext.Organisation(), teamContext.Team(), &github.ListOptions{})

	result := make([]model.Repository, 0)
	for _, value := range repository {
		result = append(result, gitHubContext.repositoryTransformer.Transform(*value))
	}

	return result
}

func (gitHubContext *GoGitHubContext) New(clientFactory func(httpClient *http.Client) *github.Client, repositoryTransformer transformer.RepositoryToModelTransformer, context context.Context, sessionContext model.SessionContext) GoGitHubContext {
	client := clientFactory(oauth2.NewClient(context, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: sessionContext.AccessToken()})))
	return GoGitHubContext{context, sessionContext, client.Teams, repositoryTransformer}
}
