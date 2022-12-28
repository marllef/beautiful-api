package server_test

import (
	"marllef/beautiful-api/frameworks/server"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	require.NotNil(t, server.NewServer())
}