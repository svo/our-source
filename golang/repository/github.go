package repository

import (
	"context"
	"net/http"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"

	"github.com/svo/our-source/golang/model"
)

type GitHubRepositoryByTeamLister interface {
	ListTeamReposBySlug(ctx context.Context, org, slug string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
}

type GitHub interface {
	Select(teamContext model.TeamContext) []*github.Repository
}

type GoGitHubContext struct {
	context                      context.Context
	sessionContext               model.SessionContext
	gitHubRepositoryByTeamLister GitHubRepositoryByTeamLister
}

func (gitHubContext *GoGitHubContext) Select(teamContext model.TeamContext) []*github.Repository {
	repository, _, _ := gitHubContext.gitHubRepositoryByTeamLister.ListTeamReposBySlug(gitHubContext.context, teamContext.Organisation(), teamContext.Team(), &github.ListOptions{})

	return repository
}

func (gitHubContext *GoGitHubContext) New(clientFactory func(httpClient *http.Client) *github.Client, context context.Context, sessionContext model.SessionContext) GoGitHubContext {
	client := clientFactory(oauth2.NewClient(context, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: sessionContext.AccessToken()})))
	return GoGitHubContext{context, sessionContext, client.Teams}
}
