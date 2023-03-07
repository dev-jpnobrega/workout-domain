package interfaces

import (
	"github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
)

// A IClientRepository represent repository client interface
type IClientRepository interface {
	// Get(params interface{}) (*[]entity.Client, *valueObject.ResponseError)
	Create(client *valueObject.ResponseError) (*valueObject.ResponseError, *valueObject.ResponseError)
	// Update(client *entity.Client) (bool, *valueObject.ResponseError)
}
