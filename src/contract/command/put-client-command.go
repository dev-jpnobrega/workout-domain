package domain

import (
	"github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"
	"github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
	"github.com/dev-jpnobrega/workout-domain/src/entity"
)

// A PutClientComannd represent business logic to edit client
type PutClientComannd struct {
	Repository interfaces.IClientRepository
}

func (c *PutClientComannd) GetModelValidate() *valueObject.ValidateModal {
	return &valueObject.ValidateModal{
		Modal: &valueObject.PutArgs{},
	}
}

func (c *PutClientComannd) Execute(input valueObject.RequestData) (
	result valueObject.ResponseData, err *valueObject.ResponseError,
) {
	client := input.Args.(*entity.Client)

	_, err = c.Repository.Update(client)

	if err != nil {
		return
	}

	result.Data = client
	result.StatusCode = 200

	return
}
