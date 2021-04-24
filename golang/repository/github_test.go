package repository

import (
  "context"
  "testing"
  "net/http"

  "github.com/google/go-github/v35/github"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/suite"
  "github.com/svo/our-source/golang/model"
)

type MockGitHubRepositoryByTeamLister struct {
  mock.Mock
}

func (lister *MockGitHubRepositoryByTeamLister) ListTeamReposBySlug(ctx context.Context, org, slug string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error) {

  args := lister.Called(ctx, org, slug, opts)
  return args.Get(0).([]*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}

type GitHubSuite struct {
  suite.Suite
}

func (suite *GitHubSuite) TestRetrievesGitHubRepositoryByTeamInOrganisation() {
  var expected []*github.Repository
  context := context.Background()
  lister := &MockGitHubRepositoryByTeamLister{}
  gitHubContext := &GitHubContext{context: context, sessionContext: model.SessionContext{}, gitHubRepositoryByTeamLister: lister}
  teamContext := model.TeamContext{}.New("bob", "mary")

  lister.On("ListTeamReposBySlug", context, "bob", "mary", mock.Anything).Return(expected, &github.Response{}, nil)

  result := gitHubContext.Select(teamContext)

  assert.Equal(suite.T(), expected, result)
  lister.AssertNumberOfCalls(suite.T(), "ListTeamReposBySlug", 1)
  lister.AssertExpectations(suite.T())
}

func (suite *GitHubSuite) TestCreatesGitHubContextWithGitHubRepositoryByTeamLister() {
  teamsService := &github.TeamsService{}
  clientFactory := func(httpClient *http.Client) *github.Client { return &github.Client{Teams: teamsService} }

  assert.Equal(suite.T(), teamsService, (&GitHubContext{}).New(clientFactory, context.Background(), model.SessionContext{}).gitHubRepositoryByTeamLister)
}

func (suite *GitHubSuite) TestCreatesGitHubContextWithContext() {
  context := context.Background()
  clientFactory := func(httpClient *http.Client) *github.Client { return &github.Client{} }

  assert.Equal(suite.T(), context, (&GitHubContext{}).New(clientFactory, context, model.SessionContext{}).context)
}

func (suite *GitHubSuite) TestCreatesGitHubContextWithSessionContext() {
  clientFactory := func(httpClient *http.Client) *github.Client { return &github.Client{} }
  sessionContext := model.SessionContext{}

  assert.Equal(suite.T(), sessionContext, (&GitHubContext{}).New(clientFactory, context.Background(), sessionContext).sessionContext)
}

func TestGitHubSuite(t *testing.T) {
  suite.Run(t, new(GitHubSuite))
}
