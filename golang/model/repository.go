package model

type Repository struct {
  name string
}

func (context Repository) New(name string) Repository {
  return Repository { name }
}

func (context Repository) Name() string {
  return context.name
}
