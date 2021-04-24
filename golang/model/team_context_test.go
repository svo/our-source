package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TeamContextSuite struct {
	suite.Suite
}

func (suite *TeamContextSuite) SetupTest() {
}

func (suite *TeamContextSuite) TestCreatesTeamContextWithOrganisation() {
	assert.Equal(suite.T(), "coconuts", TeamContext{}.New("coconuts", "foo").organisation)
}

func (suite *TeamContextSuite) TestReturnsOrganisation() {
	assert.Equal(suite.T(), "coconuts", TeamContext{organisation: "coconuts"}.Organisation())
}

func (suite *TeamContextSuite) TestCreatesTeamContextWithTeam() {
	assert.Equal(suite.T(), "coconuts", TeamContext{}.New("foo", "coconuts").team)
}

func (suite *TeamContextSuite) TestReturnsTeam() {
	assert.Equal(suite.T(), "coconuts", TeamContext{team: "coconuts"}.Team())
}

func TestTeamContextSuite(t *testing.T) {
	suite.Run(t, new(TeamContextSuite))
}
