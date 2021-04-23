package model

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
)

type UserContextSuite struct {
  suite.Suite
}

func (suite *UserContextSuite) SetupTest() {
}

func (suite *UserContextSuite) TestCreatesUserContextWithAccessToken() {
  assert.Equal(suite.T(), UserContext{accessToken: "coconuts"}, UserContext{}.New("coconuts"))
}

func (suite *UserContextSuite) TestReturnsAccessToken() {
  assert.Equal(suite.T(), "coconuts",  UserContext{accessToken: "coconuts"}.AccessToken())
}

func TestUserContextSuite(t *testing.T) {
  suite.Run(t, new(UserContextSuite))
}
