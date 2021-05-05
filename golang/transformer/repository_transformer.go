package transformer

import (
	"net/url"

	"github.com/google/go-github/v35/github"

	"github.com/svo/our-source/golang/model"
)

type RepositoryToModelTransformer interface {
	Transform(from github.Repository) model.Repository
}

type RepositoryTransformer struct{}

func (repository RepositoryTransformer) Transform(from github.Repository) model.Repository {
	name := *from.Name
	html_url := *from.HTMLURL
	repository_url, _ := url.Parse(html_url)
	return (&model.Repository{}).New(name, *repository_url)
}
