package model

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
)

type SessionContextSuite struct {
  suite.Suite
}

func (suite *SessionContextSuite) SetupTest() {
}

func (suite *SessionContextSuite) TestCreatesSessionContextWithAccessToken() {
  assert.Equal(suite.T(), SessionContext{accessToken: "coconuts"}, SessionContext{}.New("coconuts"))
}

func (suite *SessionContextSuite) TestReturnsAccessToken() {
  assert.Equal(suite.T(), "coconuts",  SessionContext{accessToken: "coconuts"}.AccessToken())
}

func TestSessionContextSuite(t *testing.T) {
  suite.Run(t, new(SessionContextSuite))
}
