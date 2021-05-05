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

type MockRepositoryTransformer struct {
	mock.Mock
}

func (transformer *MockRepositoryTransformer) Transform(from github.Repository) model.Repository {
	args := transformer.Called(from)
	return args.Get(0).(model.Repository)
}

type GitHubSuite struct {
	suite.Suite
}

func (suite *GitHubSuite) TestRetrievesGitHubRepositoryByTeamInOrganisation() {
	var expected model.Repository
	fromRepository := &github.Repository{}
	repositoryResponse := []*github.Repository{fromRepository}
	context := context.Background()
	lister := &MockGitHubRepositoryByTeamLister{}
	repositoryTransformer := new(MockRepositoryTransformer)
	gitHubContext := &GoGitHubContext{context: context, sessionContext: model.SessionContext{}, gitHubRepositoryByTeamLister: lister, repositoryTransformer: repositoryTransformer}
	teamContext := model.TeamContext{}.New("bob", "mary")

	lister.On("ListTeamReposBySlug", context, "bob", "mary", mock.Anything).Return(repositoryResponse, &github.Response{}, nil)
	repositoryTransformer.On("Transform", *fromRepository).Return(expected)

	result := gitHubContext.Select(teamContext)

	assert.ElementsMatch(suite.T(), []model.Repository{expected}, result)
	lister.AssertNumberOfCalls(suite.T(), "ListTeamReposBySlug", 1)
	lister.AssertExpectations(suite.T())
	repositoryTransformer.AssertNumberOfCalls(suite.T(), "Transform", 1)
	repositoryTransformer.AssertExpectations(suite.T())
}

func (suite *GitHubSuite) TestReturnsEmptyArrayIfRetrievesNoGitHubRepositoryByTeamInOrganisation() {
	repositoryResponse := []*github.Repository{}
	context := context.Background()
	lister := &MockGitHubRepositoryByTeamLister{}
	repositoryTransformer := new(MockRepositoryTransformer)
	gitHubContext := &GoGitHubContext{context: context, sessionContext: model.SessionContext{}, gitHubRepositoryByTeamLister: lister, repositoryTransformer: repositoryTransformer}
	teamContext := model.TeamContext{}.New("bob", "mary")

	lister.On("ListTeamReposBySlug", context, "bob", "mary", mock.Anything).Return(repositoryResponse, &github.Response{}, nil)

	result := gitHubContext.Select(teamContext)
	assert.NotNil(suite.T(), result)
	assert.Len(suite.T(), result, 0)
	lister.AssertNumberOfCalls(suite.T(), "ListTeamReposBySlug", 1)
	lister.AssertExpectations(suite.T())
}

func (suite *GitHubSuite) TestCreatesGoGitHubContextWithGitHubRepositoryByTeamLister() {
	teamsService := &github.TeamsService{}
	clientFactory := func(httpClient *http.Client) *github.Client { return &github.Client{Teams: teamsService} }

	assert.Equal(suite.T(), teamsService, (&GoGitHubContext{}).New(clientFactory, new(MockRepositoryTransformer), context.Background(), model.SessionContext{}).gitHubRepositoryByTeamLister)
}

func (suite *GitHubSuite) TestCreatesGoGitHubContextWithContext() {
	context := context.Background()
	clientFactory := func(httpClient *http.Client) *github.Client { return &github.Client{} }

	assert.Equal(suite.T(), context, (&GoGitHubContext{}).New(clientFactory, new(MockRepositoryTransformer), context, model.SessionContext{}).context)
}

func (suite *GitHubSuite) TestCreatesGoGitHubContextWithSessionContext() {
	clientFactory := func(httpClient *http.Client) *github.Client { return &github.Client{} }
	sessionContext := model.SessionContext{}

	assert.Equal(suite.T(), sessionContext, (&GoGitHubContext{}).New(clientFactory, new(MockRepositoryTransformer), context.Background(), sessionContext).sessionContext)
}

func TestGitHubSuite(t *testing.T) {
	suite.Run(t, new(GitHubSuite))
}
