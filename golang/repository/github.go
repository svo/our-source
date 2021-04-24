package repository

import (
  "context"
  "net/http"

  "github.com/google/go-github/v35/github"
  "github.com/svo/our-source/golang/model"
  "golang.org/x/oauth2"
)

type GitHubRepositoryByTeamLister interface {
  ListTeamReposBySlug(ctx context.Context, org, slug string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
}

type GitHubContext struct {
  context context.Context
  sessionContext model.SessionContext
  gitHubRepositoryByTeamLister GitHubRepositoryByTeamLister
}

func (gitHubContext *GitHubContext) Select(teamContext model.TeamContext)[]*github.Repository {
  repository, _, _ := gitHubContext.gitHubRepositoryByTeamLister.ListTeamReposBySlug(gitHubContext.context, teamContext.Organisation(), teamContext.Team(), &github.ListOptions{})

  return repository
}

func (gitHubContext *GitHubContext) New(clientFactory func(httpClient *http.Client) *github.Client, context context.Context, sessionContext model.SessionContext) GitHubContext {
  client := clientFactory(oauth2.NewClient(context, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: sessionContext.AccessToken()})))
  return GitHubContext {context, sessionContext, client.Teams }
}
