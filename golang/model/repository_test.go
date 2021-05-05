package model

import (
	"testing"

	"encoding/json"
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

func (suite *RepositorySuite) TestGeneratesCorrectJson() {
	expected_name := "coconuts"
	expected_url := "https://github.com/svo/our-source"
	test_url, _ := url.Parse(expected_url)

	result, _ := json.Marshal(Repository{}.New(expected_name, *test_url))

	assert.Equal(suite.T(), "{\"name\":\"coconuts\",\"url\":\"https://github.com/svo/our-source\"}", string(result))
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
