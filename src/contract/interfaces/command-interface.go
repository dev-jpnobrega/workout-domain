package interfaces

import "github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"

// A ICommand represent command interface
type ICommand interface {
	GetModelValidate() *valueObject.ValidateModal
	Execute(valueObject.RequestData) (valueObject.ResponseData, *valueObject.ResponseError)
}
