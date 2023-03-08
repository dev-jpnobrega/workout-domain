package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	command "github.com/dev-jpnobrega/workout-domain/src/command"
	valueObject "github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
	entity "github.com/dev-jpnobrega/workout-domain/src/entity"

	fixures "github.com/dev-jpnobrega/workout-domain/test/fixures"
)

func TestCreateClientCommandSucess(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		data := new(valueObject.RequestData)
		client := new(entity.Client)

		client.CompanyID = 1
		client.Email = "test@test.com"
		client.Name = "JP"

		commandExec := &command.CreateClientComannd{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = client

		result, err := commandExec.Execute(*data)

		fmt.Printf("%+v\n", result)
		fmt.Printf("%+v\n", err)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.EqualValues(t, 200, result.StatusCode)
	})
}
