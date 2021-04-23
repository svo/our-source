package repository

import (
  "context"
  "testing"

  "github.com/google/go-github/v35/github"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/suite"
  "github.com/svo/our-source/golang/model"
)

type MockGitHubRepositoryByTeamLister struct {
  mock.Mock
}

func (client *MockGitHubRepositoryByTeamLister) ListTeamReposBySlug(ctx context.Context, org, slug string, opts *github.ListOptions) ([]*github.Repository, *github.Response, error) {

  args := client.Called(ctx, org, slug, opts)
  return args.Get(0).([]*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}

type GitHubClientSuite struct {
  suite.Suite
}

func (suite *GitHubClientSuite) TestRetrievesGitHubRepositoryByTeamInOrganisation() {
  var expected []*github.Repository
  lister := &MockGitHubRepositoryByTeamLister{}
  theClient := &GitHubClient{gitHubRepositoryByTeamLister: lister}
  context := context.Background()
  userContext := model.UserContext{}.New("coconuts")
  teamContext := model.TeamContext{}.New("bob", "mary")

  lister.On("ListTeamReposBySlug", context, "bob", "mary", mock.Anything).Return(expected, &github.Response{}, nil)

  result := theClient.Select(context, userContext, teamContext)

  assert.Equal(suite.T(), expected, result)
  lister.AssertNumberOfCalls(suite.T(), "ListTeamReposBySlug", 1)
  lister.AssertExpectations(suite.T())
}

func TestGitHubClientSuite(t *testing.T) {
  suite.Run(t, new(GitHubClientSuite))
}
