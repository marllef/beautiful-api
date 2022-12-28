package entities_test

import (
	entities "marllef/beautiful-api/internal/app/models/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	require.NotNil(t, entities.NewProduct("Roteador", "358-57"))
}