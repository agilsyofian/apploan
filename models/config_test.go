package models

import (
	"testing"

	"github.com/agilsyofian/golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomConfig(t *testing.T) (*Config, error) {
	payload := Config{
		Name:     util.RandomString(5),
		Desc:     util.RandomString(20),
		Constant: float64(util.RandomInt(1, 10)),
	}
	result, err := testQueries.ConfigCreate(payload)
	return result, err
}

func TestCreateRandomConfig(t *testing.T) {
	result, err := createRandomConfig(t)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestGetRandomConfig(t *testing.T) {
	config, err := createRandomConfig(t)
	require.NoError(t, err)
	require.NotEmpty(t, config)

	result, err := testQueries.ConfigGet(config.ID)
	require.NoError(t, err)
	require.NotEmpty(t, config)
	require.Equal(t, config.ID, result.ID)
}
