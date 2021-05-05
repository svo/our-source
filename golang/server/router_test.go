package server

import (
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/svo/our-source/golang/model"
)

type MockGitHub struct {
	mock.Mock
}

func (gitHub *MockGitHub) Select(teamContext model.TeamContext) []model.Repository {
	args := gitHub.Called(teamContext)
	return args.Get(0).([]model.Repository)
}

func performRequest(router http.Handler, method, path string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

type GinRouterSuite struct {
	suite.Suite
	router        http.Handler
	gitHubContext *MockGitHub
}

func (suite *GinRouterSuite) SetupTest() {
	suite.gitHubContext = new(MockGitHub)
	suite.router = GinRouterFactory{}.Build(suite.gitHubContext)
}

func (suite *GinRouterSuite) TestPingStatusOk() {
	response := performRequest(suite.router, "GET", "/ping")

	assert.Equal(suite.T(), http.StatusOK, response.Code)
}

func (suite *GinRouterSuite) TestPingMessageBody() {
	body := gin.H{
		"message": "pong",
	}

	response := performRequest(suite.router, "GET", "/ping")

	var response_body map[string]string
	err := json.Unmarshal(response.Body.Bytes(), &response_body)
	value, exists := response_body["message"]

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), exists)
	assert.Equal(suite.T(), body["message"], value)
}

func (suite *GinRouterSuite) TestTeamRepositoryStatusOk() {
	var expected []model.Repository
	suite.gitHubContext.On("Select", (&model.TeamContext{}).New("coconuts", "bob")).Return(expected)

	response := performRequest(suite.router, "GET", "/organisation/coconuts/team/bob/repository")

	assert.Equal(suite.T(), http.StatusOK, response.Code)
}

func TestGinRouterSuite(t *testing.T) {
	suite.Run(t, new(GinRouterSuite))
}
