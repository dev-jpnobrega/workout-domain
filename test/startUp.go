package domain_test

import (
	"testing"

	"github.com/dev-jpnobrega/workout-domain/src/entity"
	"github.com/google/uuid"
)

var ClientsMock []entity.Client

func TestMain(m *testing.M) {
	client1 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}
	client2 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}
	client3 := entity.Client{ID: uuid.New(), Name: "JP", Age: 10}

	ClientsMock = append(ClientsMock, client1, client2, client3)

	m.Run()
}
