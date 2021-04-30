package server

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/svo/our-source/golang/repository"
)

type InitSuite struct {
	suite.Suite
}

type MockRouterFactory struct {
	mock.Mock
}

func (router *MockRouterFactory) Build(gitHubContext repository.GitHub) *gin.Engine {
	args := router.Called()
	return args.Get(0).(*gin.Engine)
}

func (suite *InitSuite) TestUsesRouterFactory() {
	routerFactory := new(MockRouterFactory)
	router := new(gin.Engine)

	routerFactory.On("Build").Return(router)

	result := Init(routerFactory)

	assert.Equal(suite.T(), router, result)
	routerFactory.AssertExpectations(suite.T())
}

func TestInitSuite(t *testing.T) {
	suite.Run(t, new(InitSuite))
}
