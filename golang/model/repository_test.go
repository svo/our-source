package model

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
  suite.Suite
}

func (suite *RepositorySuite) SetupTest() {
}

func (suite *RepositorySuite) TestCreatesRepositoryWithName() {
  assert.Equal(suite.T(), "coconuts", Repository{}.New("coconuts").name)
}

func (suite *RepositorySuite) TestReturnsName() {
  assert.Equal(suite.T(), "coconuts",  Repository{name: "coconuts"}.Name())
}

func TestRepositorySuite(t *testing.T) {
  suite.Run(t, new(RepositorySuite))
}
