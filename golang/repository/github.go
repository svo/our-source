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

type GitHubClientFactory interface {
  NewClient(httpClient *http.Client) *github.Client
}

type GitHubContext struct {
  clientFactory GitHubClientFactory
  gitHubRepositoryByTeamLister GitHubRepositoryByTeamLister
}

func (gitHubContext *GitHubContext) Select(context context.Context, userContext model.UserContext, teamContext model.TeamContext)[]*github.Repository {
  repository, _, _ := gitHubContext.gitHubRepositoryByTeamLister.ListTeamReposBySlug(context, teamContext.Organisation(), teamContext.Team(), &github.ListOptions{})

  return repository
}

func (gitHubContext *GitHubContext) NewClient(context context.Context, userContext model.UserContext) (*github.Client) {
  return gitHubContext.clientFactory.NewClient(oauth2.NewClient(context, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: userContext.AccessToken()})))
}
