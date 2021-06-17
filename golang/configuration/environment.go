package configuration

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	AccessToken string `required:"true"`
}

type Configuration interface {
	GetAccessToken() string
}

type EnvironmentConfiguration struct {
}

func (config EnvironmentConfiguration) GetAccessToken() string {
	var specification Specification

	err := envconfig.Process("our_source", &specification)

	if err != nil {
		log.Fatal(err.Error())
	}

	return specification.AccessToken
}
