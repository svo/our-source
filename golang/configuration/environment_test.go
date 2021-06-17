package configuration

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EnvironmentConfigurationSuite struct {
	suite.Suite
}

func (suite *EnvironmentConfigurationSuite) TestGetAccessToken() {
	environment := EnvironmentConfiguration{}
	assert.Equal(suite.T(), "coconuts", environment.GetAccessToken())
}

func (suite *EnvironmentConfigurationSuite) TestExitsIfNoAccessToken() {
	if os.Getenv("ASSERT_EXISTS_"+suite.T().Name()) == "1" {
		environment := EnvironmentConfiguration{}
		environment.GetAccessToken()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+suite.T().Name())
	cmd.Env = append([]string{}, "ASSERT_EXISTS_"+suite.T().Name()+"=1")

	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	suite.T().Fatal("should have exited due to missing access token environment variable")
}

func TestEnvironmentConfigurationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentConfigurationSuite))
}
