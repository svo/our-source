package model

import "net/url"

type Repository struct {
  name string
  url url.URL
}

func (repository Repository) New(name string, url url.URL) Repository {
  return Repository { name, url }
}

func (repository Repository) Name() string {
  return repository.name
}

func (repository Repository) Url() url.URL {
  return repository.url
}
