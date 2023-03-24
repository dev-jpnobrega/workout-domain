package domain_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	command "github.com/dev-jpnobrega/workout-domain/src/command"
	valueObject "github.com/dev-jpnobrega/workout-domain/src/contract/valueObject"
	entity "github.com/dev-jpnobrega/workout-domain/src/entity"

	fixures "github.com/dev-jpnobrega/workout-domain/test/fixures"
)

var ClientsMock []entity.Client

func TestMain(m *testing.M) {
	client1 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}
	client2 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}
	client3 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}

	ClientsMock = append(ClientsMock, client1, client2, client3)

	fmt.Printf("%+v\n DOMAIN", ClientsMock)

	os.Exit(m.Run())
}

func TestGetClientCommand(t *testing.T) {

	t.Run("Get clients without filters", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(valueObject.SearchArgs)

		fmt.Printf("%+v\n", ClientsMock)

		commandExec := &command.GetClientCommand{
			Repository: &fixures.ClientRepository{
				ClientsMock: ClientsMock,
			},
		}

		data.Args = args

		result, err := commandExec.Execute(*data)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.EqualValues(t, 200, result.StatusCode)
	})

	t.Run("Get clients with error", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(valueObject.SearchArgs)
		args.Name = "thorw"

		commandExec := &command.GetClientCommand{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = args

		_, err := commandExec.Execute(*data)

		assert.NotNil(t, err)
		assert.EqualValues(t, 400, err.StatusCode)
		assert.Len(t, err.Errors, 1)
	})

	t.Run("GetClient validate model", func(t *testing.T) {
		data := new(valueObject.RequestData)
		args := new(valueObject.SearchArgs)

		commandExec := &command.GetClientCommand{
			Repository: &fixures.ClientRepository{},
		}

		data.Args = args

		validated := commandExec.GetModelValidate()

		assert.Equal(t, args, validated.Modal)
	})
}
