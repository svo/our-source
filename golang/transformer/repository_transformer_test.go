package transformer

import (
	"testing"

	"net/url"

	"github.com/google/go-github/v35/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTransformerSuite struct {
	suite.Suite
}

func (suite *RepositoryTransformerSuite) SetupTest() {}

func (suite *RepositoryTransformerSuite) TestTransformsGoGitHubRepositoryName() {
	test_url := "https://github.com/svo/our-source"
	name := "coconuts"
	from := github.Repository{Name: &name, HTMLURL: &test_url}
	assert.Equal(suite.T(), name, RepositoryTransformer{}.Transform(from).Name())
}

func (suite *RepositoryTransformerSuite) TestTransformsGoGitHubRepositorUrl() {
	test_url := "https://github.com/svo/our-source"
	name := "coconuts"
	expected_url, _ := url.Parse(test_url)
	from := github.Repository{Name: &name, HTMLURL: &test_url}
	assert.Equal(suite.T(), *expected_url, RepositoryTransformer{}.Transform(from).Url())
}

func TestRepositoryTransformerSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTransformerSuite))
}
