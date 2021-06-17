package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvironmentConfigurationSuite struct {
	suite.Suite
}

func (suite *EnvironmentConfigurationSuite) TestGetAccessToken() {
	environment := EnvironmentConfiguration{}
	assert.Equal(suite.T(), environment.GetAccessToken(), "coconuts")
}

func TestEnvironmentConfigurationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentConfigurationSuite))
}
