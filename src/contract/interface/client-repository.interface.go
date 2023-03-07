package interfaces

import (
	"github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
	"github.com/dev-jpnobrega/workout-domain/src/entity"
)

// A IClientRepository represent repository client interface
type IClientRepository interface {
	Get(params interface{}) (*[]entity.Client, *valueObject.ResponseError)
	Create(client *entity.Client) (*entity.Client, *valueObject.ResponseError)
	Update(client *entity.Client) (bool, *valueObject.ResponseError)
}
