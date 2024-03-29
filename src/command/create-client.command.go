package domain

import (
	"github.com/dev-jpnobrega/workout-domain/src/contract/interfaces"
	"github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
	"github.com/dev-jpnobrega/workout-domain/src/entity"
	"github.com/google/uuid"
)

// A CreateClientComannd represent business logic to create client
type CreateClientComannd struct {
	Repository interfaces.IClientRepository
}

func (c *CreateClientComannd) GetModelValidate() *valueObject.ValidateModal {
	return &valueObject.ValidateModal{
		Modal: &entity.Client{},
	}
}

func (c *CreateClientComannd) Execute(input valueObject.RequestData) (
	result valueObject.ResponseData, err *valueObject.ResponseError,
) {
	client := input.Args.(*entity.Client)
	client.ID = uuid.New()

	if !client.IsValid() {
		err = err.New("user.not.valid", 100, 422)

		return
	}

	_, err = c.Repository.Create(client)

	if err != nil {
		return
	}

	result.Data = client
	result.StatusCode = 200

	return
}
