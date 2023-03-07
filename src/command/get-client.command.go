package domain

import (
	"github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"
	"github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
)

// A GetClientCommand represent business logic to search clients
type GetClientCommand struct {
	Repository interfaces.IClientRepository
}

func (c *GetClientCommand) GetModelValidate() *valueObject.ValidateModal {
	return &valueObject.ValidateModal{
		Modal: &valueObject.SearchArgs{},
	}
}

func (g *GetClientCommand) Execute(input valueObject.RequestData) (
	result valueObject.ResponseData, err *valueObject.ResponseError,
) {
	args := input.Args.(*valueObject.SearchArgs)

	clients, err := g.Repository.Get(args)

	if err != nil {
		return
	}

	result.Data = clients
	result.StatusCode = 200

	return
}
