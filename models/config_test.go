package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRandomConfigList(t *testing.T) {
	result, err := testQueries.ConfigGetList()
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
