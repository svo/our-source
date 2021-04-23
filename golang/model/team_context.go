package model

type TeamContext struct {
  organisation string
  team string
}

func (context TeamContext) New(organisation string, team string) TeamContext {
  return TeamContext { organisation, team }
}

func (context TeamContext) Organisation() string {
  return context.organisation
}

func (context TeamContext) Team() string {
  return context.team
}
