package model

import (
	"testing"

	"net/url"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
}

func (suite *RepositorySuite) SetupTest() {
}

func (suite *RepositorySuite) TestCreatesRepositoryWithName() {
	assert.Equal(suite.T(), "coconuts", Repository{}.New("coconuts", url.URL{}).name)
}

func (suite *RepositorySuite) TestReturnsName() {
	assert.Equal(suite.T(), "coconuts", Repository{name: "coconuts"}.Name())
}

func (suite *RepositorySuite) TestCreatesRepositoryWithUrl() {
	expected := url.URL{}

	assert.Equal(suite.T(), expected, Repository{}.New("foo", url.URL{}).url)
}

func (suite *RepositorySuite) TestReturnsUrl() {
	expected := url.URL{}

	assert.Equal(suite.T(), expected, Repository{url: expected}.Url())
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
