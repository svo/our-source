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

type MockGitHubClientFactory struct {
  mock.Mock
}

func (client *MockGitHubClientFactory) NewClient(httpClient *http.Client) *github.Client {

  args := client.Called(httpClient)
  return args.Get(0).(*github.Client)
}

type GitHubSuite struct {
  suite.Suite
}

func (suite *GitHubSuite) TestRetrievesGitHubRepositoryByTeamInOrganisation() {
  var expected []*github.Repository
  lister := &MockGitHubRepositoryByTeamLister{}
  gitHubContext := &GitHubContext{gitHubRepositoryByTeamLister: lister}
  context := context.Background()
  userContext := model.UserContext{}.New("coconuts")
  teamContext := model.TeamContext{}.New("bob", "mary")

  lister.On("ListTeamReposBySlug", context, "bob", "mary", mock.Anything).Return(expected, &github.Response{}, nil)

  result := gitHubContext.Select(context, userContext, teamContext)

  assert.Equal(suite.T(), expected, result)
  lister.AssertNumberOfCalls(suite.T(), "ListTeamReposBySlug", 1)
  lister.AssertExpectations(suite.T())
}

func (suite *GitHubSuite) TestGetsConnectionToGitHub() {
  var expected *github.Client
  clientFactory := &MockGitHubClientFactory{}
  gitHubContext := &GitHubContext{clientFactory: clientFactory}
  context := context.Background()
  userContext := model.UserContext{}.New("coconuts")

  clientFactory.On("NewClient", mock.Anything).Return(expected)

  result := gitHubContext.NewClient(context, userContext)

  assert.Equal(suite.T(), expected, result)
  clientFactory.AssertNumberOfCalls(suite.T(), "NewClient", 1)
  clientFactory.AssertExpectations(suite.T())
}

func TestGitHubSuite(t *testing.T) {
  suite.Run(t, new(GitHubSuite))
}
