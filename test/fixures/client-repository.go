package domain

import (
	values "github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
	entity "github.com/dev-jpnobrega/workout-domain/src/entity"
)

// A ClientRepository represent client database
type ClientRepository struct{}

func (c *ClientRepository) Get(params interface{}) (*[]entity.Client, *values.ResponseError) {
	clients := &[]entity.Client{}

	return clients, nil
}

func (c *ClientRepository) Create(client *entity.Client) (*entity.Client, *values.ResponseError) {

	return client, nil
}

func (c *ClientRepository) Update(client *entity.Client) (bool, *values.ResponseError) {

	return true, nil
}
