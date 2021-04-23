package repository

import (
  "context"

  "github.com/google/go-github/v35/github"
  "github.com/svo/our-source/golang/model"
  "golang.org/x/oauth2"
)

type GitHubRepositoryByTeamLister interface {
  ListTeamReposBySlug(ctx context.Context, org, slug string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error)
}

type GitHubClient struct {
  gitHubRepositoryByTeamLister GitHubRepositoryByTeamLister
}

func (repositoryClient *GitHubClient) Select(context context.Context, userContext model.UserContext, teamContext model.TeamContext)[]*github.Repository {
  repository, _, _ := repositoryClient.gitHubRepositoryByTeamLister.ListTeamReposBySlug(context, teamContext.Organisation(), teamContext.Team(), &github.ListOptions{})

  return repository
}

func NewClient(userContext model.UserContext, context context.Context) (*github.Client) {
  return github.NewClient(oauth2.NewClient(context, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: userContext.AccessToken()})))
}
