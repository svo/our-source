package model

type UserContext struct {
  accessToken string
}

func (context UserContext) New(accessToken string) UserContext {
  return UserContext { accessToken }
}

func (context UserContext) AccessToken() string {
  return context.accessToken
}
