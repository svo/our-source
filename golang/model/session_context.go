package model

type SessionContext struct {
  accessToken string
}

func (context SessionContext) New(accessToken string) SessionContext {
  return SessionContext { accessToken }
}

func (context SessionContext) AccessToken() string {
  return context.accessToken
}
